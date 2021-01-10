package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// group all routes under /v1/api
	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		// prefix the user routes
		user := v1.Group("/user")
		{
			user.POST("", s.CreateUser())
		}

		// prefix the weight routes
		weight := v1.Group("/weight")
		{
			weight.POST("", s.CreateWeightEntry())
		}
	}

	return router
}
