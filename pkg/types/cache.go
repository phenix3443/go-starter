package types

import (
	"fmt"
)

// RedisConfig Redis 配置.
type RedisConfig struct {
	Host         string `mapstructure:"host" validate:"required"`
	Port         int    `mapstructure:"port" validate:"required,min=1,max=65535"`
	Password     string `mapstructure:"password"`
	Database     int    `mapstructure:"database" validate:"min=0,max=15"`
	PoolSize     int    `mapstructure:"pool_size" validate:"min=1"`
	MinIdleConns int    `mapstructure:"min_idle_conns" validate:"min=0"`
}

// GetRedisDSN 获取 Redis URL 字符串
func (c *RedisConfig) GetRedisDSN() string {
	if c.Password != "" {
		return fmt.Sprintf("redis://:%s@%s:%d/%d", c.Password, c.Host, c.Port, c.Database)
	}
	return fmt.Sprintf("redis://%s:%d/%d", c.Host, c.Port, c.Database)
}
