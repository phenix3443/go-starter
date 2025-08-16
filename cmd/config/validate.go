package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"

	"github.com/phenix3443/go-starter/pkg/config"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configuration file",
	Long:  `Validate the configuration file for syntax and required fields`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile, _ := cmd.Flags().GetString("config")

		if configFile == "" {
			return fmt.Errorf("config file path is required")
		}

		// 检查文件是否存在
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			return fmt.Errorf("config file does not exist: %s", configFile)
		}

		// 加载并验证配置
		cfg, err := config.LoadFromFile(configFile)
		if err != nil {
			return fmt.Errorf("config validation failed: %w", err)
		}

		// 打印配置摘要
		fmt.Printf("✅ Configuration validation successful!\n\n")
		fmt.Printf("📋 Configuration Summary:\n")
		fmt.Printf("   App Name: %s\n", cfg.App.Name)
		fmt.Printf("   Version: %s\n", cfg.App.Version)
		fmt.Printf("   Environment: %s\n", cfg.App.Environment)
		fmt.Printf("   Debug Mode: %t\n", cfg.App.Debug)
		fmt.Printf("   API Address: %s\n", cfg.API.Address)
		fmt.Printf("   Log Level: %s\n", cfg.Log.Level)

		if cfg.MySQL != "" {
			fmt.Printf("   MySQL: %s\n", cfg.MySQL)
		}
		if cfg.Postgres != "" {
			fmt.Printf("   PostgreSQL: %s\n", cfg.Postgres)
		}
		if cfg.Redis != "" {
			fmt.Printf("   Redis: %s\n", cfg.Redis)
		}

		return nil
	},
}

func init() {
	validateCmd.Flags().StringP("config", "c", "", "config file path")
	validateCmd.MarkFlagRequired("config")
}

func main() {
	if err := validateCmd.Execute(); err != nil {
		log.Error("Validation failed", "error", err)
		os.Exit(1)
	}
}
