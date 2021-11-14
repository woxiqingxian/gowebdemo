package mysql

import (
	"fmt"
	"gowebdemo/core/config"
	"gowebdemo/core/logger"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	goSqlDriverMysql "github.com/go-sql-driver/mysql"
	gormDriverMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	mysqlClientMap sync.Map
)

type MySQLGroup struct {
	name      string
	master    *gorm.DB
	slaveList []*gorm.DB
	next      uint64
	total     uint64
}

// Master返回master实例
func (self *MySQLGroup) Master() *gorm.DB {
	return self.master
}

// Slave返回一个slave实例，使用轮转算法
func (self *MySQLGroup) Slave() *gorm.DB {
	if self.total == 0 {
		return self.master
	}
	next := atomic.AddUint64(&self.next, 1)
	return self.slaveList[next%self.total]
}

// Instance函数如果isMaster是true， 返回master实例，否则返回slave实例
func (self *MySQLGroup) Instance(isMaster bool) *gorm.DB {
	if isMaster {
		return self.Master()
	}
	return self.Slave()
}

func SetUp() {
	for _, mysqlConfig := range config.ServerConfig.MysqlConfigList {
		if _, ok := mysqlClientMap.Load(mysqlConfig.Name); ok {
			continue
		}

		mysqlGroup, err := newMySQLGroup(mysqlConfig)
		if err != nil {
			logger.ServerLog().Panic(fmt.Sprintf("SetUp mysql Error %s", err))
		}
		mysqlClientMap.LoadOrStore(mysqlConfig.Name, mysqlGroup)
		logger.ServerLog().Info(fmt.Sprintf("mysql %s setup success", mysqlConfig.Name))
	}
}

func newMySQLGroup(mysqlConfig config.MysqlConf) (*MySQLGroup, error) {
	var err error
	mysqlGroup := MySQLGroup{name: mysqlConfig.Name}

	mysqlGroup.master, err = openConn(mysqlConfig.Master, mysqlConfig.LogLevel, mysqlConfig.SlowThreshold)
	if err != nil {
		return nil, err
	}
	mysqlGroup.slaveList = make([]*gorm.DB, 0, len(mysqlConfig.SlaveList))
	mysqlGroup.total = 0
	for _, slave := range mysqlConfig.SlaveList {
		c, err := openConn(slave, mysqlConfig.LogLevel, mysqlConfig.SlowThreshold)
		if err != nil {
			return nil, err
		}
		mysqlGroup.slaveList = append(mysqlGroup.slaveList, c)
		mysqlGroup.total++

	}
	return &mysqlGroup, nil
}

func openConn(dsn string, logLevel string, slowThreshold int) (*gorm.DB, error) {
	var err error

	// dsn 解析
	dsnParse, err := goSqlDriverMysql.ParseDSN(dsn)

	if err != nil {
		return nil, err
	}
	maxIdle, _ := strconv.Atoi(dsnParse.Params["max_idle"])
	if maxIdle == 0 {
		maxIdle = 15
	}
	maxActive, _ := strconv.Atoi(dsnParse.Params["max_active"])
	if maxActive == 0 {
		maxActive = 20
	}
	maxLifetime, _ := strconv.Atoi(dsnParse.Params["max_lifetime_sec"])
	if maxLifetime == 0 {
		maxLifetime = 1800
	}
	delete(dsnParse.Params, "max_idle")
	delete(dsnParse.Params, "max_active")
	delete(dsnParse.Params, "max_lifetime_sec")
	dsn = dsnParse.FormatDSN()

	// 获取日志级别 Info：4 Warn：3 Silent：1 Error：2
	logLevelMap := map[string]gormLogger.LogLevel{
		"slient": gormLogger.Silent,
		"error":  gormLogger.Error,
		"warn":   gormLogger.Warn,
		"info":   gormLogger.Info,
	}
	newLogger := logger.NewMysqlLogger(
		gormLogger.Config{
			SlowThreshold:             time.Duration(slowThreshold) * time.Millisecond, // 慢SQL 阈值
			LogLevel:                  logLevelMap[logLevel],                           // Log level
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
		logger.MysqlLog(),
	)

	//打开连接
	dbConn, err := gorm.Open(gormDriverMysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	currentDb, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	//连接池设置
	currentDb.SetMaxIdleConns(maxIdle)
	currentDb.SetMaxOpenConns(maxActive)
	currentDb.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)

	return dbConn, err
}
