package service

var SysDictService = newSysDictService()

type sysDictService struct {
}

func newSysDictService() *sysDictService {
	inst := &sysDictService{}
	return inst
}
