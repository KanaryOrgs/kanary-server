package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type JobController struct {
	kh *k8s.K8sHandler
}

func NewJobController(kh *k8s.K8sHandler) *JobController {
	return &JobController{kh: kh}
}

func (jc *JobController) GetJobs(c *gin.Context) {
	jobs, err := jc.kh.ListJobs(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get jobs: %v", err)})
		return
	}
	serializedJobs := serializer.SerializeJobList(jobs)
	c.JSON(http.StatusOK, serializedJobs)
}

func (jc *JobController) GetJob(c *gin.Context) {
	jobName := c.Param("name")
	namespace := c.Param("namespace")

	job, err := jc.kh.GetJob(jobName, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get job: %v", err)})
		return
	}

	serializedJob := serializer.SerializeJobDetails(job)
	c.JSON(http.StatusOK, serializedJob)
}
