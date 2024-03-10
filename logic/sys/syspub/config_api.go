package syspub

import "github.com/fengjx/luchen"

const (
	ScopeSys = "sys"
)

var ConfigAPI configAPI

type configAPI interface {
	// GetConfigString 返回key对应的配置
	GetConfigString(scope string, key string) (string, error)

	// GetConfig 返回key对应的配置，并序列化成对象
	GetConfig(scope string, key string, data any) error
}

func SetConfigAPI(impl configAPI) {
	luchen.RootLogger().Info("set ConfigAPI")
	ConfigAPI = impl
}
