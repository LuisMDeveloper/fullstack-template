// Package webui serves the Vue build embedded in the binary.
// Run `npm run build` in web/ first: it writes to ./dist.
package webui

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
)

//go:embed all:dist
var embedded embed.FS

// Handler serves the static files and falls back to index.html for
// client-side routes (SPA).
func Handler() http.Handler {
	dist, err := fs.Sub(embedded, "dist")
	if err != nil {
		panic(err)
	}
	files := http.FileServerFS(dist)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := path.Clean(r.URL.Path)[1:]
		if name == "" {
			name = "index.html"
		}
		if _, err := fs.Stat(dist, name); err != nil {
			r.URL.Path = "/"
		}
		files.ServeHTTP(w, r)
	})
}
