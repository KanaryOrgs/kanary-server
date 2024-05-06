package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
	"net/http"
)

type NodeController struct {
	kh *k8s.K8sHandler
}

func NewNodeController(kh *k8s.K8sHandler) *NodeController {
	return &NodeController{kh: kh}
}

// GetNodes handles the GET requests to list Nodes.
func (nc *NodeController) GetNodes(c *gin.Context) {
	nodes, err := nc.kh.ListNodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get nodes: %v", err)})
		return
	}
	serializedPods := serializer.SerializeNodeList(nodes)
	c.JSON(http.StatusOK, serializedPods)
}
