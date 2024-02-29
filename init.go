package main

import (
	"time"

	"github.com/json-iterator/go/extra"
)

func init() {
	extra.RegisterTimeAsInt64Codec(time.Millisecond)
}
