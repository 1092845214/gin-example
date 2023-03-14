package config

import (
	"github.com/spf13/viper"
	"github.com/yangkaiyue/gin-exp/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

func InitLogger() *zap.SugaredLogger {

	encoder := getEncoder()
	writeSyncer := getWriteSyncer()

	logMode := zapcore.InfoLevel
	if viper.GetBool("server.debug") {
		logMode = zapcore.DebugLevel
	}

	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, os.Stdout), logMode)
	return zap.New(core, zap.AddCaller()).
		WithOptions(zap.AddCallerSkip(2)).
		Sugar()
}

func getEncoder() zapcore.Encoder {

	// 默认格式
	cfg := zap.NewProductionEncoderConfig()

	// 修改默认格式
	cfg.TimeKey = "time"
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.In(global.CstZone).Format(global.DateTimeFormat))
	}

	//return zapcore.NewJSONEncoder(cfg)
	return zapcore.NewConsoleEncoder(cfg)
}

func getWriteSyncer() zapcore.WriteSyncer {

	logFile := path.Join(global.ProjectPath, "log", time.Now().In(global.CstZone).Format(global.DateOnlyFormat))

	// 日志切割使用 lumberjack
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   logFile + ".log",
		MaxSize:    viper.GetInt("log.max_size"), // MB
		MaxAge:     viper.GetInt("log.max_backup"),
		MaxBackups: viper.GetInt("log.max_age"),
	}

	return zapcore.AddSync(lumberjackSyncer)
}
