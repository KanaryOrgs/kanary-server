package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type DaemonSetController struct {
	kh *k8s.K8sHandler
}

func NewDaemonSetController(kh *k8s.K8sHandler) *DaemonSetController {
	return &DaemonSetController{kh: kh}
}

// GetDaemonSets handles the GET requests to list DaemonSets.
func (dsc *DaemonSetController) GetDaemonSets(c *gin.Context) {
	daemonSets, err := dsc.kh.GetDaemonSets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get daemonsets: %v", err)})
		return
	}
	serializedDaemonSets := serializer.SerializeDaemonSetList(daemonSets)
	c.JSON(http.StatusOK, serializedDaemonSets)
}

// GetDaemonSet handles the GET requests to retrieve a specific DaemonSet.
func (dsc *DaemonSetController) GetDaemonSet(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")

	daemonSet, err := dsc.kh.GetDaemonSet(namespace, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get daemonset: %v", err)})
		return
	}

	serializedDaemonSet := serializer.SerializeDaemonSetDetails(daemonSet)
	c.JSON(http.StatusOK, serializedDaemonSet)
}
