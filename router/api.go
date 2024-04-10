package router

import (
	"my-driver/handler"
	"my-driver/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API)SetupRouter() {
	// api.Echo.Use(middleware.Logger())
	api.Echo.Use(middleware.Recover())

	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	auth := api.Echo.Group("/auth", security.JWTMiddleware())
	auth.GET("/profile", api.UserHandler.HandleProfile)
}