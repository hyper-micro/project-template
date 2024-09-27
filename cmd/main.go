package main

import (
	"context"
	"os"
	"syscall"
	"time"

	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/logger"
	"github.com/hyper-micro/hyper/provider/server"
	"github.com/hyper-micro/project-template/cmd/wireinject"
)

var (
	buildDate   = "-"
	buildCommit = "-"
	version     = "v0.0.1"
	appName     = "ProjectTemplateService"
)

func main() {
	if err := run(); err != nil {
		logger.Errorf("Unrecoverable err: %v", err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv, _, err := server.NewProvider(server.Option{
		AppName:               appName,
		Version:               version,
		BuildCommit:           buildCommit,
		BuildDate:             buildDate,
		ShutdownSigs:          []os.Signal{syscall.SIGTERM, syscall.SIGINT},
		ShutdownDelayDuration: 1 * time.Second,
		ConfigDefault:         "conf",
	})
	if err != nil {
		return err
	}

	err = srv.RegServes(
		func(cfg config.Config) ([]server.App, server.CleanUpHandler, error) {
			s, cleanUp, err := wireinject.NewServer(ctx, cfg)
			return []server.App{
				s.HttpServer,
				s.RpcServer,
			}, cleanUp, err
		},
	)
	if err != nil {
		return err
	}

	return srv.Run()
}
