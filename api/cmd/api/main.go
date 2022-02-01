package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ippsav/awesome-links/api/cmd/api/config"
)

func main() {
	// parse config
	config := &config.Config{}
	if err := config.Parse(); err != nil {
		fmt.Printf("could not parse config, err: %s ", err.Error())
		os.Exit(1)
	}

	// router
	r := gin.Default()

	// server
	srv := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
	// routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// start server
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("could not start server, err: %s", err.Error())
	}
}
