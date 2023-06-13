package config

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

func LoadLocale(config *AppConfig) {
	log.Debug("load locale")
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := bundle.LoadMessageFile(rootPath() + "/resources/message/en.json")
	if err != nil {
		log.Panic("Failed to load locale file ", err)
	}
	config.Localizer = i18n.NewLocalizer(bundle, language.English.String())
}
