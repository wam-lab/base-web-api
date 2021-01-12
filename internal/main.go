package main

import (
	"flag"
	"github.com/wam-lab/base-web-api/internal/core"
	"github.com/wam-lab/base-web-api/internal/global/initialize"
)

var configFile = flag.String("f", "etc/config.yml", "the config file")
var mode = flag.String("m", "dev", "the development mode")

func main() {
	initialize.Config(*configFile, *mode)
	initialize.Log()
	initialize.Mysql()

	core.Run()
}
