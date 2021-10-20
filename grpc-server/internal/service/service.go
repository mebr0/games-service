package service

import (
	"context"
	"github.com/gorilla/websocket"
	pb "github.com/mebr0/grpc-server/pkg/api/v1"
)

type Game interface {
	Send(ctx context.Context, game *pb.Game) error
}

type Services struct {
	Game Game
}

func NewServices(c *websocket.Conn) *Services {
	return &Services{
		Game: newGameService(c),
	}
}
