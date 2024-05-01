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

// GetPods handles the GET requests to list Pods.
func (pc *PodController) GetPods(c *gin.Context) {
	pods, err := pc.kh.ListPods(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get pods: %v", err)})
		return
	}
	serializedPods := serializer.SerializePodList(pods)
	c.JSON(http.StatusOK, serializedPods)
}
