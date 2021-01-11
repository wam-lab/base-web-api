package initialize

import (
	"github.com/spf13/viper"
	"github.com/wam-lab/base-web-api/internal/global"
	"path/filepath"
	"strings"
)

func Config(f string, mode string) {
	dir, fn := filepath.Split(f)
	ext := strings.TrimLeft(filepath.Ext(fn), ".")
	c := viper.New()
	c.SetConfigName("config")
	c.SetConfigType(ext)
	c.AddConfigPath(dir)
	c.AddConfigPath(".")
	err := c.ReadInConfig()
	if err != nil {
		panic(err)
	}
	c.Set("mode", mode)
	global.Config = c
}
