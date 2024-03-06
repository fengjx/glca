package types

import (
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
	"github.com/fengjx/glca/protocol"
)

func BuildUserInfoPBFromEntity(user *entity.SysUser) *protocol.UserInfo {
	if user == nil {
		return nil
	}
	return &protocol.UserInfo{
		UID:      user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
	}
}
