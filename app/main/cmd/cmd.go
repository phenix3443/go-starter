package cmd

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/phenix3443/go-starter/pkg/types"
)

var appConfig types.AppConfig

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "config file (default is ./config.yaml)")
}

func initConfig() {
	configFile, _ := rootCmd.PersistentFlags().GetString("config")

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Crit("Failed to read config file", "error", err)
	}
	log.Info("Using config file:", "path:", viper.ConfigFileUsed())
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Crit("Failed to unmarshal config", "error", err)
	}
}
