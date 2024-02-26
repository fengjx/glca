package common

import "sync"

type Inst struct {
	Endpoints *endpoints
	CommonSvc *commonService
}

var ins *Inst
var insOnce sync.Once

func GetInst() *Inst {
	insOnce.Do(func() {
		ins = &Inst{
			Endpoints: newEndpoints(),
			CommonSvc: newCommonService(),
		}
	})
	return ins
}
