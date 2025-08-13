package factory_method

type Logger interface {
	Log(message string) string
}
type FileLogger struct {
}

type ConsoleLogger struct {
}

func (fl *FileLogger) Log(message string) string {
	return "FileLogger: " + message
}

func (cl *ConsoleLogger) Log(message string) string {
	return "ConsoleLogger: " + message
}

type LoggerFactory interface {
	CreateLogger() Logger
}

type FileLoggerFactory struct {
}

type ConsoleLoggerFactory struct {
}

func (flf *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

func (clf *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
