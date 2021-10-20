package v1

import "github.com/mebr0/grpc-server/internal/service"

type API struct {
	services *service.Services
	game     *Game
}

func NewAPI(services *service.Services) *API {
	return &API{
		services: services,
		game: &Game{
			services: services,
		},
	}
}

func (a *API) Init() *Game {
	return a.game
}
