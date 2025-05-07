package config

import (
	_ "embed"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	yaml "gopkg.in/yaml.v3"
)

var (
	//go:embed config.yaml
	configData []byte
	v          *validator.Validate
)

var AppConfig *appConfig

func Load() error {
	config := &appConfig{}

	if err := defaults.Set(config); err != nil {
		return err
	}

	if err := yaml.Unmarshal(configData, &config); err != nil {
		return err
	}

	v = validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(config); err != nil {
		return err
	}

	AppConfig = config
	return nil
}
