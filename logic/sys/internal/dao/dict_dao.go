package dao

import (
	"reflect"

	"github.com/fengjx/daox"

	"github.com/fengjx/glca/integration/db"
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
)

var DictDao = newSysDictDao()

type sysDictDao struct {
	*daox.Dao
}

func newSysDictDao() *sysDictDao {
	inst := &sysDictDao{}
	inst.Dao = daox.NewDAO(
		db.GetDefaultDB(),
		"sys_dict",
		"id",
		reflect.TypeOf(&entity.SysDict{}),
	)
	return inst
}
