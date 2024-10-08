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
		setUpDeploymentRoutes(v1, kh)
		setUpPersistentVolumeRoutes(v1, kh)
		setUpPersistentVolumeClaimRoutes(v1, kh)
		setUpStorageClassRoutes(v1, kh)
		setUpStatefulSetRoutes(v1, kh)
		setUpJobRoutes(v1, kh)
		setUpCronJobRoutes(v1, kh)
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
// /v1/ingresses/
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

// setUpDeploymentRoutes sets up routing for deployment related endpoints.
// /v1/deployments/
func setUpDeploymentRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	deploymentController := controllers.NewDeploymentController(kh)

	api.GET("/deployments", deploymentController.GetDeployments)
	api.GET("/deployments/:namespace/:name", deploymentController.GetDeployment)
}

// setUpPersistentVolumeRoutes sets up routing for persistent volume related endpoints.
// /v1/pvs/
func setUpPersistentVolumeRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	persistentVolumeController := controllers.NewVolumeController(kh)

	api.GET("/pvs", persistentVolumeController.GetPersistentVolumes)
	api.GET("/pvs/:name", persistentVolumeController.GetPersistentVolume)
}

// setUpPersistentVolumeClaimoutes sets up routing for persistent volume claim related endpoints.
// /v1/pvcs/
func setUpPersistentVolumeClaimRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	persistentVolumeClaimController := controllers.NewVolumeController(kh)

	api.GET("/pvcs", persistentVolumeClaimController.GetPersistentVolumeClaims)
	api.GET("/pvcs/:namespace/:name", persistentVolumeClaimController.GetPersistentVolumeClaim)
}

// setUpStorageClassRoutes sets up routing for storage class related endpoints.
// /v1/storageclasses/
func setUpStorageClassRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	storageClassController := controllers.NewVolumeController(kh)

	api.GET("/storageclasses", storageClassController.GetStorageClasses)
	api.GET("/storageclasses/:name", storageClassController.GetStorageClass)
}

// setUpStatefulSetRoutes sets up routing for statefulset related endpoints.
// /v1/statefulsets/
func setUpStatefulSetRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	statefulSetController := controllers.NewStatefulSetController(kh)

	api.GET("/statefulsets", statefulSetController.GetStatefulSets)
	api.GET("/statefulsets/:namespace/:name", statefulSetController.GetStatefulSet)
}

// setUpJobRoutes sets up routing for job related endpoints.
// /v1/jobs/
func setUpJobRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	jobController := controllers.NewJobController(kh)

	api.GET("/jobs", jobController.GetJobs)
	api.GET("/jobs/:namespace/:name", jobController.GetJob)
}

// setUpCronJobRoutes sets up routing for cronjob related endpoints.
// /v1/cronjobs/
func setUpCronJobRoutes(api *gin.RouterGroup, kh *k8s.K8sHandler) {
	cronJobController := controllers.NewCronJobController(kh)

	api.GET("/cronjobs", cronJobController.GetCronJobs)
	api.GET("/cronjobs/:namespace/:name", cronJobController.GetCronJob)
}
