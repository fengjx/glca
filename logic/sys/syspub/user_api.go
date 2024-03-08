package syspub

import (
	"context"

	"github.com/fengjx/luchen"
)

var UserAPI userAPI

type userAPI interface {
	// GetByUsername 根据用户名查询用户信息
	GetByUsername(ctx context.Context, username string) (*UserDetailInfoDTO, error)
}

func SetUserAPI(impl userAPI) {
	if UserAPI != nil {
		luchen.RootLogger().Warn("SetUserAPI duplicate")
		return
	}
	UserAPI = impl
}
