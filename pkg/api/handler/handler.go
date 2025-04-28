package handler

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"

	"github.com/phenix3443/go-starter/pkg/cache"
	"github.com/phenix3443/go-starter/pkg/database"
)

func getAppDB(ctx *gin.Context) *database.AppDB {
	db, ok := ctx.Get("app_db")
	if !ok {
		log.Crit("app_db not found in context")
		return nil
	}
	return db.(*database.AppDB)
}

func getAppCache(ctx *gin.Context) *cache.AppCache {
	c, ok := ctx.Get("app_cache")
	if !ok {
		log.Crit("app_redis not found in context")
		return nil
	}
	return c.(*cache.AppCache)
}
