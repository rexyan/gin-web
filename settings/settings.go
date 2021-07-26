package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Config = new(AppConfig)

type AppConfig struct {
	*ServerConfig `mapstructure:"server"`
	*LoggerConfig `mapstructure:"logger"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
}

type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type RedisConfig struct {
	DB       int    `mapstructure:"db"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"poolSize"`
}

type ServerConfig struct {
	Mode   string `mapstructure:"mode"`
	Name   string `mapstructure:"name"`
	Port   int    `mapstructure:"port"`
	Secret string `mapstructure:"secret"`
}

func Init() (err error) {
	// viper.SetConfigFile("完整路径 + 完整名称")
	viper.SetConfigName("config") // 根据文件名查找，而不是文件名+后缀
	// viper.SetConfigType("yaml") 从远程读取才需要指定，例如 etcd
	viper.AddConfigPath("./")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper read config file error: %v\n", err)
		return err
	}

	// 配置转为结构体
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper unmarshal error: %v\n", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("viper config change event!")
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper unmarshal error: %v\n", err)
		}
	})
	return
}
