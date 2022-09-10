package mysql

import (
	"context"

	"gorm.io/gorm"
)

// SQLConn struct
type SQLConn struct {
	name string
}

// GetSQLConn 获取 sql 链接
func GetSQLConn(name string) *SQLConn {
	return &SQLConn{name}
}

// Master 获取主库链接
func (self *SQLConn) Master(ctx context.Context) *gorm.DB {
	if client, ok := mysqlClientMap.Load(self.name); ok {
		if v, ok := client.(*MySQLGroup); ok {
			return v.Master().WithContext(ctx)
		}
	}
	return nil
}

// Slave 获取从库链接
func (self *SQLConn) Slave(ctx context.Context) *gorm.DB {
	if client, ok := mysqlClientMap.Load(self.name); ok {
		if v, ok := client.(*MySQLGroup); ok {
			return v.Slave().WithContext(ctx)
		}
	}
	return nil
}
