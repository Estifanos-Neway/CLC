package config

import (
	_ "embed"
)

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
		SystemInstruction string `yaml:"systemInstruction"`
	} `yaml:"gemini" validate:"required"`
	Environment string `yaml:"environment" default:"dev" validation:"oneof=dev prod"`
}

const (
	EnvironmentDev  string = "dev"
	EnvironmentProd string = "prod"
)
