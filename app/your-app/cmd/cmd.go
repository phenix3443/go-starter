package cmd

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"

	"github.com/phenix3443/go-starter/pkg/config"
	"github.com/phenix3443/go-starter/pkg/types"
)

var appConfig *types.AppConfig

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file path (optional, will search in common locations if not specified)")
	rootCmd.PersistentFlags().StringP("env", "e", "", "environment (development, staging, production)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug mode")
}

func initConfig() {
	// 获取配置文件路径
	configFile, _ := rootCmd.PersistentFlags().GetString("config")

	// 获取环境参数
	env, _ := rootCmd.PersistentFlags().GetString("env")
	debug, _ := rootCmd.PersistentFlags().GetBool("debug")

	// 设置环境变量（如果通过命令行指定）
	if env != "" {
		os.Setenv("GO_STARTER_APP_ENVIRONMENT", env)
	}
	if debug {
		os.Setenv("GO_STARTER_APP_DEBUG", "true")
	}

	var err error
	if configFile != "" {
		// 从指定文件加载配置
		appConfig, err = config.LoadFromFile(configFile)
		if err != nil {
			log.Crit("Failed to load config from file", "file", configFile, "error", err)
		}
	} else {
		// 从环境变量和默认值加载配置
		appConfig, err = config.LoadFromEnv()
		if err != nil {
			log.Crit("Failed to load config from environment", "error", err)
		}
	}

	// 验证配置
	if err := appConfig.Validate(); err != nil {
		log.Crit("Config validation failed", "error", err)
	}

	// 打印配置信息
	log.Info("Configuration loaded successfully",
		"app_name", appConfig.App.Name,
		"version", appConfig.App.Version,
		"environment", appConfig.App.Environment,
		"debug", appConfig.App.Debug,
		"api_address", appConfig.API.Address,
		"log_level", appConfig.Log.Level,
	)

	// 如果是开发环境，打印更多配置信息
	if appConfig.IsDevelopment() {
		log.Debug("Development mode - showing detailed config",
			"mysql_host", getMySQLHost(),
			"redis_host", getRedisHost(),
			"kafka_brokers", getKafkaBrokers(),
		)
	}
}

// getMySQLHost 获取MySQL主机地址（安全地）
func getMySQLHost() string {
	if appConfig.MySQL != "" {
		return appConfig.MySQL
	}
	return "not configured"
}

// getRedisHost 获取Redis主机地址（安全地）
func getRedisHost() string {
	if appConfig.Redis != "" {
		return appConfig.Redis
	}
	return "not configured"
}

// getKafkaBrokers 获取Kafka代理地址（安全地）
func getKafkaBrokers() []string {
	if appConfig.Kafka != nil {
		return appConfig.Kafka.Brokers
	}
	return []string{"not configured"}
}

// GetAppConfig 获取应用配置
func GetAppConfig() *types.AppConfig {
	return appConfig
}
