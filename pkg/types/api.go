package types

// CORSConfig CORS 配置
type CORSConfig struct {
	AllowOrigin      string `mapstructure:"allow_origin" validate:"required"`
	AllowMethods     string `mapstructure:"allow_methods" validate:"required"`
	AllowHeaders     string `mapstructure:"allow_headers" validate:"required"`
	ExposeHeaders    string `mapstructure:"expose_headers"`
	MaxAge           int    `mapstructure:"max_age" validate:"min=0"`
	AllowCredentials bool   `mapstructure:"allow_credentials"`
}

// RateLimitConfig 限流配置
type RateLimitConfig struct {
	Enabled     bool `mapstructure:"enabled"`
	MaxRequests int  `mapstructure:"max_requests" validate:"min=1"`
	WindowSize  int  `mapstructure:"window_size" validate:"min=1"`
	BurstSize   int  `mapstructure:"burst_size" validate:"min=1"`
}

// APIConfig API 配置
type APIConfig struct {
	Version      string          `mapstructure:"version" validate:"required"`
	Address      string          `mapstructure:"address" validate:"required"`
	CORS         CORSConfig      `mapstructure:"cors" validate:"required"`
	RateLimit    RateLimitConfig `mapstructure:"rate_limit"`
	ReadTimeout  int             `mapstructure:"read_timeout" validate:"min=1"`
	WriteTimeout int             `mapstructure:"write_timeout" validate:"min=1"`
	IdleTimeout  int             `mapstructure:"idle_timeout" validate:"min=1"`
}
