package frontend

import (
	"embed"
	"fmt"
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

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("/", clientHandler())
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
