package rpc

import (
	"github.com/kgemio/kaspad/infrastructure/logger"
	"github.com/kgemio/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
