package prefixmanager

import (
	"github.com/kgemio/kaspad/infrastructure/logger"
	"github.com/kgemio/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
