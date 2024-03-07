package syspub

import (
	"time"

	"github.com/fengjx/glca/protocol"
)

// UserDetailInfoDTO dto for entity.SysUser
type UserDetailInfoDTO struct {
	UserInfoDTO
	Pwd    string    `json:"pwd"`
	Salt   string    `json:"salt"`
	Remark string    `json:"remark"`
	Utime  time.Time `json:"utime"`
	Ctime  time.Time `json:"ctime"`
}

func (dto *UserDetailInfoDTO) ToUserInfoPB() *protocol.UserInfo {
	if dto == nil {
		return nil
	}
	return &protocol.UserInfo{
		UID:      dto.ID,
		Username: dto.Username,
		Nickname: dto.Nickname,
		Email:    dto.Email,
		Avatar:   dto.Avatar,
		Phone:    dto.Phone,
	}
}

// UserInfoDTO dto for entity.SysUser
type UserInfoDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}
