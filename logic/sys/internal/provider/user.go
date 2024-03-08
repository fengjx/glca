package provider

import (
	"context"

	"github.com/fengjx/glca/logic/sys/internal/service"
	"github.com/fengjx/glca/logic/sys/syspub"
)

var UserProvider = &userProvider{}

type userProvider struct {
}

func (impl *userProvider) GetByUsername(ctx context.Context, username string) (*syspub.UserDetailInfoDTO, error) {
	sysUser, err := service.SysUserService.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	userDetailInfo := &syspub.UserDetailInfoDTO{
		UserInfoDTO: syspub.UserInfoDTO{
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
