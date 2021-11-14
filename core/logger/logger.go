package logger

import (
	"context"
	"fmt"
	"gowebdemo/core/config"
	"path"
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerMap sync.Map

var serverLogName string = "server"
var accessLogName string = "access"
var errorLogName string = "error"

func SetUp() {
	logNameList := []string{}

	defaultLogNameList := []string{serverLogName, accessLogName, errorLogName}
	logNameList = append(logNameList, defaultLogNameList...)
	logNameList = append(logNameList, config.ServerConfig.LogConfig.LogNameList...)

	for _, logName := range logNameList {
		logFilePath := path.Join(config.ServerConfig.LogConfig.LogDir, logName+".log")
		fmt.Println(logFilePath)
		logger := initLogger(logFilePath)
		loggerMap.Store(logName, logger)

	}
	return
}

func LoggerSync() {
	loggerMap.Range(func(key, value interface{}) bool {
		if logger, ok := value.(*zap.Logger); ok {
			logger.Sync()
		}
		return true
	})
}

func initLogger(logFilePath string) *zap.Logger {
	// WriterSyncer ：指定日志将写到哪里去
	writeSyncerList := []zapcore.WriteSyncer{}
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath,                              // 日志文件位置
		MaxSize:    config.ServerConfig.LogConfig.MaxSize,    // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxBackups: config.ServerConfig.LogConfig.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     config.ServerConfig.LogConfig.MaxAge,     // 保留旧文件的最大天数
		Compress:   config.ServerConfig.LogConfig.Compress,   // 是否压缩/归档旧文件
	})
	writeSyncerList = append(writeSyncerList, fileWriteSyncer)
	// // 输出到控制台
	// if config.ServerConfig.AppConfig.Debug {
	// 	consoleWriteSyncer := zapcore.AddSync(os.Stdout)
	// 	writeSyncerList = append(writeSyncerList, consoleWriteSyncer)
	// }
	writeSyncer := zapcore.NewMultiWriteSyncer(writeSyncerList...)

	// Encoder:编码器(如何写入日志)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// encoder := zapcore.NewConsoleEncoder(encoderConfig)
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Log Level：哪种级别的日志将被写入
	// zapcore.DebugLevel  zapcore.InfoLevel  zapcore.ErrorLevel
	zapcoreLogLevelMap := map[string]zapcore.LevelEnabler{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	logLevel := zapcoreLogLevelMap[config.ServerConfig.LogConfig.Level]
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger
}

func Log(ctx context.Context, nameList ...string) *zap.SugaredLogger {
	logName := errorLogName
	if len(nameList) == 1 {
		logName = nameList[0]
	}
	return loadLogWithTraceId(ctx, logName).Sugar()
}

func ServerLog() *zap.Logger {
	return loadLog(serverLogName)
}

func MysqlLog() *zap.Logger {
	return loadLog(errorLogName)
}

func loadLogWithTraceId(ctx context.Context, logName string) *zap.Logger {
	logger := loadLog(logName)
	traceId, _ := ctx.Value("traceId").(string)
	return logger.With(zap.String("traceId", traceId))
}

func loadLog(logName string) *zap.Logger {
	var logger *zap.Logger
	if val, ok := loggerMap.Load(logName); ok {
		logger, _ = val.(*zap.Logger)
	}
	return logger
}
