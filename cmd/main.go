package main

import (
	"context"
	"os"
	"syscall"

	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/logger"
	"github.com/hyper-micro/hyper/provider/server"
	"github.com/hyper-micro/project-layout/cmd/wireinject"
)

var (
	buildDate   = "-"
	buildCommit = "-"
	version     = "v0.0.1"
	appName     = "HyperApp"
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
		AppName:      appName,
		Version:      version,
		BuildCommit:  buildCommit,
		BuildDate:    buildDate,
		ShutdownSigs: []os.Signal{syscall.SIGTERM},
	})
	if err != nil {
		return err
	}

	if err := srv.RegServe(func(cfg config.Config) (server.App, server.CleanUpHandler, error) {
		return wireinject.NewHttpServer(ctx, cfg)
	}); err != nil {
		return err
	}

	if err := srv.RegServe(func(cfg config.Config) (server.App, server.CleanUpHandler, error) {
		return wireinject.NewRpcServer(ctx, cfg)
	}); err != nil {
		return err
	}

	return srv.Run()
}
