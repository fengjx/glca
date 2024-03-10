package service

import (
	"context"

	"github.com/fengjx/luchen"
	"go.uber.org/zap"

	"github.com/fengjx/glca/logic/sys/internal/dao"
	"github.com/fengjx/glca/logic/sys/internal/data/types"
	"github.com/fengjx/glca/logic/sys/syspub"
)

var SysConfigSvc = newSysConfigService()

type sysConfigService struct {
	configMap map[string][]*syspub.ConfigDTO
}

func newSysConfigService() *sysConfigService {
	inst := &sysConfigService{}
	return inst
}

func (svc sysConfigService) RefreshConfig(ctx context.Context) {
	log := luchen.Logger(ctx)
	list, err := dao.ConfigDao.ListAll(ctx)
	if err != nil {
		log.Error("list all sys_config err", zap.Error(err))
		return
	}
	configMap := map[string][]*syspub.ConfigDTO{}
	for _, item := range list {
		configMap[item.Scope] = append(configMap[item.Scope], types.BuildConfigDTO(item))
	}
	svc.configMap = configMap
	log.Infof("refresh config, size: %d", len(list))
}

func (svc sysConfigService) ListScopeConfig(scope string) []*syspub.ConfigDTO {
	return svc.configMap[scope]
}

func (svc sysConfigService) GetAllConfig() map[string][]*syspub.ConfigDTO {
	return svc.configMap
}
