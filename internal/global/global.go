package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *viper.Viper
	Log    *zap.Logger
	Db     *gorm.DB
)
