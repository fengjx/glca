package service

import (
	"context"
	"strings"

	"github.com/fengjx/go-halo/utils"

	"github.com/fengjx/glca/connom/kit"
	"github.com/fengjx/glca/logic/sys/syspub"
)

func newSysUserConfig() tableConfig {
	m := syspub.SysUserMeta

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
	cfg := tableConfig{
		tableName:          m.TableName,
		insertFieldsFilter: insertFieldsFilter,
		insertDataWrapper:  insertDataWrapper,
	}
	return cfg
}
