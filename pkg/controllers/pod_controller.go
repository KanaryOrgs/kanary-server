package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
	"net/http"
)

type PodController struct {
	kh *k8s.K8sHandler
}

func NewPodController(kh *k8s.K8sHandler) *PodController {
	return &PodController{kh: kh}
}

// GetPodList handles the GET requests to list Pods.
func (pc *PodController) GetPodList(c *gin.Context) {
	pods, err := pc.kh.ListPods(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get pods: %v", err)})
		return
	}
	serializedPods := serializer.SerializePodList(pods)
	c.JSON(http.StatusOK, serializedPods)
}

func (pc *PodController) GetPod(c *gin.Context) {
	podName := c.Param("podName")
	namespace := c.Query("namespace")

	pod, err := pc.kh.GetPod(podName, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get pod: %v", err)})
		return
	}

	cpuUsage, memUsage, err := pc.kh.GetPodUsage(podName, namespace)
	if err != nil {
		c.Error(err).SetMeta("Failed to get pod usage")
	}

	serializePod := serializer.SerializePodDetails(pod, cpuUsage, memUsage)
	c.JSON(http.StatusOK, serializePod)
}
