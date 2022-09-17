package llogger

var (
	TraceLevel = 1
	DebugLevel = 2
	InfoLevel  = 3
	WarnLevel  = 4
	ErrorLevel = 5
)

func getTextByLevel(level int) string {
	switch level {
	case TraceLevel:
		return "[TRACE] "
	case DebugLevel:
		return "[DEBUG] "
	case ErrorLevel:
		return "[ERROR] "
	case WarnLevel:
		return "[WARN ] "
	default:
		return "[INFO ] "
	}
}
