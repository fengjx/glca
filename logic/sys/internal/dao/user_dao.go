package dao

import (
	"reflect"

	"github.com/fengjx/daox"

	"github.com/fengjx/glca/integration/db"
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
)

var SysUserDao *sysUserDao

type sysUserDao struct {
	*daox.Dao
}

func init() {
	SysUserDao = newSysUserDao()
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
