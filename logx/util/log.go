package util

import (
	"sync/atomic"
)

var (
	logLevel uint32 = 0
)

func SetRootLevel(level uint32) {
	atomic.StoreUint32(&logLevel, level)
}

func ShallLog(level uint32) bool {
	return atomic.LoadUint32(&logLevel) <= level
}
