package types

import (
	"fmt"
)

// DatabaseConfig 数据库连接配置
type DatabaseConfig struct {
	Host            string `mapstructure:"host" validate:"required"`
	Port            int    `mapstructure:"port" validate:"required,min=1,max=65535"`
	Username        string `mapstructure:"username" validate:"required"`
	Password        string `mapstructure:"password" validate:"required"`
	Database        string `mapstructure:"database" validate:"required"`
	Charset         string `mapstructure:"charset"`
	ParseTime       bool   `mapstructure:"parse_time"`
	Loc             string `mapstructure:"loc"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" validate:"min=1"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" validate:"min=1"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime" validate:"min=1"`
}

// MySQLConfig MySQL 特定配置
type MySQLConfig struct {
	DatabaseConfig `mapstructure:",squash"`
	SSLMode        string `mapstructure:"ssl_mode"`
}

// PostgresConfig PostgreSQL 特定配置
type PostgresConfig struct {
	DatabaseConfig `mapstructure:",squash"`
	SSLMode        string `mapstructure:"ssl_mode"`
	TimeZone       string `mapstructure:"timezone"`
}

// GetDSN 获取 MySQL DSN 字符串
func (c *MySQLConfig) GetDSN() string {
	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}
	if c.Loc == "" {
		c.Loc = "Local"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset, c.ParseTime, c.Loc)

	if c.SSLMode != "" {
		dsn += "&tls=" + c.SSLMode
	}

	return dsn
}

// GetDSN 获取 PostgreSQL DSN 字符串
func (c *PostgresConfig) GetDSN() string {
	if c.TimeZone == "" {
		c.TimeZone = "UTC"
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.Host, c.Port, c.Username, c.Password, c.Database)

	if c.SSLMode != "" {
		dsn += " sslmode=" + c.SSLMode
	}

	dsn += " TimeZone=" + c.TimeZone
	return dsn
}
