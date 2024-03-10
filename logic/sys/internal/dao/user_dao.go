package dao

import (
	"context"
	"reflect"
	"strings"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/utils"
	"github.com/samber/lo"

	"github.com/fengjx/glca/connom/consts"
	"github.com/fengjx/glca/connom/kit"
	"github.com/fengjx/glca/integration/db"
	"github.com/fengjx/glca/logic/common/commpub"
	"github.com/fengjx/glca/logic/sys/internal/data/entity"
	"github.com/fengjx/glca/logic/sys/internal/data/meta"
)

var SysUserDao *sysUserDao

type sysUserDao struct {
	*daox.Dao
}

func init() {
	SysUserDao = newSysUserDao()
}

func registerUserTableConfig() {
	m := meta.SysUserMeta
	fieldsFilter := func(ctx context.Context) []string {
		return []string{"login_time"}
	}
	insertDataWrapper := func(ctx context.Context, src map[string]any) map[string]any {
		if _, ok := src[m.Status]; !ok {
			src[m.Status] = consts.StatusNormal
		}
		pwd := utils.ToString(src[m.Pwd])
		if len(pwd) == 0 {
			delete(src, m.Pwd)
			delete(src, m.Salt)
			return src
		}
		salt := utils.RandomString(6)
		sb := strings.Builder{}
		sb.WriteString(pwd)
		sb.WriteString(salt)
		md5Pwd := kit.MD5Hash(sb.String())
		src[m.Pwd] = md5Pwd
		src[m.Salt] = salt
		return src
	}
	columns := lo.Filter[string](m.Columns(), func(item string, i int) bool {
		return !lo.Contains[string]([]string{"pwd", "salt"}, item)
	})

	updateDataWrapper := func(ctx context.Context, src map[string]any) map[string]any {
		pwd := utils.ToString(src[m.Pwd])
		if len(pwd) > 0 {
			salt := utils.RandomString(6)
			sb := strings.Builder{}
			sb.WriteString(pwd)
			sb.WriteString(salt)
			md5Pwd := kit.MD5Hash(sb.String())
			src[m.Pwd] = md5Pwd
			src[m.Salt] = salt
		}
		return src
	}

	// 通用查询过滤删除记录
	selectConditionWrapper := func(cond []daox.Condition) []daox.Condition {
		cond = append(cond, daox.Condition{
			Op:            daox.OpAnd,
			Field:         m.Status,
			ConditionType: daox.ConditionTypeNotEq,
			Vals:          []any{"del"},
		})
		return cond
	}

	cfg := commpub.TableConfig{
		TableName:              m.TableName(),
		PrimaryKey:             m.PrimaryKey(),
		SelectColumns:          columns,
		SelectConditionWrapper: selectConditionWrapper,
		InsertFieldsFilter:     fieldsFilter,
		InsertDataWrapper:      insertDataWrapper,
		UpdateFieldsFilter:     fieldsFilter,
		UpdateDataWrapper:      updateDataWrapper,
	}
	commpub.CommonAPI.RegisterTableConfig(cfg)
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
