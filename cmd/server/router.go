package server

import (
	"fmt"
	"strings"
	"web-api-scaffold/internal/app/handler"
	"web-api-scaffold/internal/app/middleware"
	"web-api-scaffold/internal/app/repository"
	"web-api-scaffold/internal/app/service"
	"web-api-scaffold/internal/pkg/requestor"
)

func (s *Server) SetMiddleware() *Server {
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.RequestID())
	s.router.Use(middleware.Logger(s.logger))
	return s
}

func (s *Server) RegisterRoutes() *Server {
	userRepository := repository.NewUserRepository(s.database)
	userService := service.NewUserService(userRepository)

	deviceInfoHandler := handler.NewDeviceInfoHandler()
	deviceListHandler := handler.NewDeviceListHandler(userService)
	deviceBindHandler := handler.NewDeviceBindHandler(userService)

	fileRepository := repository.NewFileRepository(s.database)
	fileService := service.NewFileService(fileRepository)

	fileInitHandler := handler.NewFileInitHandler(fileService)
	fileCreateHandler := handler.NewFileCreateHandler(fileService)

	v1 := s.router.Group(version(1))
	v1.Get("/device/info", requestor.Bind(deviceInfoHandler))

	v2 := s.router.Group(version(2))
	device := v2.Group("/device")
	device.Get("/dashboard", middleware.Monitor())
	device.Get("/info", requestor.Bind(deviceInfoHandler))
	device.Get("/list", requestor.Bind(deviceListHandler))
	device.Post("/bind", requestor.Bind(deviceBindHandler))

	file := v2.Group("/file", middleware.Authorizer(s.database))
	file.Post("/init", requestor.Bind(fileInitHandler))
	file.Post("/create", requestor.Bind(fileCreateHandler))

	return s
}

func version(i uint8) string {
	return buildRouteWithPrefix("api", fmt.Sprintf("v%d", i))
}

func buildRouteWithPrefix(prefix, route string) string {
	return strings.TrimRight(prefix, "/") + "/" + strings.TrimLeft(route, "/")
}
