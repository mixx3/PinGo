package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ApiStatus Log godoc
// @Summary Status
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/status [get]
func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
		}

		c.JSON(http.StatusOK, response)
	}
}

// GetAll HealthCheck godoc
// @Summary Get all
// @Tags log
// @Accept */*
// @Produce json
// @Success 200
// @Router /v1/log [get]
func (s *Server) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		res, err := s.logService.GetAll()
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
