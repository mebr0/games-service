package v1

import (
	"context"
	"github.com/mebr0/grpc-server/internal/service"
	pb "github.com/mebr0/grpc-server/pkg/api/v1"
	log "github.com/sirupsen/logrus"
)

type Game struct {
	services *service.Services
}

func (g *Game) Send(ctx context.Context, game *pb.Game) (*pb.Response, error) {
	log.Infof("game received with id: %d, teams: %s (%d) and %s (%d)",
		game.Id, game.Team_1, game.ScoreTeam_1, game.Team_2, game.ScoreTeam_2)

	if err := g.services.Game.Send(ctx, game); err != nil {
		log.Error(err)
		return nil, err
	}

	return &pb.Response{
		Message: "ok",
	}, nil
}
