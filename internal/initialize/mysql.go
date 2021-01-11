package initialize

import (
	"github.com/wam-lab/base-web-api/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Mysql() {
	c := global.Config
	dsn := c.GetString("Mysql.DataSource")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.GetString("Mysql.Prefix"),
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err)
	}

	if sqlDb, err := db.DB(); err != nil {
		panic(err)
	} else {
		sqlDb.SetMaxIdleConns(c.GetInt("Mysql.MaxIdle"))
		sqlDb.SetMaxOpenConns(c.GetInt("Mysql.MaxOpen"))
		err := sqlDb.Ping()
		if err != nil {
			panic(err)
		}
	}

	global.Db = db
}
