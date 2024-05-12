package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/config"
	"github.com/kanaryorgs/kanary-server/pkg/controllers"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
)

func NewRouter(kh *k8s.K8sHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(config.LoggerFormatter))
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		setUpPodRoutes(v1, kh)
		setUpNodeRoutes(v1, kh)
	}

	return router
}

// setUpPodRoutes sets up routing for pod related endpoints.
// /v1/pods/
func setUpPodRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	podController := controllers.NewPodController(kh)

	api.GET("/pods", podController.GetPodList)
	api.GET("/pods/:podName", podController.GetPod)

}

// setUpNodeRoutes sets up routing for node related endpoints.
// /v1/pods/
func setUpNodeRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	nodeController := controllers.NewNodeController(kh)

	api.GET("/nodes", nodeController.GetNodes)
}
