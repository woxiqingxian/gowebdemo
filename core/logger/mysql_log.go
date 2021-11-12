package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

func NewMysqlLogger(config gormLogger.Config, writerLog zap.Logger) gormLogger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	return &logger{
		Config:       config,
		writer:       writerLog,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type logger struct {
	gormLogger.Config
	writer                              zap.Logger
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (self *logger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newlogger := *self
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (self logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if self.LogLevel >= gormLogger.Info {
		self.writer.Info(fmt.Sprintf(self.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
	}
}

// Warn print warn messages
func (self logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if self.LogLevel >= gormLogger.Warn {
		self.writer.Warn(fmt.Sprintf(self.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
	}
}

// Error print error messages
func (self logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if self.LogLevel >= gormLogger.Error {
		self.writer.Error(fmt.Sprintf(self.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
	}
}

// // Trace print sql message
// func (self logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
// 	if self.LogLevel <= gormLogger.Silent {
// 		return
// 	}
//
// 	elapsed := time.Since(begin)
// 	switch {
// 	case err != nil && self.LogLevel >= gormLogger.Error && (!errors.Is(err, gormLogger.ErrRecordNotFound) || !self.IgnoreRecordNotFoundError):
// 		sql, rows := fc()
// 		if rows == -1 {
// 			self.writer.Error(fmt.Sprintf(self.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
// 		} else {
// 			self.writer.Error(fmt.Sprintf(self.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
// 		}
// 	case elapsed > self.SlowThreshold && self.SlowThreshold != 0 && self.LogLevel >= gormLogger.Warn:
// 		sql, rows := fc()
// 		slowLog := fmt.Sprintf("SLOW SQL >= %v", self.SlowThreshold)
// 		if rows == -1 {
// 			self.writer.Warn(fmt.Sprintf(self.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql))
// 		} else {
// 			self.writer.Warn(fmt.Sprintf(self.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
// 		}
// 	case self.LogLevel == gormLogger.Info:
// 		sql, rows := fc()
// 		if rows == -1 {
// 			self.writer.Info(fmt.Sprintf(self.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql))
// 		} else {
// 			self.writer.Info(fmt.Sprintf(self.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
// 		}
// 	}
// }

// Trace print sql message
func (self logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	traceId, _ := ctx.Value("traceId").(string)
	writer := self.writer.With(zap.String("traceId", traceId))
	if self.LogLevel <= gormLogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil && self.LogLevel >= gormLogger.Error && (!errors.Is(err, gormLogger.ErrRecordNotFound) || !self.IgnoreRecordNotFoundError):
		// 错误日志
		writer.Error(fmt.Sprintf("[Line:%s] [Err:%s] [cost:%.3fms] [rows:%d] [sql:%s]", utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
	case elapsed > self.SlowThreshold && self.SlowThreshold != 0 && self.LogLevel >= gormLogger.Warn:
		// 慢日志
		slowLog := fmt.Sprintf("SLOW SQL >= %v", self.SlowThreshold)
		writer.Warn(fmt.Sprintf("[Line:%s] [Err:%s] [cost:%.3fms] [rows:%d] [:%s]", utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
	case self.LogLevel == gormLogger.Info:
		// 日志记录
		writer.Info(fmt.Sprintf("[Line:%s] [cost:%.3fms] [rows:%d] [%s]", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
	}
	return

}
