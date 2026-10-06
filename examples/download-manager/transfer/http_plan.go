package transfer

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// plan is the negotiated shape of the response body.
type plan struct {
	start      int64 // file offset the body writes at
	truncate   bool  // open the partial with O_TRUNC first
	total      int64 // whole-entity length, Unknown when hidden
	validators Validators
	resumed    bool
	complete   bool // a 416 proved the partial already holds the entity
}

// negotiate turns a response into a write plan. It refuses any body it
// cannot append at a proven offset.
func negotiate(resp *http.Response, req Request, prep prepared) (plan, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return planFrom200(resp, prep)
	case http.StatusPartialContent:
		return planFrom206(resp, req, prep)
	case http.StatusRequestedRangeNotSatisfiable:
		return planFrom416(resp, req, prep)
	default:
		return plan{}, fmt.Errorf("%w: HTTP %d", ErrBadStatus, resp.StatusCode)
	}
}

// planFrom200 is a whole-body response. A previous offset is dropped: the
// server either ignored Range or If-Range proved the entity changed.
func planFrom200(resp *http.Response, prep prepared) (plan, error) {
	if err := encodingOK(resp); err != nil {
		return plan{}, err
	}

	return plan{
		start:      0,
		truncate:   true,
		total:      contentLength(resp),
		validators: validatorsOf(resp),
		resumed:    false,
	}, nil
}

// planFrom206 accepts a range answer only when its start equals the offset
// we asked for and the entity length did not change.
func planFrom206(resp *http.Response, req Request, prep prepared) (plan, error) {
	if err := encodingOK(resp); err != nil {
		return plan{}, err
	}

	start, _, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return plan{}, err
	}

	if start != prep.offset {
		return plan{}, fmt.Errorf("%w: server started at %d, wanted %d", ErrBadRange, start, prep.offset)
	}

	if total >= 0 && req.Expected > 0 && total != req.Expected {
		return plan{}, fmt.Errorf("%w: length changed from %d to %d", ErrBadRange, req.Expected, total)
	}

	if total < 0 {
		total = contentLength(resp)
	}

	return plan{
		start:      start,
		truncate:   false,
		total:      total,
		validators: mergeValidators(validatorsOf(resp), req.Validators),
		resumed:    prep.resumed,
	}, nil
}

// planFrom416 treats a 416 as success only when the partial file already
// matches the expected length, which is the shape an interrupted complete
// download leaves behind.
func planFrom416(resp *http.Response, req Request, prep prepared) (plan, error) {
	if req.Expected > 0 && prep.offset == req.Expected {
		if _, _, total, err := parseContentRange(resp.Header.Get("Content-Range")); err == nil {
			if total >= 0 && total != req.Expected {
				return plan{}, fmt.Errorf("%w: length changed from %d to %d", ErrBadRange, req.Expected, total)
			}
		}

		return plan{complete: true, total: req.Expected, validators: req.Validators}, nil
	}

	return plan{}, fmt.Errorf("%w: HTTP 416 at offset %d of %d", ErrBadRange, prep.offset, req.Expected)
}

// contentLength reads Content-Length, mapping anything unusable to Unknown.
func contentLength(resp *http.Response) int64 {
	if resp.ContentLength < 0 {
		return Unknown
	}

	return resp.ContentLength
}

// validatorsOf copies ETag and Last-Modified off a response.
func validatorsOf(resp *http.Response) Validators {
	return Validators{
		ETag:         resp.Header.Get("ETag"),
		LastModified: resp.Header.Get("Last-Modified"),
	}
}

// mergeValidators keeps the response validator when present and falls back
// to the saved one, so a silent server cannot erase resume state.
func mergeValidators(fresh, saved Validators) Validators {
	if fresh.ETag == "" {
		fresh.ETag = saved.ETag
	}

	if fresh.LastModified == "" {
		fresh.LastModified = saved.LastModified
	}

	return fresh
}

// encodingOK rejects a body that arrived encoded: the bytes do not line up
// with Content-Length offsets and cannot be checksummed as the entity.
func encodingOK(resp *http.Response) error {
	enc := resp.Header.Get("Content-Encoding")
	if enc == "" || enc == "identity" {
		return nil
	}

	return fmt.Errorf("%w: Content-Encoding %s", ErrUnsupportedResume, enc)
}

// parseContentRange splits "bytes 10-19/100" or "bytes */100".
func parseContentRange(value string) (start, end, total int64, err error) {
	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, "bytes ")

	span, totalText, ok := strings.Cut(value, "/")
	if !ok {
		return 0, 0, 0, fmt.Errorf("%w: bad Content-Range %q", ErrBadRange, value)
	}

	total, err = contentRangeTotal(totalText)
	if err != nil {
		return 0, 0, 0, err
	}

	if span == "*" {
		return 0, 0, total, nil
	}

	left, right, ok := strings.Cut(span, "-")
	if !ok {
		return 0, 0, 0, fmt.Errorf("%w: bad Content-Range %q", ErrBadRange, value)
	}

	start, err = strconv.ParseInt(left, 10, 64)
	if err != nil || start < 0 {
		return 0, 0, 0, fmt.Errorf("%w: bad Content-Range %q", ErrBadRange, value)
	}

	end, err = strconv.ParseInt(right, 10, 64)
	if err != nil || end < start {
		return 0, 0, 0, fmt.Errorf("%w: bad Content-Range %q", ErrBadRange, value)
	}

	return start, end, total, nil
}

// contentRangeTotal parses the total field; "*" and empty mean Unknown.
func contentRangeTotal(text string) (int64, error) {
	text = strings.TrimSpace(text)
	if text == "" || text == "*" {
		return Unknown, nil
	}

	total, err := strconv.ParseInt(text, 10, 64)
	if err != nil || total < 0 {
		return 0, fmt.Errorf("%w: bad total %q", ErrBadRange, text)
	}

	return total, nil
}
