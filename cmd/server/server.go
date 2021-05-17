package server

import (
	"fmt"
	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-api-scaffold/internal/pkg/config"
	"web-api-scaffold/internal/pkg/constant"
	"web-api-scaffold/internal/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	router   *fiber.App
	database *gorm.DB
	logger   *logger.Logger
}

func New() *Server {
	return &Server{
		router: fiber.New(
			fiber.Config{
				ReadTimeout:           1 * time.Second,
				BodyLimit:             constant.MaxBodyLimit,
				ReadBufferSize:        constant.MaxBufferSize,
				WriteBufferSize:       constant.MaxBufferSize,
				ReduceMemoryUsage:     true,
				DisableStartupMessage: true,
				ServerHeader: fmt.Sprintf("%s/%.2f",
					constant.BinFileSererName, constant.BinFileVersion),
			},
		),
		logger: logger.NewServerLogger(),
	}
}

func (s *Server) ListenAndServe() {
	fmt.Printf(
		"\n\t██████╗ ██╗███╗   ██╗\t███████╗██╗██╗     ███████╗"+
			"\n\t██╔══██╗██║████╗  ██║\t██╔════╝██║██║     ██╔════╝"+
			"\n\t██████╔╝██║██╔██╗ ██║\t█████╗  ██║██║     █████╗"+
			"\n\t██╔══██╗██║██║╚██╗██║\t██╔══╝  ██║██║     ██╔══╝"+
			"\n\t██████╔╝██║██║ ╚████║\t██║     ██║███████╗███████╗"+
			"\n\t╚═════╝ ╚═╝╚═╝  ╚═══╝\t╚═╝     ╚═╝╚══════╝╚══════╝"+
			"  [%.2f]\n\n", constant.BinFileVersion)

	var err error
	var listener net.Listener

	go func() {
		addr := fmt.Sprintf("%s:%d",
			config.Instance().Server.Host, config.Instance().Server.Port)

		if listener, err = net.Listen(fiber.NetworkTCP4, addr); err != nil {
			s.logger.Fatal().Err(err).Msgf("HTTP server addr: %s listen failed", addr)
		}

		s.logger.Debug().Msgf("HTTP server listening on http://%s", addr)
		if err = s.router.Listener(listener); err != nil {
			s.logger.Fatal().Err(err).Msg("HTTP server error happen")
		}
	}()

	s.HandleQuitSignal()
}

func (s *Server) HandleQuitSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	s.logger.Debug().Msgf("Got exit signal: %v", <-quit)
	if pool, err := s.database.DB(); err == nil {
		s.logger.Debug().Msg("Close database connection...")
		if err = pool.Close(); err != nil {
			s.logger.Fatal().Err(err).Msg("Close database connection failed")
		}
	}

	s.logger.Debug().Msg("Server exiting...")
	if err := s.router.Shutdown(); err != nil {
		s.logger.Fatal().Err(err).Msg("Server exit failed")
	}
	s.logger.Debug().Msg("Server exit.")
}
