package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/phenix3443/go-starter/pkg/types"
	"github.com/spf13/viper"
)

// Loader 配置加载器
type Loader struct {
	configPath string
	configName string
	configType string
}

// NewLoader 创建新的配置加载器
func NewLoader() *Loader {
	return &Loader{
		configName: "config",
		configType: "yaml",
	}
}

// SetConfigPath 设置配置文件路径
func (l *Loader) SetConfigPath(path string) *Loader {
	l.configPath = path
	return l
}

// SetConfigName 设置配置文件名
func (l *Loader) SetConfigName(name string) *Loader {
	l.configName = name
	return l
}

// SetConfigType 设置配置文件类型
func (l *Loader) SetConfigType(configType string) *Loader {
	l.configType = configType
	return l
}

// Load 加载配置
func (l *Loader) Load() (*types.AppConfig, error) {
	// 设置配置文件路径
	if l.configPath != "" {
		viper.SetConfigFile(l.configPath)
	} else {
		viper.SetConfigName(l.configName)
		viper.SetConfigType(l.configType)

		// 添加配置文件搜索路径
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("./app/your-app/config")
		viper.AddConfigPath("./app/your-app")
		viper.AddConfigPath("/etc/go-starter")
		viper.AddConfigPath("$HOME/.go-starter")
	}

	// 设置环境变量前缀
	viper.SetEnvPrefix("GO_STARTER")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// 设置默认值
	l.setDefaults()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		log.Warn("Config file not found, using defaults and environment variables")
	} else {
		log.Info("Using config file", "path", viper.ConfigFileUsed())
	}

	// 解析配置
	var config types.AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// 验证配置
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	// 应用环境变量覆盖
	l.applyEnvironmentOverrides(&config)

	return &config, nil
}

// setDefaults 设置默认配置值
func (l *Loader) setDefaults() {
	// 应用默认值
	viper.SetDefault("app.name", "go-starter")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("app.debug", true)

	// 日志默认值
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")
	viper.SetDefault("log.output", "stdout")

	// API 默认值
	viper.SetDefault("api.version", "v1")
	viper.SetDefault("api.address", ":8080")
	viper.SetDefault("api.read_timeout", 30)
	viper.SetDefault("api.write_timeout", 30)
	viper.SetDefault("api.idle_timeout", 60)

	// CORS 默认值
	viper.SetDefault("api.cors.allow_origin", "*")
	viper.SetDefault("api.cors.allow_methods", "GET, POST, PUT, DELETE, OPTIONS")
	viper.SetDefault("api.cors.allow_headers", "Content-Type, Authorization")
	viper.SetDefault("api.cors.max_age", 86400)
	viper.SetDefault("api.cors.allow_credentials", true)

	// 限流默认值
	viper.SetDefault("api.rate_limit.enabled", true)
	viper.SetDefault("api.rate_limit.max_requests", 100)
	viper.SetDefault("api.rate_limit.window_size", 60)
	viper.SetDefault("api.rate_limit.burst_size", 10)

	// 数据库连接池默认值
	viper.SetDefault("mysql.max_open_conns", 100)
	viper.SetDefault("mysql.max_idle_conns", 10)
	viper.SetDefault("mysql.conn_max_lifetime", 3600)
	viper.SetDefault("mysql.charset", "utf8mb4")
	viper.SetDefault("mysql.parse_time", true)
	viper.SetDefault("mysql.loc", "Local")

	viper.SetDefault("postgres.max_open_conns", 100)
	viper.SetDefault("postgres.max_idle_conns", 10)
	viper.SetDefault("postgres.conn_max_lifetime", 3600)
	viper.SetDefault("postgres.timezone", "UTC")

	// Redis 默认值
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.database", 0)
	viper.SetDefault("redis.pool_size", 10)
	viper.SetDefault("redis.min_idle_conns", 5)
}

// applyEnvironmentOverrides 应用环境变量覆盖
func (l *Loader) applyEnvironmentOverrides(config *types.AppConfig) {
	// 应用环境变量覆盖
	if env := os.Getenv("GO_STARTER_APP_ENVIRONMENT"); env != "" {
		config.App.Environment = env
	}

	if debug := os.Getenv("GO_STARTER_APP_DEBUG"); debug != "" {
		config.App.Debug = debug == "true"
	}

	if logLevel := os.Getenv("GO_STARTER_LOG_LEVEL"); logLevel != "" {
		config.Log.Level = logLevel
	}

	if apiAddr := os.Getenv("GO_STARTER_API_ADDRESS"); apiAddr != "" {
		config.API.Address = apiAddr
	}

	// 数据库环境变量覆盖
	if mysqlDSN := os.Getenv("GO_STARTER_MYSQL_DSN"); mysqlDSN != "" {
		config.MySQL = mysqlDSN
	}

	if postgresDSN := os.Getenv("GO_STARTER_POSTGRES_DSN"); postgresDSN != "" {
		config.Postgres = postgresDSN
	}

	// Redis 环境变量覆盖
	if redisURL := os.Getenv("GO_STARTER_REDIS_URL"); redisURL != "" {
		config.Redis = redisURL
	}
}

// LoadFromFile 从文件加载配置
func LoadFromFile(configPath string) (*types.AppConfig, error) {
	return NewLoader().SetConfigPath(configPath).Load()
}

// LoadFromEnv 从环境变量加载配置
func LoadFromEnv() (*types.AppConfig, error) {
	return NewLoader().Load()
}

// CreateConfigFile 创建配置文件
func CreateConfigFile(configPath string, config *types.AppConfig) error {
	// 确保目录存在
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// 创建配置文件
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	// 写入配置内容
	viper.SetConfigFile(configPath)
	viper.Set("app", config.App)
	viper.Set("mysql", config.MySQL)
	viper.Set("postgres", config.Postgres)
	viper.Set("redis", config.Redis)
	viper.Set("kafka", config.Kafka)
	viper.Set("api", config.API)
	viper.Set("log", config.Log)

	return viper.WriteConfig()
}
