package logs

import (
	_ "encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	Log *zap.Logger
)

const (
	DESC    = "desc"
	ERR     = "err"
	OTHER   = "other"
	INFO    = "[INFO]: "
	ERROR   = "[ERROR]: "
	WARNING = "[WARN]: "
)

// logpath 日志文件路径
// loglevel 日志级别
func InitLogger(logpath string, loglevel string) {

	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = myTimeEncode
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)
	Log = zap.New(core)
	Log.Info("[INFO]: ", zap.String("desc", "DefaultLogger init success"))
}

func myTimeEncode(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 封装错误
func MyError(desc string, err interface{}, others ...interface{}) {
	Log.Error(ERROR, zap.String(DESC, desc), zap.Any(ERR, err), zap.String(OTHER, fmt.Sprint(others)))
}

// 封装提示
func MyInfo(desc string, others ...interface{}) {
	Log.Info(INFO, zap.String(DESC, desc), zap.String(OTHER, fmt.Sprint(others)))
}
