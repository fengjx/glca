package hello

import "sync"

type Inst struct {
	GreetSvc  *greetService
	Endpoints *endpoints
}

var ins *Inst
var insOnce sync.Once

func GetInst() *Inst {
	insOnce.Do(func() {
		ins = &Inst{
			GreetSvc:  newGreetService(),
			Endpoints: newEndpoints(),
		}
	})
	return ins
}
