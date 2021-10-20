package app

import (
	"context"
	"github.com/mebr0/grpc-client/internal/config"
	"github.com/mebr0/grpc-client/internal/game"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	// Load configs
	cfg := config.LoadConfig(configPath)

	c, err := game.NewGenerator(":" + cfg.GRPC.Port, cfg.GRPC.Period)

	if err != nil {
		log.Fatal(err)
	}

	if err = c.Start(); err != nil {
		log.Errorf("failed to stop client: %v", err)
	}

	log.Info("Client started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = c.Stop(ctx); err != nil {
		log.Errorf("failed to stop client: %v", err)
	}

	log.Info("Client stopped")
}
