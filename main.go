package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hypernova120.net/go-webserver-template/api"
)

//go:embed static
var embeddedFiles embed.FS

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)

	subFS, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		panic(err)
	}

	r.NoRoute(gin.WrapH(http.FileServer(http.FS(subFS))))
	r.Run(":8080")
}
