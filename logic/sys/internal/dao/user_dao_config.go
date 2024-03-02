package dao

import (
	"context"
	"strings"

	"github.com/fengjx/go-halo/utils"
	"github.com/samber/lo"

	"github.com/fengjx/glca/connom/kit"
	"github.com/fengjx/glca/logic/common/commdto"
	"github.com/fengjx/glca/logic/common/commpub"
	"github.com/fengjx/glca/logic/sys/internal/data/meta"
)

func init() {
	m := meta.SysUserMeta
	insertFieldsFilter := func(ctx context.Context) []string {
		return []string{"login_time"}
	}
	insertDataWrapper := func(ctx context.Context, src map[string]any) map[string]any {
		pwd := utils.ToString(src[m.Pwd])
		if pwd == "" {
			delete(src, m.Pwd)
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
	columns := lo.Filter[string](meta.SysUserColumns, func(item string, i int) bool {
		return !lo.Contains[string]([]string{"pwd", "salt"}, item)
	})
	cfg := commdto.TableConfig{
		TableName:          meta.SysUserTableName,
		Columns:            columns,
		InsertFieldsFilter: insertFieldsFilter,
		InsertDataWrapper:  insertDataWrapper,
	}
	commpub.RegisterTableConfig(cfg)
}
