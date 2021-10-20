package service

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	pb "github.com/mebr0/grpc-server/pkg/api/v1"
	log "github.com/sirupsen/logrus"
)

type GameService struct {
	c *websocket.Conn
}

func newGameService(c *websocket.Conn) *GameService {
	return &GameService{
		c: c,
	}
}

func (s *GameService) Send(ctx context.Context, game *pb.Game) error {
	u, err := json.Marshal(game)

	if err != nil {
		return err
	}

	if err = s.c.WriteMessage(websocket.TextMessage, u); err != nil {
		return err
	}

	_, message, err := s.c.ReadMessage()

	if err != nil {
		return err
	}

	log.Infof("recv: %s", message)

	return nil
}
