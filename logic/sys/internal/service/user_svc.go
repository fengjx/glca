package service

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"
	"go.uber.org/zap"

	"github.com/fengjx/glca/logic/sys/internal/dao"
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
	"github.com/fengjx/glca/logic/sys/internal/data/meta"
)

var SysUserService *sysUserService

type sysUserService struct {
}

func init() {
	SysUserService = newSysUserService()
}

func newSysUserService() *sysUserService {
	inst := &sysUserService{}
	return inst
}

func (svc *sysUserService) GetByUsername(ctx context.Context, username string) (*entity.SysUser, error) {
	log := luchen.Logger(ctx)
	user := &entity.SysUser{}
	ok, err := dao.SysUserDao.GetByColumn(daox.OfKv(meta.SysUserMeta.Username, username), user)
	if err != nil {
		log.Error("get user by username err", zap.Error(err))
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return user, nil
}
