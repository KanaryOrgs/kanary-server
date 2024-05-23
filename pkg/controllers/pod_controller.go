package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type PodController struct {
	kh *k8s.K8sHandler
}

func NewPodController(kh *k8s.K8sHandler) *PodController {
	return &PodController{kh: kh}
}

// GetPods godoc
// @Summary Show pod list.
// @Schemes
// @Description get pod list in k8s cluster.
// @Tags pods
// @Accept */*
// @Produce json
// @Success 200 {array} serializer.PodList
// @Router /pods [get]
func (pc *PodController) GetPods(c *gin.Context) {
	pods, err := pc.kh.ListPods(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get pods: %v", err)})
		return
	}
	serializedPods := serializer.SerializePodList(pods)
	c.JSON(http.StatusOK, serializedPods)
}
