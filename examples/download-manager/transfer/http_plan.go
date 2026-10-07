package transfer

import (
	"fmt"
	"net/http"
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
		return planFrom200(resp)
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
func planFrom200(resp *http.Response) (plan, error) {
	if err := encodingOK(resp); err != nil {
		return plan{}, err
	}

	return plan{
		start:      0,
		truncate:   true,
		total:      contentLength(resp),
		validators: validatorsOf(resp),
	}, nil
}
