/*
Copyright © 2025 phenix3443
*/
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"

	"github.com/phenix3443/go-starter/app/your-app/cmd"
)

func main() {
	log.SetDefault(log.NewLogger(log.JSONHandlerWithLevel(os.Stdout, log.LevelDebug)))
	cmd.Execute()
}
