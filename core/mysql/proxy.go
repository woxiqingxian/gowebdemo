package mysql

import (
	"context"

	"gorm.io/gorm"
)

type SQL struct {
	name string
}

func InitSQL(name string) *SQL {
	return &SQL{name}
}

func (self *SQL) Master(ctx context.Context) *gorm.DB {
	if client, ok := mysqlClientMap.Load(self.name); ok {
		if v, ok := client.(*MySQLGroup); ok {
			return v.Master().WithContext(ctx)
		}
	}
	return nil
}

func (self *SQL) Slave(ctx context.Context) *gorm.DB {
	if client, ok := mysqlClientMap.Load(self.name); ok {
		if v, ok := client.(*MySQLGroup); ok {
			return v.Slave().WithContext(ctx)
		}
	}
	return nil
}
