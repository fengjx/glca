package service

import (
	"context"

	"github.com/fengjx/glca/logic/sys/syspub"
)

var UserSvc = &userService{}

type userService struct {
}

func (svc *userService) GetByUsername(ctx context.Context, username string) (*syspub.UserDetailInfoDTO, error) {
	userDetailDTO, err := syspub.UserAPI.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return userDetailDTO, nil
}
