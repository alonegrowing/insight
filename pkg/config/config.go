package config

import (
	"os"

	"github.com/alonegrowing/insight/pkg/basic/util"
	"github.com/alonegrowing/insight/pkg/treasure/log"
	"github.com/alonegrowing/insight/pkg/treasure/redis"
	"github.com/alonegrowing/insight/pkg/treasure/sql"
	"github.com/alonegrowing/insight/pkg/treasure/tomlconfig"
)

var (
	err           error
	conf          = "./conf/prod/config.toml"
	ServiceConfig Config
)

func init() {
	util.PanicIfError(tomlconfig.ParseTomlConfig(conf, &ServiceConfig))
	InitLoggerConfig(ServiceConfig.Logger)
}

type Service struct {
	WEBPort int64 `toml:"web_port"`
	RPCPort int64 `toml:"rpc_port"`
}

type LoggerConfig struct {
	Level   log.Level `toml:"level"`
	LogPath string    `toml:"logpath"`
}

type Config struct {
	ServiceName string            `toml:"service_name"`
	Service     Service           `toml:"service"`
	Logger      LoggerConfig      `toml:"log"`
	Database    []sql.GroupConfig `toml:"database"`
	Redis       []redis.Config    `toml:"redis"`
}

func InitLoggerConfig(conf LoggerConfig) {
	file, err := os.OpenFile(conf.LogPath, os.O_CREATE|os.O_WRONLY, 0666)
	util.PanicIfError(err)
	log.SetOutput(file)
	log.SetLevel(log.InfoLevel)
}
