package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type NodeController struct {
	kh *k8s.K8sHandler
}

func NewNodeController(kh *k8s.K8sHandler) *NodeController {
	return &NodeController{kh: kh}
}

// GetNodes godoc
// @Summary Show node list.
// @Schemes
// @Description get node list in k8s cluster.
// @Tags nodes
// @Accept */*
// @Produce json
// @Success 200 {array} serializer.NodeList
// @Router /nodes [get]
func (nc *NodeController) GetNodes(c *gin.Context) {
	nodes, err := nc.kh.ListNodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get nodes: %v", err)})
		return
	}
	serializedPods := serializer.SerializeNodeList(nodes)
	c.JSON(http.StatusOK, serializedPods)
}
