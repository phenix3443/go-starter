package types

// AppInfoConfig 应用信息配置
type AppInfoConfig struct {
	Name        string `mapstructure:"name" validate:"required"`
	Version     string `mapstructure:"version" validate:"required"`
	Environment string `mapstructure:"environment" validate:"oneof=development staging production"`
	Debug       bool   `mapstructure:"debug"`
}
