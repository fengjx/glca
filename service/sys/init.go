package sys

import "sync"

type Inst struct {
	Endpoints *endpoints
	userDao   *sysUserDao
	UserSvc   *sysUserService
}

var ins *Inst
var insOnce sync.Once

func GetInst() *Inst {
	insOnce.Do(func() {
		ins = &Inst{
			Endpoints: newEndpoints(),
			userDao:   newSysUserDao(),
			UserSvc:   newSysUserService(),
		}
	})
	return ins
}
