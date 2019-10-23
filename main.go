package main

import (
	"fmt"
	"github.com/getupandgo/highly-uncertain-test-task/internal/config"
	"github.com/getupandgo/highly-uncertain-test-task/internal/controller"
	"github.com/getupandgo/highly-uncertain-test-task/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.InitConfiguration()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to get configuration file")
	}

	if !conf.GetBool("http_debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	c := controller.Controller{}

	s := server.Init(c)

	httpPort := conf.GetInt("http_port")

	if err = s.Engine.Run(fmt.Sprintf(":%d", httpPort)); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start server")
	}

	log.Info().Msgf("starting API server on port %d", httpPort)
}
