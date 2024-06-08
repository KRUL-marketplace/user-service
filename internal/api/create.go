package api

import (
	"context"
	"user-service/internal/converter"
	desc "user-service/pkg/user-service"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserInfoRequest) (*desc.CreateUserInfoResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateUserInfoResponse{
		Id: id,
	}, nil
}
