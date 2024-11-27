package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type Config struct {
	Name             string `mapstructure:"name"`
	Port             int    `mapstructure:"port"`
	Version          string `mapstructure:"version"`
	StartTime        string `mapstructure:"start_time"`
	MachineID        int    `mapstructure:"machine_id"`
	Mode             string `mapstructure:"mode"`
	TranslatorLocale string `mapstructure:"translator_locale"`

	*LoggerConfig `mapstructure:"logger"`
	*MysqlConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
}

type LoggerConfig struct {
	LogFile string `mapstructure:"log_file"`
	MaxSize int    `mapstructure:"max_size"`
	MaxAge  int    `mapstructure:"max_age"`
	Level   string `mapstructure:"level"`
}

type MysqlConfig struct {
	UserName string `mapstructure:"username"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

var Conf = new(Config)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("Unable to decode config file to struct, %v", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("Unable to decode config file to struct, %v", err)
		}
	})
	return
}
