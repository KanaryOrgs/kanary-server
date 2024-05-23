package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type IngressController struct {
	kh *k8s.K8sHandler
}

func NewIngressController(kh *k8s.K8sHandler) *IngressController {
	return &IngressController{kh: kh}
}

// GetIngresses godoc
// @Summary Show ingress list.
// @Schemes
// @Description get ingress list in k8s cluster.
// @Tags ingresses
// @Accept */*
// @Produce json
// @Success 200 {array} serializer.IngressList
// @Router /ingresses [get]
func (pc *IngressController) GetIngresses(c *gin.Context) {
	ingresses, err := pc.kh.ListIngresses(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get ingresses: %v", err)})
		return
	}
	serializedIngresses := serializer.SerializeIngressList(ingresses)
	c.JSON(http.StatusOK, serializedIngresses)
}
