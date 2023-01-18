package handler

import (
	"allenmicro/proto/rand"
	"allenmicro/rand-service/service"
	"context"
)

type handler struct {
}

func (h handler) GetRand(ctx context.Context, request *rand.RandRequest, response *rand.RandResponse) error {
	max := request.GetMax()
	response.Result = service.GetRand(max)
	return nil
}

func Handler() rand.RandHandler {
	return handler{}
}
