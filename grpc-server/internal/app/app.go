package app

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/mebr0/grpc-server/internal/api"
	"github.com/mebr0/grpc-server/internal/config"
	"github.com/mebr0/grpc-server/internal/server"
	"github.com/mebr0/grpc-server/internal/service"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	// Load configs
	cfg := config.LoadConfig(configPath)

	u := url.URL{
		Scheme: "ws",
		Host:   cfg.WS.Host,
		Path:   "/api/v1/games",
	}

	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal(err)
	}

	s := service.NewServices(c)
	a := api.NewAPI(s)

	g := a.Init()

	srv := server.NewServer(g)

	go func() {
		if err := srv.Start(cfg); err != nil {
			log.Errorf("failed to serve: %s", err)
		}
	}()

	log.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}

	if err = c.Close(); err != nil {
		log.Errorf("failed to disconnect from websocket: %v", err)
	}

	log.Info("Server stopped")
}
