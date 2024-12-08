package config

import (
	"log"

	mylog "github.com/imniynaiy/ticket-system/internal/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var ConfigFile = pflag.StringP("config", "c", "./configs/config.yaml", "Set config file")

type config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Log      mylog.LogConfig
}

type ServerConfig struct {
	Port          string
	ReleaseMode   bool
	Cors          bool
	StaticPath    string `mapstructure:"static-path"`
	AuthSalt      string `mapstructure:"auth-salt"`
	JwtSigningKey string `mapstructure:"jwt-signing-key"`
}

type DatabaseConfig struct {
	Address  string
	Port     string
	Username string
	Password string
	DBName   string `mapstructure:"dbname"`
}

var GlobalConfig config

func ParseConfig() {
	viper.SetConfigFile(*ConfigFile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		log.Panicf("fatal error reading config file %s \n", err)
	}
	err := viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Panicf("fatal error parsing config file %s \n", err)
	}
}
