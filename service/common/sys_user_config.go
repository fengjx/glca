package common

import (
	"context"
	"strings"

	"github.com/fengjx/go-halo/utils"

	"github.com/fengjx/glca/connom/kit"
	"github.com/fengjx/glca/data/meta"
)

func newSysUserConfig() tableConfig {
	m := meta.SysUserMeta

	insertFieldsFilter := func(ctx context.Context) []string {
		return []string{"login_time"}
	}

	insertDataWrapper := func(ctx context.Context, src map[string]any) map[string]any {
		pwd := utils.ToString(src[m.Pwd])
		salt := utils.ToString(src[m.Salt])
		if pwd == "" || salt == "" {
			delete(src, m.Pwd)
			delete(src, m.Salt)
			return src
		}
		sb := strings.Builder{}
		sb.WriteString(pwd)
		sb.WriteString(salt)
		md5Pwd := kit.MD5Hash(sb.String())
		src[m.Pwd] = md5Pwd
		return src
	}
	cfg := tableConfig{
		tableName:          m.TableName,
		insertFieldsFilter: insertFieldsFilter,
		insertDataWrapper:  insertDataWrapper,
	}
	return cfg
}
