package infrastructure

import (
	"api/infrastructure/interface/rest"
	rest_rooms "api/infrastructure/interface/rest/rooms"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server *echo.Echo
}

func NewServer() *Server {
	server := new(Server)
	server.server = echo.New()
	server.setup()

	return server
}

func (s *Server) Run(port string) {
	s.server.Logger.Fatal(s.server.Start(port))
}

func (s *Server) RunForServerless(w http.ResponseWriter, r *http.Request) {
	s.setup()
	s.server.ServeHTTP(w, r)
}

func (s *Server) setup() {
	s.setConfig()
	s.setRoutes()
}

func (s *Server) setConfig() {
	s.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	s.server.Use(middleware.Recover())
	s.server.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody))
		fmt.Fprintf(os.Stderr, "Headers: %v\n", c.Request().Header)
		fmt.Fprintf(os.Stderr, "Method: %v\n", c.Request().Method)
	}))
}

func (s *Server) setRoutes() {
	s.server.GET("/login", rest.Login)

	s.server.POST("/rooms", rest.CreateRoom)

	s.server.GET("/rooms/:room_id/users", rest_rooms.GetUsers)

	s.server.GET("/rooms/:room_id/payments", rest_rooms.GetPayments)
	s.server.POST("/rooms/:room_id/payments", rest_rooms.CreatePayment)
	s.server.DELETE("/rooms/:room_id/payments", rest_rooms.DeletePayment)
}
