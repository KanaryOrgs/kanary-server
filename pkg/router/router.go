package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/kanaryorgs/kanary-server/docs"
	"github.com/kanaryorgs/kanary-server/pkg/config"
	"github.com/kanaryorgs/kanary-server/pkg/controllers"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(kh *k8s.K8sHandler) *gin.Engine {
	router := gin.New()
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Title = "kanary-server API"
	router.Use(gin.LoggerWithFormatter(config.LoggerFormatter))
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := router.Group("/v1")
	{
		setUpPodRoutes(v1, kh)
		setUpNodeRoutes(v1, kh)
		setUpServiceRoutes(v1, kh)
		setUpIngressRoutes(v1, kh)
		setUpDaemonSetRoutes(v1, kh)

	}

	return router
}

// setUpPodRoutes sets up routing for pod related endpoints.
// /v1/pods/
func setUpPodRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	podController := controllers.NewPodController(kh)

	api.GET("/pods", podController.GetPods)
	api.GET("/pods/:namespace/:name", podController.GetPod)
	api.GET("/pods/logs/:namespace/:name", podController.GetLogsOfPod)

}

// setUpNodeRoutes sets up routing for node related endpoints.
// /v1/nodes/
func setUpNodeRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	nodeController := controllers.NewNodeController(kh)

	api.GET("/nodes", nodeController.GetNodes)
	api.GET("/nodes/:name", nodeController.GetNode)
	api.GET("/nodes/count", nodeController.GetNodeCount)

}

// setUpServiceRoutes sets up routing for service related endpoints.
// /v1/services/
func setUpServiceRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	serviceController := controllers.NewServiceController(kh)

	api.GET("/services", serviceController.GetServices)
}

// setUpIngressRoutes sets up routing for ingress related endpoints.
// /v1/ingress/
func setUpIngressRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	ingressController := controllers.NewIngressController(kh)

	api.GET("/ingresses", ingressController.GetIngresses)
}

// setUpDaemonSetRoutes sets up routing for daemonset related endpoints.
// /v1/daemonsets/
func setUpDaemonSetRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	daemonSetController := controllers.NewDaemonSetController(kh)

	api.GET("/daemonsets", daemonSetController.GetDaemonSets)
	api.GET("/daemonsets/:namespace/:name", daemonSetController.GetDaemonSet)
}
