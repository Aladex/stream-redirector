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

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, token")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	r := gin.Default()
	r.Use(Options)
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
