package initialize

import (
	"github.com/wam-lab/base-web-api/internal/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

// use single database
func Mysql() {
	c := global.Config
	db, err := gorm.Open(mysql.Open(c.GetString("Mysql.DataSource")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.GetString("Mysql.Prefix"),
			SingularTable: false,
		},
		// 使用
		Logger: Default.LogMode(logger.Info),
	})
	if err != nil {
		global.Log.Panic("Mysql connect failed", zap.Any("err", err))
		os.Exit(1)
	}

	if sqlDb, err := db.DB(); err != nil {
		panic(err)
	} else {
		sqlDb.SetMaxIdleConns(c.GetInt("Mysql.MaxIdle"))
		sqlDb.SetMaxOpenConns(c.GetInt("Mysql.MaxOpen"))
		err := sqlDb.Ping()
		if err != nil {
			global.Log.Panic("Ping mysql failed", zap.Any("err", err))
			os.Exit(1)
		}
	}

	global.Db = db
}
