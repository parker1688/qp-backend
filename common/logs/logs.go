package logs

import (
	"bootpkg/common/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func NewLog(filename string) *zap.SugaredLogger {
	if len(filename) == 0 {
		filename = "boot.log"
	}
	hook := lumberjack.Logger{
		Filename: filename,
		MaxSize:  1024,
		MaxAge:   8,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "message",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		CallerKey:      "caller",
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	multiWrite := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	if global.CONFIG.General.ENV == "Debug" {
		multiWrite = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		multiWrite,
		atomicLevel,
	)
	return zap.New(core, zap.AddCaller()).Sugar()
}
