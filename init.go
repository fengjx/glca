package main

import (
	"time"

	"github.com/fengjx/go-halo/json"
)

func init() {
	json.RegisterTimeAsInt64Codec(time.Millisecond)
}
