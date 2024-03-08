package provider

import "github.com/fengjx/glca/logic/sys/syspub"

func Init() {
	syspub.SetUserAPI(&UserProvider{})
}
