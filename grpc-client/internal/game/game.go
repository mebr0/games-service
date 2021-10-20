package game

import (
	pb "github.com/mebr0/grpc-client/pkg/api/v1"
	"math/rand"
	"time"
)

const (
	min = 0
	max = 10
)

var (
	id    int32 = 1
	teams       = []string{
		"navi",
		"team liquid",
		"team spirit",
		"astralis",
	}
)

func Random() *pb.Game {
	rand.Seed(time.Now().UnixNano())

	g := &pb.Game{
		Id:          id,
		Team_1:      teams[rand.Intn(len(teams))],
		Team_2:      teams[rand.Intn(len(teams))],
		ScoreTeam_1: rand.Int31n(max-min+1) + min,
		ScoreTeam_2: rand.Int31n(max-min+1) + min,
	}

	id += 1

	return g
}
