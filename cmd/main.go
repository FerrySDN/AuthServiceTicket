package main

import (
	"log"

	"github.com/FerrySDN/auth-service/internal/core/auth"
	"github.com/FerrySDN/auth-service/internal/db"
	httphandler "github.com/FerrySDN/auth-service/internal/handler/http"
	jwtadapter "github.com/FerrySDN/auth-service/internal/repository/jwt"
	pg "github.com/FerrySDN/auth-service/internal/repository/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	d := db.Connect()
	userRepo := pg.NewUserRepository(d)
	tokenSvc := jwtadapter.NewJWTService()
	authSvc := auth.NewService(userRepo, tokenSvc)
	h := httphandler.NewHandler(authSvc)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "auth-service OK"})
	})

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
	}

	log.Println("Auth service running on :9000")
	if err := r.Run(":9000"); err != nil {
		log.Fatal(err)
	}
}
