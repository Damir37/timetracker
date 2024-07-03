package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"timertracker/internel/handler"
	"timertracker/internel/service"
)

type Server struct {
	ServiceContext service.ServiceContext
	FiberApp       *fiber.App
}

func NewServer(serviceContext service.ServiceContext) *Server {
	return &Server{
		ServiceContext: serviceContext,
	}
}

func (server *Server) StartWebServer() error {
	// ИНЦИАЛИЗАЦИЯ И НАСТРОЙКА
	app := fiber.New(fiber.Config{
		AppName:           server.ServiceContext.Config.AppName,
		EnablePrintRoutes: true,
		ProxyHeader:       fiber.HeaderXForwardedFor,
	})

	//НАСТРОЙКА КОРС

	//ЛОГГЕР
	app.Use(logger.New())

	handler.RegisterHandlers(app, server.ServiceContext)

	server.FiberApp = app

	err := app.Listen(server.ServiceContext.Config.Bind)
	if server.ServiceContext.Config.Debug {
		log.Println(fmt.Sprintf("DEBUG: Start web server %s", server.ServiceContext.Config.Bind))
	}

	return err
}
