package conf

import (
	"errors"
	"log/slog"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	App AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Port      string `mapstructure:"port"`
	MongoURI  string `mapstructure:"mongoUri"`
	DBname    string `mapstructure:"dbName"`
	UserColl  string `mapstructure:"userColl"`
	JWTSecret string `mapstructure:"jwtSecret"`
	MD5Secret string `mapstructure:"md5Secret"`
}

func ParseConfig() error {
	viper.SetConfigName("leaf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("viper read config error")
	}
	if err := viper.Unmarshal(Conf); err != nil {
		return errors.New("viper unmarshal error")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			slog.Error("viper reunmarshal error", "error message", err)
			return
		}
	})
	return nil
}
