package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

// https://www.akmittal.dev/posts/go-embed-files/  <- compiler angry if you don't put go:embed ABOVE embed.FS, it's positional


//go:embed build
var content embed.FS


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
