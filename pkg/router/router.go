package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/config"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
)

func NewRouter(kh *k8s.K8sHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(config.LoggerFormatter))
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		setUpPodRoutes(v1)
		// setUpNodeRoutes(v1)
	}

	return router
}

func setUpPodRoutes(api *gin.RouterGroup) {

}

// setUpNodeRoutes
