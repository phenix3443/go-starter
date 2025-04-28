/*
Copyright © 2025 phenix3443 <phenix3443@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/phenix3443/go-starter/pkg/api/router/v1"
	"github.com/phenix3443/go-starter/pkg/cache"
	"github.com/phenix3443/go-starter/pkg/database"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "api is a command to provide api endpoints",
	RunE: func(cmd *cobra.Command, _ []string) error {
		srv, err := newAPIServer()
		if err != nil {
			return fmt.Errorf("failed to create api server: %w", err)
		}
		return srv.start(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

type apiServer struct {
	appDB    *database.AppDB
	appCache *cache.AppCache
}

func newAPIServer() (*apiServer, error) {
	db, err := database.NewAppDB(appConfig.MySQL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	cache, err := cache.NewAppCache(appConfig.Redis)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}
	return &apiServer{
		appDB:    db,
		appCache: cache,
	}, nil
}

func (s *apiServer) start(_ context.Context) error {
	r := router.SetupRouter(s.appDB, s.appCache)
	return r.Run(appConfig.API.Address)
}
