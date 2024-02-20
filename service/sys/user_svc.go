package sys

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
	"go.uber.org/zap"

	"github.com/fengjx/glca/data/entity"
	"github.com/fengjx/glca/data/meta"
)

type sysUserService struct {
}

func newSysUserService() *sysUserService {
	inst := &sysUserService{}
	return inst
}

func (svc *sysUserService) getByUsername(ctx context.Context, username string) (*entity.SysUser, error) {
	log := luchen.Logger(ctx)
	user := &entity.SysUser{}
	userDao := GetInst().userDao
	ok, err := userDao.GetByColumn(daox.OfKv(meta.SysUserMeta.Username, username), user)
	if err != nil {
		log.Error("get user by username err", zap.Error(err))
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return user, nil
}
