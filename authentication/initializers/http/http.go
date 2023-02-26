package http

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

var server *echo.Echo

// Gzip compression level.
const gzip = 5

type Config struct {
	ServerPort string `env:"SERVER_PORT"`
}

func New() {
	server = echo.New()
	server.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	//nolint:exhaustivestruct
	loggerConfig := middleware.LoggerConfig{
		Output: log.StandardLogger().Writer(),
	}

	server.Use(middleware.LoggerWithConfig(loggerConfig))
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())
	//nolint:exhaustivestruct
	server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzip,
	}))

}

func Listen() {
	var cfg Config = Config{
		ServerPort: os.Getenv("SERVER_PORT"),
	}
	server.Logger.Info("Server config: ", cfg)
	server.Logger.Info("Starting server on port: ", cfg.ServerPort)
	server.Logger.Fatal(server.Start(":" + cfg.ServerPort))
}

func Get() *echo.Echo {
	return server
}
