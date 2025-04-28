package handler

import (
	"net/http"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/phenix3443/go-starter/pkg/database/models"
)

type savePersonRequest struct {
	Age uint64 `json:"age" binding:"required"`
}

type savePersonResponse struct {
	Message string `json:"message"`
}

func SavePerson(c *gin.Context) {
	name := c.Param("name")
	var req savePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p := &models.Person{
		Name: name,
		Age:  req.Age,
	}
	if err := getAppDB(c).SavePerson(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save person to app db"})
		return
	}
	if err := getAppCache(c).SavePerson(c.Request.Context(), p.Name, p.Age); err != nil {
		log.Error("failed to save person to redis", "error", err)
	}
	resp := savePersonResponse{
		Message: "success",
	}
	c.JSON(http.StatusOK, resp)
}

func GetPerson(c *gin.Context) {
	name := c.Param("name")
	age, err := getAppCache(c).GetPerson(c.Request.Context(), name)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"name": name, "age": age})
		return
	}
	p, err := getAppDB(c).GetPerson(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get person from app db"})
		return
	}
	log.Debug("get person from app db", "name", name, "age", p.Age)
	c.JSON(http.StatusOK, p)
}
