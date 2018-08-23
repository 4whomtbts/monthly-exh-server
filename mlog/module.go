package mlog

import (
	"go.uber.org/zap"
	"fmt"
	"go.uber.org/zap/zapcore"
)
type Field = zapcore.Field

type LoggerConfiguration struct {
	EnableConsole bool
	ConsoleJson   bool
	ConsoleLevel  string
	EnableFile    bool
	FileJson      bool
	FileLevel     string
	FileLocation  string
}

type Logger struct {
	zap *zap.Logger
	consoleLevel zap.AtomicLevel
	fileLevel zap.AtomicLevel
}

var globalLogger *Logger

func InitGlobalLogger(logger *Logger) {
	globalLogger = logger
	Debug = globalLogger.Debug
	Info = globalLogger.Info
	Warn = globalLogger.Warn
	Error = globalLogger.Error
	Critical = globalLogger.Critical
}

func NewLogger(config zap.Config) *Logger {

	//logger,_ := zap.NewDevelopment()

	fmt.Println("NEwLogger 콘피그 ",config)
	l0 := &Logger{
		zap : nil,
	}
	//config.EncoderConfig.EncodeLevel
	logger , err := config.Build()




	if err != nil {
		fmt.Println(err)
	}

	l0.zap = logger
	return l0

	}

func (l *Logger) Info(message string, fields ...Field) {
	l.zap.Info(message, fields...)
}
func (l *Logger) Warn(message string, fields ...Field) {
	l.zap.Warn(message, fields...)
}
func (l *Logger) Debug(message string, fields ...Field) {
	l.zap.Debug(message, fields...)
}

func (l *Logger) Error(message string, fields ...Field) {
	l.zap.Error(message, fields...)
}

func (l *Logger) Critical(message string, fields ...Field) {
	l.zap.Error(message, fields...)
}


type LogFunc func(string, ...Field)

var Debug LogFunc = defaultDebugLog
var Info LogFunc = defaultInfoLog
var Warn LogFunc = defaultWarnLog
var Error LogFunc = defaultErrorLog
var Critical LogFunc = defaultCriticalLog
