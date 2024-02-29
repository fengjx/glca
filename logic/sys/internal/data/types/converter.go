package types

import (
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
	"github.com/fengjx/glca/pb"
)

func BuildUserInfoPBFromEntity(user *entity.SysUser) *pb.UserInfo {
	if user == nil {
		return nil
	}
	return &pb.UserInfo{
		Uid:      user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
	}
}
