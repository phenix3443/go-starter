package types

// LogConfig 日志配置
type LogConfig struct {
	Level  string `mapstructure:"level" validate:"oneof=debug info warn error"`
	Format string `mapstructure:"format" validate:"oneof=json text"`
	Output string `mapstructure:"output" validate:"oneof=stdout stderr file"`
	File   string `mapstructure:"file"`
}
