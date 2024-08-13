package main

import (
	"context"
	"fmt"

	"github.com/mineamihai2001/fm/internal/api/router"
	"github.com/mineamihai2001/fm/internal/helpers"
	"github.com/mineamihai2001/fm/internal/infrastructure/opentelemetry"
	"github.com/mineamihai2001/fm/pkg/amqp"
	"github.com/mineamihai2001/fm/pkg/logger"
	"github.com/mineamihai2001/fm/pkg/tracing"
	"github.com/rs/zerolog/log"
)

func main() {
	env := helpers.Env()

	amqpClient, err := amqp.
		NewClient(amqp.ClientConfig{
			Protocol: env.Log.Writer.Protocol,
			Host:     env.Log.Writer.Host,
			User:     env.Log.Writer.User,
			Password: env.Log.Writer.Password,
			Port:     env.Log.Writer.Port,
			VHost:    env.Log.Writer.VHost,
		})

	if err != nil {
		logger.Init(opentelemetry.GetTraceId)
	} else {
		defer amqpClient.Conn().Close()

		logger.Init(
			opentelemetry.GetTraceId,
			amqpClient.RegisterQueue(env.Log.Writer.Queue),
		)
	}

	tracer := tracing.Init()
	defer func() {
		if err := tracer.Shutdown(context.Background()); err != nil {
			log.Fatal().Err(err).Msg("Shutting down tracer provider")
		}
	}()

	app := router.Create()

	log.
		Debug().
		Msg(fmt.Sprintf("Server starting on  http://[::1]:%d", env.App.Port))

	if err := app.Run(fmt.Sprintf(":%d", env.App.Port)); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}
}
