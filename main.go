package main

import (
	"flag"
	"log"

	"github.com/ikasamt/go-epgstation/pkg"

	"github.com/gin-gonic/gin"
)

var hddBase string
var serverHost string

func main() {
	flag.StringVar(&serverHost, "serverHost", "", "serverHost")
	flag.StringVar(&hddBase, "HDDBase", "", "HDDBase")
	flag.Parse()

	if serverHost == `` {
		log.Println("no serverHost")
		return
	}

	if hddBase == `` {
		log.Println("no hddBase")
		return
	}

	log.SetFlags(log.Llongfile)

	// init variables
	pkg.HDDBase = hddBase
	pkg.ServerHost = serverHost
	pkg.CachedChannel = pkg.APIGETChannel()
	pkg.CachedResponse = pkg.VideoResponse{}

	// gin server
	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
	r.GET("/api/refresh", func(c *gin.Context) {
		pkg.SetCacheVideoResponse()
		c.JSON(200, gin.H{"size": pkg.CachedResponse.Total})
	})
	r.GET("/api/total", func(c *gin.Context) {
		c.JSON(200, gin.H{"size": pkg.CachedResponse.Total})
	})
	r.GET("/jpeg/:ID", pkg.JPEGIndexHandler)
	r.GET("/mp4/:Key", pkg.VideoServeHandler)
	r.GET("/api/remove/:ID", pkg.APIRemoveHandler)

	r.Run("0.0.0.0:3001")
}
