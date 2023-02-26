package app

import (
	"PinGo/pkg/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ApiStatus PinGo godoc
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

// GetAll PinGo godoc
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

// PostRequest PinGo godoc
// @Summary Post request
// @Tags request
// @Accept json
// @Produce json
// @Param log body api.RequestPostSchema true "schema"
// @Success 200
// @Router /v1/request [post]
func (s *Server) PostRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newRequest api.RequestPostSchema
		err := c.ShouldBindJSON(&newRequest)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.requestService.Create(&newRequest)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new request created",
		}

		c.JSON(http.StatusOK, response)
	}
}
