package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

var content embed.FS

//go:embed build

func clientHandler() http.Handler {
	fsys := fs.FS(content)
	contentStatic, _ := fs.Sub(fsys, "build")
	return http.FileServer(http.FS(contentStatic))

}

func erve() {
	mux := http.NewServeMux()
	mux.Handle("/", clientHandler())
	http.ListenAndServe(":3000", mux)
}
