package sys

import (
	"context"

	"github.com/fengjx/luchen"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"

	"github.com/fengjx/glca/connom/errno"
	"github.com/fengjx/glca/pb"
)

func (e *endpoints) MakeLoginEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log := luchen.Logger(ctx)
		req := request.(*pb.LoginReq)
		userSvc := GetInst().UserSvc
		user, err := userSvc.getByUsername(ctx, req.Username)
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
		if user.Pwd != req.Password {
			return nil, errno.PasswordErr
		}
		resp := &pb.LoginResp{}
		resp.UserInfo = buildUserInfoPB(user)
		return resp, nil
	}
}
