package logger

import "go.uber.org/zap"

type Logger struct {
	Logger *zap.Logger
}

func SetLogger() *Logger {
	l, err := zap.NewDevelopment()

	if err != nil {
		panic(err)
	}
	return &Logger{
		Logger: l,
	}
}

func init() {
	Logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(Logger)
}
