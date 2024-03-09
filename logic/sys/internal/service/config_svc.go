package service

var ConfigService = newSysConfigService()

type sysConfigService struct {
}

func newSysConfigService() *sysConfigService {
	inst := &sysConfigService{}
	return inst
}
