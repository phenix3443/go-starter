/*
Copyright © 2025 phenix3443 <phenix3443@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"

	"github.com/phenix3443/go-starter/pkg/api/router/v1"
	"github.com/phenix3443/go-starter/pkg/cache"
	"github.com/phenix3443/go-starter/pkg/database"
	"github.com/phenix3443/go-starter/pkg/types"
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
	config   *types.AppConfig
}

func newAPIServer() (*apiServer, error) {
	config := GetAppConfig()

	// 创建数据库连接
	var db *database.AppDB
	var err error

	if config.MySQL != "" {
		db, err = database.NewAppDB(config.MySQL)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL database: %w", err)
		}
	} else if config.Postgres != "" {
		db, err = database.NewAppDB(config.Postgres)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no database configuration found")
	}

	// 创建缓存连接
	cache, err := cache.NewAppCache(config.Redis)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &apiServer{
		appDB:    db,
		appCache: cache,
		config:   config,
	}, nil
}

func (s *apiServer) start(ctx context.Context) error {
	r := router.SetupRouter(s.appDB, s.appCache)

	// 设置服务器超时
	srv := &http.Server{
		Addr:         s.config.API.Address,
		Handler:      r,
		ReadTimeout:  time.Duration(s.config.API.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.API.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.config.API.IdleTimeout) * time.Second,
	}

	// 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Server failed to start", "error", err)
		}
	}()

	// 等待上下文取消
	<-ctx.Done()

	// 优雅关闭
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("Server shutdown failed", "error", err)
		return err
	}

	log.Info("Server shutdown gracefully")
	return nil
}
