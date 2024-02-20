package sys

import (
	"reflect"

	"github.com/fengjx/daox"

	"github.com/fengjx/glca/data/entity"
	"github.com/fengjx/glca/integration/db"
)

type sysUserDao struct {
	*daox.Dao
}

func newSysUserDao() *sysUserDao {
	inst := &sysUserDao{}
	inst.Dao = daox.NewDAO(
		db.GetDefaultDB(),
		"sys_user",
		"id",
		reflect.TypeOf(&entity.SysUser{}),
	)
	return inst
}
