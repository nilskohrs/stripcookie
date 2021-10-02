// Package stripcookie a plugin to strip cookies.
package stripcookie

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	Cookies []string `json:"cookies,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// CookieStrip a CookieStrip plugin.
type CookieStrip struct {
	next        http.Handler
	cookies     []string
	name        string
	splitRegexp *regexp.Regexp
}

// New created a new CookieStrip plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Cookies == nil || len(config.Cookies) == 0 {
		return nil, fmt.Errorf("cookies cannot be empty")
	}

	return &CookieStrip{
		cookies:     config.Cookies,
		next:        next,
		name:        name,
		splitRegexp: regexp.MustCompile(` *([^=;]+?) *=[^;]+`),
	}, nil
}

func (c *CookieStrip) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	cookieHeaders := req.Header.Values("cookie")
	req.Header.Del("cookie")
	for _, cookieHeader := range cookieHeaders {
		cookies := c.splitRegexp.FindAllStringSubmatch(cookieHeader, -1)
		var keep []string
		for _, cookie := range cookies {
			if !stringInSlice(cookie[1], c.cookies) {
				keep = append(keep, cookie[0])
			}
		}
		if len(keep) > 0 {
			req.Header.Add("cookie", strings.TrimSpace(strings.Join(keep, ";")))
		}
	}
	c.next.ServeHTTP(rw, req)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
