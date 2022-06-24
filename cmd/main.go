package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"stream-redirector/api"
)

// content holds our static web server content.
//go:embed frontend/*
var content embed.FS

func main() {
	r := gin.Default()
	serverRoot, err := fs.Sub(content, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}
	r.GET("/api/get-url", api.GetUrl)
	r.POST("/api/set-url", api.SetUrl)
	r.GET("/api/redirect", api.RedirectToUrl)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/app")
	})
	r.StaticFS("/app", http.FS(serverRoot))

	err = r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
