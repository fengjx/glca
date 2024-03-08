package provider

import "github.com/fengjx/glca/logic/common/commpub"

func Init() {
	commpub.SetCommonAPI(&CommonProvider{})
}
