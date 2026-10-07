package transfer

import (
	"fmt"
	"net/http"
)

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
		return plan{}, fmt.Errorf("%w: server started at %d, wanted %d",
			ErrBadRange, start, prep.offset)
	}

	if total >= 0 && req.Expected > 0 && total != req.Expected {
		return plan{}, fmt.Errorf("%w: length changed from %d to %d",
			ErrBadRange, req.Expected, total)
	}

	if total < 0 {
		total = contentLength(resp)
	}

	return plan{
		start:      start,
		total:      total,
		validators: mergeValidators(validatorsOf(resp), req.Validators),
		resumed:    prep.resumed,
	}, nil
}
