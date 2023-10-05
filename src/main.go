package main

import (
	"fmt"
	"log"
	"os"

	adapters "github.com/NBN23dev/go-service-template/src/adapters/grpc"
	"github.com/NBN23dev/go-service-template/src/core/services"
	"github.com/NBN23dev/go-service-template/src/plugins/logger"
	"github.com/NBN23dev/go-service-template/src/plugins/tracer"
	"github.com/NBN23dev/go-service-template/src/server"
	"github.com/NBN23dev/go-service-template/src/utils"
)

func main() {
	// Env vars
	name, _ := utils.GetEnvOr("SERVICE_NAME", "unknown")
	port, _ := utils.GetEnvOr("PORT", 8080)

	// Tracer
	if err := tracer.Init(); err != nil {
		log.Fatal(err)
	}

	// Logger
	if err := logger.Init(name, logger.LevelInfo); err != nil {
		log.Fatal(err)
	}

	// Service
	repos := services.Repositories{}
	service, err := services.NewService(repos)

	if err != nil {
		log.Fatal(err)
	}

	adapter := adapters.NewGRPCAdapter(service)

	// Create server
	server, err := server.NewServer(adapter)

	if err != nil {
		log.Fatal(err)
	}

	// Shutdown
	go server.GracefulShutdown(func(sig os.Signal) {
		logger.Info(fmt.Sprintf("'%s' service it is about to end", name), nil)
	})

	logger.Info(fmt.Sprintf("'%s' service it is about to start", name), nil)

	// Start server
	server.Start(port)
}
