package configs

import (
	_ "embed"
	"flag"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	defaultConfigFile = "./configs/default_configs.toml"
	config            = new(Config)

	//go:embed default_configs.toml
	defaultConfigs []byte
)

type Config struct {
	Server struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	} `toml:"server"`

	LogConfig struct {
		MaxSizeMB      int    `toml:"maxSizeMb"`
		MaxBackupSize  int    `toml:"maxBackupSize"`
		MaxAge         int    `toml:"maxAge"`
		AccessLogPath  string `toml:"accessLogPath"`
		BackendLogPath string `toml:"backendLogPath"`
	} `toml:"logConfig"`

	Language struct {
		Local string `toml:"local"`
	} `toml:"language"`
}

func Get() Config {
	if config.Server.Host == "" {
		config.Server.Host = "127.0.0.1"
	}
	if config.Server.Port == 0 {
		config.Server.Port = 8080
	}
	return *config
}

// TODO 使用配置中心统一配置
func init() {
	configFile := flag.String("config", defaultConfigFile, "config file path")
	flag.Parse()

	viper.SetConfigType("toml")
	viper.SetConfigFile(*configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}
