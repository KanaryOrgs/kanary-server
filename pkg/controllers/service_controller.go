package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type ServiceController struct {
	kh *k8s.K8sHandler
}

func NewServiceController(kh *k8s.K8sHandler) *ServiceController {
	return &ServiceController{kh: kh}
}

// GetServices godoc
// @Summary Show service list.
// @Schemes
// @Description get service list in k8s cluster.
// @Tags services
// @Accept */*
// @Produce json
// @Success 200 {array} serializer.ServiceList
// @Router /services [get]
func (pc *ServiceController) GetServices(c *gin.Context) {
	services, err := pc.kh.ListServices(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get services: %v", err)})
		return
	}
	serializedServices := serializer.SerializeServiceList(services)
	c.JSON(http.StatusOK, serializedServices)
}
