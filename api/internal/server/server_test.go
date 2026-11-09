package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPIRoutes(t *testing.T) {
	h := New()
	for path, want := range map[string]string{
		"/api/health": `"status":"ok"`,
		"/api/hello":  `"message":"Hello, World!"`,
		"/api/info":   `"name":"fullstack-template"`,
	} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("GET", path, nil))
		if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), want) {
			t.Errorf("%s: got %d %q, want body containing %s", path, rec.Code, rec.Body.String(), want)
		}
		if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
			t.Errorf("%s: content-type %q", path, ct)
		}
	}
}

func TestUnknownAPIRouteIsJSON(t *testing.T) {
	for _, req := range []*http.Request{
		httptest.NewRequest("GET", "/api/nope", nil),
		httptest.NewRequest("POST", "/api/health", nil),
	} {
		rec := httptest.NewRecorder()
		New().ServeHTTP(rec, req)
		if rec.Code != http.StatusNotFound || rec.Header().Get("Content-Type") != "application/json" {
			t.Errorf("%s %s: got %d %q", req.Method, req.URL.Path, rec.Code, rec.Header().Get("Content-Type"))
		}
	}
}
