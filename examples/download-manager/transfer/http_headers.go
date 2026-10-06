package transfer

import (
	"fmt"
	"net/http"
)

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
