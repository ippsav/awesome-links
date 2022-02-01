package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
