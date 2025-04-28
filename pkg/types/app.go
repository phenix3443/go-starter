package types

type KafkaConfig struct {
	Brokers []string `mapstructure:"brokers"`
	Topic   string   `mapstructure:"topic"`
}

type CORSConfig struct {
	AllowOrigin      string `mapstructure:"allow_origin"`
	AllowMethods     string `mapstructure:"allow_methods"`
	AllowHeaders     string `mapstructure:"allow_headers"`
	ExposeHeaders    string `mapstructure:"expose_headers"`
	MaxAge           int    `mapstructure:"max_age"`
	AllowCredentials bool   `mapstructure:"allow_credentials"`
}

type RateLimitConfig struct {
	Enabled     bool `mapstructure:"enabled"`
	MaxRequests int  `mapstructure:"max_requests"`
	WindowSize  int  `mapstructure:"window_size"`
	BurstSize   int  `mapstructure:"burst_size"`
}

type APIConfig struct {
	Version   string          `mapstructure:"version"`
	Address   string          `mapstructure:"address"`
	CORS      CORSConfig      `mapstructure:"cors"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
}
type AppConfig struct {
	MySQL    string       `mapstructure:"mysql"`
	Postgres string       `mapstructure:"postgres"`
	Redis    string       `mapstructure:"redis"`
	Kafka    *KafkaConfig `mapstructure:"kafka"`
	API      *APIConfig   `mapstructure:"api"`
}
