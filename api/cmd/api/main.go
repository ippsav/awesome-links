package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ippsav/awesome-links/api/cmd/api/config"
	"github.com/ippsav/awesome-links/api/cmd/api/handler"
	"github.com/ippsav/awesome-links/api/cmd/api/middleware"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/ent/hook"
	"github.com/ippsav/awesome-links/api/internal/adapter/controller"
	"github.com/ippsav/awesome-links/api/internal/adapter/repository"
	"github.com/ippsav/awesome-links/api/internal/adapter/service"
	_ "github.com/lib/pq"
)

func main() {
	// parse config
	config := &config.Config{}
	if err := config.Parse(); err != nil {
		fmt.Printf("could not parse config, err: %s ", err.Error())
		os.Exit(1)
	}
	// ent client
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.DBName, config.Database.Password, config.Database.SSLMode))
	if err != nil {
		fmt.Printf("could not open connection , err: %s ", err.Error())
	}
	defer client.Close()
	client.Use(hook.UserPasswordHook)

	// repositories
	ur := repository.NewUserRepository(client)
	lr := repository.NewLinkRepository(client)
	cr, err := repository.NewCloudinaryRepository(config)
	if err != nil {
		fmt.Printf("could not create cloudinary repository, err: %s ", err.Error())
	}
	// services
	us := service.NewUserService(ur, cr)
	ls := service.NewLinkService(lr, cr)
	// controllers
	controller := &controller.Controller{
		User: us,
		Link: ls,
	}

	// router
	r := gin.Default()
	// setup session store
	store := cookie.NewStore([]byte(config.Server.CookieSecret))
	// server
	srv := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
	// middlewares
	r.Use(sessions.Sessions("auth-store", store))
	r.Use(middleware.SessionMiddleware())
	r.Use(middleware.GinContextToContextMiddleware())
	r.Use(middleware.AuthMiddleware(config))
	// routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.POST("/query", handler.GraphqlHandler(client, controller, config))
	r.GET("/", handler.PlaygroundHandler())

	// start server
	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("could not start server, err: %s", err.Error())
	}
}
