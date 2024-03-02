package commpub

import (
	"github.com/fengjx/glca/logic/common/commdto"
	"github.com/fengjx/glca/logic/common/internal/service"
)

func RegisterTableConfig(config commdto.TableConfig) {
	service.CommonService.RegisterTableConfig(config)
}
