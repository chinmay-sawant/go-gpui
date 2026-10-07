package transfer

import (
	"fmt"
	"net/http"
)

// planFrom416 treats a 416 as success only when the partial file already
// matches the expected length, the shape an interrupted complete download
// leaves behind.
func planFrom416(resp *http.Response, req Request, prep prepared) (plan, error) {
	if req.Expected <= 0 || prep.offset != req.Expected {
		return plan{}, fmt.Errorf("%w: HTTP 416 at offset %d of %d",
			ErrBadRange, prep.offset, req.Expected)
	}

	if _, _, total, err := parseContentRange(resp.Header.Get("Content-Range")); err == nil {
		if total >= 0 && total != req.Expected {
			return plan{}, fmt.Errorf("%w: length changed from %d to %d",
				ErrBadRange, req.Expected, total)
		}
	}

	return plan{complete: true, total: req.Expected, validators: req.Validators}, nil
}
