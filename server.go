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

	e.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig {
				AllowOrigins: []string{ "*" },
				AllowMethods: []string{ "*" },
			},
		),
	)

	// Public Endpoints

	// Login
	e.POST("/login", handlers.LoginHandler)

	// Users
	e.POST("/users", handlers.CreateUserHandler)
	e.GET("/users", handlers.GetAllUsersHandler)
	e.GET("/users/:id", handlers.GetUserByIDHandler)
	e.DELETE("/users/:id", handlers.RemoveUserByIDHandler)
	e.PUT("/users/:id", handlers.UpdateUserByIDHandler)

	// Users
	e.POST("/customers", handlers.CreateCustomerHandler)
	e.GET("/customers", handlers.GetAllCustomersHandler)
	e.GET("/customers/:id", handlers.GetCustomerByIDHandler)
	e.DELETE("/customers/:id", handlers.RemoveCustomerByIDHandler)
	e.PUT("/customers/:id", handlers.UpdateCustomerByIDHandler)

	// Providers
	e.POST("/providers", handlers.CreateProviderHandler)
	e.GET("/providers", handlers.GetAllProvidersHandler)
	e.GET("/providers/:id", handlers.GetProviderByIDHandler)
	e.DELETE("/providers/:id", handlers.RemoveProviderByIDHandler)
	e.PUT("/providers/:id", handlers.UpdateProviderByIDHandler)

	// Roosters
	e.POST("/roosters", handlers.CreateRoosterHandler)
	e.GET("/roosters", handlers.GetAllRoostersHandler)
	e.GET("/roosters/:id", handlers.GetRoosterByIDHandler)
	e.DELETE("/roosters/:id", handlers.RemoveRoosterByIDHandler)
	e.PUT("/roosters/:id", handlers.UpdateRoosterByIDHandler)

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
