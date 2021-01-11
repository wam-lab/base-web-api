package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Config *viper.Viper
	Db     *gorm.DB
)
