package main

import (
	"embed"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hypernova120.net/go-webserver-template/api"
)

var embeddedFiles embed.FS

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)

	//subFS, err := fs.Sub(embeddedFiles, "static")
	dir, err := static.EmbedFolder(embeddedFiles, "static")
	if err != nil {
		panic(err)
	}

	//r.StaticFS("/", http.FS(subFS))
	r.Use(static.Serve("/", dir))
	r.Run(":8080")
}
