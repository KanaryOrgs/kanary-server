package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type CronJobController struct {
	kh *k8s.K8sHandler
}

func NewCronJobController(kh *k8s.K8sHandler) *CronJobController {
	return &CronJobController{kh: kh}
}

// GetCronJobs handles the GET requests to list CronJobs.
func (cjc *CronJobController) GetCronJobs(c *gin.Context) {
	cronJobs, err := cjc.kh.GetCronJobs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get cronjobs: %v", err)})
		return
	}
	serializedCronJobs := serializer.SerializeCronJobList(cronJobs)
	c.JSON(http.StatusOK, serializedCronJobs)
}

// GetCronJob handles the GET requests to retrieve a specific CronJob.
func (cjc *CronJobController) GetCronJob(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")

	cronJob, err := cjc.kh.GetCronJob(namespace, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get cronjob: %v", err)})
		return
	}

	serializedCronJob := serializer.SerializeCronJobDetails(cronJob)
	c.JSON(http.StatusOK, serializedCronJob)
}
