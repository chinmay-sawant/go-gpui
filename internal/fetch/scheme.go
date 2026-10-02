package fetch

import (
	"errors"
	"net/http"
	"net/url"
)

// client is shared. Jar stays nil, so no cookie is stored or sent.
var client = &http.Client{CheckRedirect: redirect}

func checkURL(raw string) error {
	if raw == "" {
		return ErrScheme
	}

	u, err := url.Parse(raw)
	if err != nil {
		return err
	}
	if !web(u) {
		return ErrScheme
	}

	return nil
}

func web(u *url.URL) bool {
	return u != nil && (u.Scheme == "http" || u.Scheme == "https")
}

func redirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	if !web(req.URL) {
		return ErrScheme
	}

	return nil
}
