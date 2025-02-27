package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "greet",
	Short: "A simple greeting application",
}

var greetCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print a hello message",
	Run: func(cmd *cobra.Command, args []string) {
		// 读取配置信息
		username := viper.GetString("username")
		greeting := viper.GetString("greet")
		fmt.Printf("%s, %s!\n", greeting, username)
	},
}

var configFile string

func init() {
	// 设置根命令的 flag
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.config.json)")

	// 读取配置文件
	cobra.OnInitialize(initConfig)

	// 将 greetCmd 添加到根命令中
	rootCmd.AddCommand(greetCmd)
}

// initConfig 用于初始化配置文件
func initConfig() {
	// 如果指定了配置文件路径，则优先使用它
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}