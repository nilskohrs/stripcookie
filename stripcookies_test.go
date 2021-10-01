package traefik_stripcookie_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nilskohrs/traefik-stripcookie"
)

func TestDemo(t *testing.T) {
	cfg := traefik_stripcookie.CreateConfig()
	cfg.Cookies = []string{"testCookie", "otherCookie"}

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_stripcookie.New(ctx, next, cfg, "stripcookie-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Cookie", "testCookie=testValue; testCookie2=testValue2; testCookie3=testValue3")
	req.Header.Add("Cookie", "testCookie=testValue; otherCookie=otherValue")

	handler.ServeHTTP(recorder, req)

	assertCookies(t, req, "testCookie2=testValue2; testCookie3=testValue3")
}

func assertCookies(t *testing.T, req *http.Request, expected string) {
	t.Helper()
	if len(req.Header.Values("Cookie")) > 1 {
		t.Errorf("too many headers")
	}
	if req.Header.Get("Cookie") != expected {
		t.Errorf("invalid header value: %s", req.Header.Get("Cookie"))
	}
}
