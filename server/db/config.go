/**
配置数据源信息
*/
package db

import (
	"github.com/ChinasoftNobody/gochat/server/common"
	"github.com/ChinasoftNobody/gochat/server/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wonderivan/logger"
)

var DB *gorm.DB

var DbCreate = []interface{}{
	&dto.CommonMessage{},
	&dto.ClientConnectDto{},
}

/**
初始化数据库实例
*/
func init() {
	var err error
	DB, err = gorm.Open("mysql", common.DbUrl)
	if err != nil {
		panic(err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(200)
	DB.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	logger.Info("gorm初始化成功")
	createTable()
}

/**
创建表信息
*/
func createTable() {
	for _, table := range DbCreate {
		if !DB.HasTable(table) {
			DB.CreateTable(table)
		}
	}
}
