package consensus

import (
	"github.com/kgemio/kaspad/infrastructure/logger"
	"github.com/kgemio/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
