package provider

import (
	"github.com/fengjx/glca/logic/common/commpub"
	"github.com/fengjx/glca/logic/common/internal/service"
)

type CommonProvider struct {
}

func (p *CommonProvider) RegisterTableConfig(config commpub.TableConfig) {
	service.CommonService.RegisterTableConfig(config)
}
