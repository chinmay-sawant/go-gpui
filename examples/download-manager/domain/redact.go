package domain

import (
	"net/url"
	"strings"
)

// redacted stands in for anything that may carry a secret.
const redacted = "[redacted]"

// RedactURL strips userinfo and any query or fragment values that may hold
// credentials, and returns the URL with a redacted userinfo and a query
// that keeps only key names.
func RedactURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return redacted
	}

	if u.User != nil {
		u.User = url.User(redacted)
	}

	if u.RawQuery != "" {
		u.RawQuery = redactQuery(u.Query())
	}

	if u.Fragment != "" {
		u.Fragment = redacted
	}

	return u.String()
}

// redactQuery keeps parameter names and blanks the values.
func redactQuery(values url.Values) string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}

	// url.Values.Encode sorts keys, so the loop above only builds the slice.
	out := url.Values{}
	for _, key := range keys {
		out.Set(key, redacted)
	}

	return out.Encode()
}

// Redact blanks the value in an Authorization-style header line.
func Redact(header string) string {
	if i := strings.IndexByte(header, ':'); i >= 0 {
		return header[:i+1] + " " + redacted
	}

	return redacted
}
