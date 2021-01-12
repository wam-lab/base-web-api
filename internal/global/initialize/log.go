package initialize

import (
	"github.com/wam-lab/base-web-api/common/conno"
	"github.com/wam-lab/base-web-api/internal/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func Log() {
	var logger *zap.Logger
	c := global.Config.GetStringMapString("log")
	mode := global.Config.GetString("mode")
	name := global.Config.GetString("Name")
	core := zapcore.NewCore(newZapEncoder(mode), newWriteSyncer(mode), zap.NewAtomicLevelAt(level(c["Level"])))
	field := zap.Fields(zap.String("APP", name))

	if mode == conno.DEV {
		caller := zap.AddCaller()
		dev := zap.Development()
		logger = zap.New(core, caller, dev, field)
	} else {
		logger = zap.New(core, field)
	}
	global.Log = logger
}

func newWriteSyncer(mode string) zapcore.WriteSyncer {
	//if mode == conno.DEV {
	//	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	//} else {
	//	logConfig := global.Config.Log
	//	now := time.Now()
	//	hook := lumberjack.Logger{
	//		Filename:   logConfig.Path + string(os.PathSeparator) + fmt.Sprintf("%04d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()) + ".log",
	//		MaxSize:    logConfig.MaxSize,
	//		MaxBackups: logConfig.MaxBackups,
	//		MaxAge:     logConfig.MaxAge,
	//		Compress:   logConfig.Compress,
	//	}
	//	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	//}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
}

func newZapEncoder(mode string) zapcore.Encoder {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "lineNum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	if mode == conno.PRO {
		return zapcore.NewJSONEncoder(encoderCfg)
	} else {
		return zapcore.NewConsoleEncoder(encoderCfg)
	}
}

func level(l string) zapcore.Level {
	switch strings.ToUpper(l) {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "DPANIC":
		return zapcore.DPanicLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
