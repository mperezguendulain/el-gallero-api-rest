package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mperezguendulain/el-gallero-api-rest/db"
	"github.com/mperezguendulain/el-gallero-api-rest/handlers"
	"github.com/mperezguendulain/el-gallero-api-rest/services/jwtserv"
)

func main() {

	defer db.MgoSession.Close()

	e := echo.New()

	// Public Endpoints
	e.POST("/login", handlers.LoginHandler)
	e.POST("/users", handlers.CreateUserHandler)
	e.GET("/users", handlers.GetAllUsersHandler)
	e.GET("/users/:id", handlers.GetUserByIDHandler)
	e.DELETE("/users/:id", handlers.RemoveUserByIDHandler)
	e.PUT("/users/:id", handlers.UpdateUserByIDHandler)

	// Protected Endpoints
	authGuard := e.Group("/elgallero/api/v1")

	// Auth Middlerware
	authGuard.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtserv.SecretKey,
	}))

	authGuard.GET("/saluda", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola, bienvenido...")
	})

	e.Logger.Fatal(e.Start(":5000"))

}
