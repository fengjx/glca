package endpoint

import (
	"context"
	"strings"

	"github.com/fengjx/luchen"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"

	"github.com/fengjx/glca/connom/auth"
	"github.com/fengjx/glca/connom/errno"
	"github.com/fengjx/glca/connom/kit"
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
	"github.com/fengjx/glca/logic/sys/internal/data/types"
	"github.com/fengjx/glca/logic/sys/internal/service"
	"github.com/fengjx/glca/pb"
)

func MakeLoginEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.LoginReq)
		log := luchen.Logger(ctx).With(zap.String("username", req.Username))
		user, err := service.SysUserService.GetByUsername(ctx, req.Username)
		if err != nil {
			log.Error("user login err",
				zap.String("username", req.Username),
				zap.Error(err),
			)
			return nil, err
		}
		if user == nil {
			log.Warn("user not exist", zap.String("username", req.Username))
			return nil, errno.UserNotExistErr
		}
		if !checkPassword(user, req.Password) {
			return nil, errno.PasswordErr
		}
		token, err := auth.GenToken(auth.LoginPayload{
			UID: user.ID,
		})
		if err != nil {
			log.Error("gen token err", zap.Error(err))
			return nil, err
		}
		resp := &pb.LoginResp{}
		resp.UserInfo = types.BuildUserInfoPBFromEntity(user)
		resp.Token = token
		return resp, nil
	}
}

// checkPassword 检查密码是否匹配
func checkPassword(user *entity.SysUser, password string) bool {
	sb := strings.Builder{}
	sb.WriteString(password)
	sb.WriteString(user.Salt)
	md5Pwd := kit.MD5Hash(sb.String())
	return user.Pwd == md5Pwd
}
