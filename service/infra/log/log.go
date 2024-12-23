package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
	"time"
)

var sugar *zap.SugaredLogger

var ServiceName = "main"

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {

	pc, _, _, _ := runtime.Caller(7)
	f := runtime.FuncForPC(pc)

	enc.AppendString(f.Name())
}

func initLogPath(serviceName string) string {
	rootPath := GetPath()
	logPath := strings.Join([]string{rootPath, "logs"}, "/")

	if !IsDirExists(logPath) {
		MkdirFile(logPath)
	}

	//date := getDate()
	//outputLog := strings.Join([]string{serviceName + "_output.log", date}, "_")
	outputLog := strings.Join([]string{serviceName + "_output.log"}, "_")
	outputLogPath := strings.Join([]string{logPath, outputLog}, "/")

	return outputLogPath
}

func init() {
	outputLogPath := initLogPath(ServiceName)

	logger := NewLogger(outputLogPath, zapcore.DebugLevel, 200, 10, 14, true, ServiceName)
	sugar = logger.Sugar()
}

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

func InitLogger(serviceName string) {
	ServiceName = serviceName
	outputLogPath := initLogPath(ServiceName)

	logger := NewLogger(outputLogPath, zapcore.DebugLevel, 100, 10, 7, true, ServiceName)
	sugar = logger.Sugar()
}

/**
 * zapcore构造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		//Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		CallerKey:    "caller",
		MessageKey:   "msg",
		EncodeLevel:  zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:   zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeCaller: callerEncoder,
		EncodeName:   zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func Error(args ...interface{}) {

	sugar.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

func Panic(args ...interface{}) {
	sugar.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	sugar.Panicf(template, args...)
}

func getDate() string {

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}
