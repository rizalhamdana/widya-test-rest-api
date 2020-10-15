package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
	"github.com/rizalhamdana/widya-test/middleware"
)

const DefaultPort = 8081

func (s *Service) HTTPServerMain() {

	e := echo.New()
	e.Use(middleware.ServerHeader, middleware.Logger)
	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	authGroup := e.Group("api/v1")
	s.AuthHandler.Mount(authGroup)

	profileGroup := e.Group("api/v1")
	profileGroup.Use(s.AuthMiddlewareCustom.AuthorizationCustom)
	s.ProfileHandler.Mount(profileGroup)

	port := DefaultPort

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))
}
