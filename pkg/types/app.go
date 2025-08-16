package types

import (
	"fmt"
)

// AppConfig 应用主配置
type AppConfig struct {
	App      AppInfoConfig `mapstructure:"app" validate:"required"`
	MySQL    string        `mapstructure:"mysql"`
	Postgres string        `mapstructure:"postgres"`
	Redis    string        `mapstructure:"redis" validate:"required"`
	Kafka    *KafkaConfig  `mapstructure:"kafka"`
	API      *APIConfig    `mapstructure:"api" validate:"required"`
	Log      LogConfig     `mapstructure:"log" validate:"required"`
}

// Validate 验证配置
func (c *AppConfig) Validate() error {
	if c.App.Environment == "production" && c.App.Debug {
		return fmt.Errorf("debug mode should not be enabled in production environment")
	}

	if c.API != nil && c.API.ReadTimeout <= 0 {
		return fmt.Errorf("API read timeout must be greater than 0")
	}

	if c.API != nil && c.API.WriteTimeout <= 0 {
		return fmt.Errorf("API write timeout must be greater than 0")
	}

	return nil
}

// IsProduction 检查是否为生产环境
func (c *AppConfig) IsProduction() bool {
	return c.App.Environment == "production"
}

// IsDevelopment 检查是否为开发环境
func (c *AppConfig) IsDevelopment() bool {
	return c.App.Environment == "development"
}

// IsStaging 检查是否为预发布环境
func (c *AppConfig) IsStaging() bool {
	return c.App.Environment == "staging"
}
