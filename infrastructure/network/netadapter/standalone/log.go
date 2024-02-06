package standalone

import (
	"github.com/kgemio/kaspad/infrastructure/logger"
	"github.com/kgemio/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
