package router

import (
	"github.com/gin-gonic/gin"

	"github.com/phenix3443/go-starter/pkg/cache"
	"github.com/phenix3443/go-starter/pkg/database"
)

// SetupRouter 创建并配置 Gin 路由引擎
func SetupRouter(appDB *database.AppDB, appCache *cache.AppCache) *gin.Engine {
	r := gin.Default()
	r.Use(injectAppDB(appDB), injectAppCache(appCache))
	apiV1 := r.Group("/api/v1")
	registerV1Routes(apiV1)
	return r
}

// injectAppDB 返回一个注入数据库实例的 Gin 中间件
func injectAppDB(appDB *database.AppDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_db", appDB)
		c.Next()
	}
}

func injectAppCache(appCache *cache.AppCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_cache", appCache)
		c.Next()
	}
}
