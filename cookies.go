package gotenberg

import (
	"errors"
)

const (
	sameSiteStrict = "Strict"
	sameSiteLax    = "Lax"
	sameSiteNone   = "None"
)

var errRequiredCookieFieldEmpty = errors.New("required cookie field empty")

type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Domain   string `json:"domain"`
	Path     string `json:"path,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	HTTPOnly bool   `json:"http_only,omitempty"`
	SameSite string `json:"same_site,omitempty"`
}

func (c *Cookie) validate() error {
	if c.Name == "" || c.Value == "" || c.Domain == "" {
		return errRequiredCookieFieldEmpty
	}

	if c.SameSite != sameSiteStrict && c.SameSite != sameSiteLax && c.SameSite != sameSiteNone {
		c.SameSite = ""
	}

	return nil
}
