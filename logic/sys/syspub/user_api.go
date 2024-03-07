package syspub

import (
	"context"

	"github.com/fengjx/glca/logic/sys/internal/service"
)

var UserAPI userAPI = &userAPIImpl{}

type userAPI interface {
	// GetByUsername 根据用户名查询用户信息
	GetByUsername(ctx context.Context, username string) (*UserDetailInfoDTO, error)
}

type userAPIImpl struct {
}

func (impl *userAPIImpl) GetByUsername(ctx context.Context, username string) (*UserDetailInfoDTO, error) {
	sysUser, err := service.SysUserService.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	userDetailInfo := &UserDetailInfoDTO{
		UserInfoDTO: UserInfoDTO{
			ID:       sysUser.ID,
			Username: sysUser.Username,
			Nickname: sysUser.Nickname,
			Email:    sysUser.Email,
			Avatar:   sysUser.Avatar,
			Phone:    sysUser.Phone,
			Status:   sysUser.Status,
		},
		Pwd:    sysUser.Pwd,
		Salt:   sysUser.Salt,
		Remark: sysUser.Remark,
		Utime:  sysUser.Utime,
		Ctime:  sysUser.Ctime,
	}
	return userDetailInfo, nil
}
