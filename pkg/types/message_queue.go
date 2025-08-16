package types

// KafkaConfig Kafka 配置
type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers" validate:"required,min=1"`
	Topic   string   `mapstructure:"topic" validate:"required"`
}
