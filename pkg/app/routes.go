package app

import (
	docs "PinGo/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router
	docs.SwaggerInfo.Version = "0.0.1"
	v1 := router.Group("/v1")
	{
		v1.GET("/status", s.ApiStatus())
		log := v1.Group("/log")
		log.GET("", s.GetAll())
		request := v1.Group("/request")
		request.POST("", s.PostRequest())
		receiver := v1.Group("/receiver")
		receiver.POST("", s.PostReceiver())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
