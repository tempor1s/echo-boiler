package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tempor1s/echo-boiler/handlers"
)

// Routes will setup all the routes and
func (s *Server) Routes() {
	s.e.Use(middleware.Logger())
	// s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))

	// create a new handler and dependency inject the db and the tournament manager
	r := handlers.New(s.db)

	// Heya!
	s.e.GET("/", r.Hello)
	// Setup all your routes here!
}
