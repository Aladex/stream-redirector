package api

import (
	"github.com/gin-gonic/gin"
	"os"
)

var StreamUrl string

type StreamUrlRequest struct {
	Url string `json:"url" form:"url"`
}

func init() {
	StreamUrl = os.Getenv("STREAM_URL")
	if StreamUrl == "" {
		StreamUrl = "https://google.com"
	}
}

func GetUrl(c *gin.Context) {
	c.JSON(200, gin.H{"url": StreamUrl})
}

func SetUrl(c *gin.Context) {
	var streamUrlReq StreamUrlRequest
	if err := c.ShouldBindJSON(&streamUrlReq); err == nil {
		StreamUrl = streamUrlReq.Url
	}
	c.JSON(200, gin.H{"url": StreamUrl})
}

func RedirectToUrl(c *gin.Context) {
	c.Redirect(301, StreamUrl)
}
