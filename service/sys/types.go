package sys

import (
	"github.com/fengjx/glca/data/entity"
	"github.com/fengjx/glca/pb"
)

// UserDTO dto for entity.SysUser
type UserDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}

func buildUserInfoPB(user *entity.SysUser) *pb.UserInfo {
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
