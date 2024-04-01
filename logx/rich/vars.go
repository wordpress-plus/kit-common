package rich

const CallerDepth = 4

const (
	// DebugLevel logs everything
	DebugLevel uint32 = iota
	// InfoLevel include debugs
	InfoLevel
	// WarnLevel include debugs, infos
	WarnLevel
	// ErrorLevel includes errors, slows, stacks
	ErrorLevel
)

const (
	callerKey   = "caller"
	durationKey = "duration"
	spanKey     = "spanId"
	traceKey    = "traceId"
)
