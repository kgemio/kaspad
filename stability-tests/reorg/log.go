package main

import (
	"github.com/kgemio/kaspad/infrastructure/logger"
	"github.com/kgemio/kaspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
