package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GinMode           string `mapstructure:"GIN_MODE"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	DBMaxIdleConn     int    `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxOpenConn     int    `mapstructure:"DB_MAX_OPEN_CONN"`
	JWTTokenLifeSpan  int    `mapstructure:"TOKEN_HOUR_LIFESPAN"`
	JWTTokenSecretKey string `mapstructure:"JWT_TOKEN_SECRET"`
	Localizer         *i18n.Localizer
	Log               *log.Logger
}

func LoadConfig(activeProfile string) (config AppConfig, err error) {
	log.Debug("Load Config")
	path := rootPath() + "/env"
	viper.AddConfigPath(path)
	if len(activeProfile) > 0 {
		viper.SetConfigName(activeProfile)
	} else {
		viper.SetConfigName("app")
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config,", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	err = os.Setenv("JWT_TOKEN_SECRET", config.JWTTokenSecretKey)
	if err != nil {
		return
	}
	return
}

func rootPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../")
}

func Translate(messageID *string, localizer i18n.Localizer) string {
	localizeConf1g := i18n.LocalizeConfig{
		MessageID: *messageID,
	}
	LocalizedString, err := localizer.Localize(&localizeConf1g)
	if err != nil {
		log.Error("Error occurred while parsing string.", err)
	}
	return LocalizedString
}
