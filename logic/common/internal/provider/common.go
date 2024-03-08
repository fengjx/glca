package provider

import (
	"github.com/fengjx/glca/logic/common/commpub"
	"github.com/fengjx/glca/logic/common/internal/service"
)

var CommonProvider = &commonProvider{}

type commonProvider struct {
}

func (p *commonProvider) RegisterTableConfig(config commpub.TableConfig) {
	service.CommonService.RegisterTableConfig(config)
}
