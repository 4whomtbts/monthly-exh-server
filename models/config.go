package models

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const DEFAULT_LOG = "logs/monthlyExhibition.log"

type Config struct {
	LogSettings  LogSettings
	AwsSettings AwsSettings

}

type LogSettings struct {
	LogConfig *zap.Config
	LogFilePath string
	EnableConsole          bool
	ConsoleLevel           string
	ConsoleJson            *bool
	EnableFile             bool
	FileLevel              string
	FileJson               *bool
	FileLocation           string
	EnableWebhookDebugging bool
	EnableDiagnostics      *bool
}


func (s *LogSettings) SetDefaults() *LogSettings {


	s.LogConfig = &zap.Config{
		Encoding : "json",
		Level : zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths : []string{DEFAULT_LOG},
		ErrorOutputPaths : []string{DEFAULT_LOG},

		EncoderConfig : zapcore.EncoderConfig{
			MessageKey : "msg",

			LevelKey : "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,

			TimeKey  : "time",
			EncodeTime : zapcore.ISO8601TimeEncoder,

			CallerKey : "called_from",
			EncodeCaller : zapcore.ShortCallerEncoder,
		},
	}

	return s;
}


type AwsSettings struct {
	Region string

}

func (s *AwsSettings) SetDefaults() *AwsSettings {
	s.Region = "ap-northeast-2"
	return s;
}


func (o *Config) SetDefaults() {
	o.LogSettings.SetDefaults()
	o.AwsSettings.SetDefaults()
}