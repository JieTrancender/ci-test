package main

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "pong",
		})
	})

	return router
}

func main() {
	host := flag.String("h", "0.0.0.0", "host name or ip address")
	port := flag.Int("p", 8080, "port")

	flag.CommandLine.Parse(os.Args[1:])

	router := setupRouter()

	server := &http.Server{
		Addr:           *host + ":" + strconv.Itoa(*port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
