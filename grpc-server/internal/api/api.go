package api

import (
	"github.com/mebr0/grpc-server/internal/api/v1"
	"github.com/mebr0/grpc-server/internal/service"
)

type API struct {
	services *service.Services
}

func NewAPI(services *service.Services) *API {
	return &API{
		services: services,
	}
}

func (a *API) Init() *v1.Game {
	apiV1 := v1.NewAPI(a.services)

	return apiV1.Init()
}
