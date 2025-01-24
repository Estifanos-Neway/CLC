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

// Configuration values
type appConfig struct {
	Gemini struct {
		ApiKey           string `yaml:"apiKey" validate:"required"`
		Url              string `yaml:"url" validate:"required"`
		GenerationConfig struct {
			Temperature      float32 `yaml:"temperature"`
			TopK             int     `yaml:"topK"`
			TopP             float32 `yaml:"topP"`
			MaxOutputTokens  int     `yaml:"maxOutputTokens"`
			ResponseMimeType string  `yaml:"responseMimeType"`
		} `yaml:"generationConfig" validate:"required"`
	} `yaml:"gemini" validate:"required"`
}

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
