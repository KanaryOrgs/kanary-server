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

func (nc *NodeController) GetNodeCount(c *gin.Context) {
	node, err := nc.kh.ListNodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get nodes: %v", err)})
		return
	}
	serializedPods := serializer.SerializeNodeCount(node)
	c.JSON(http.StatusOK, serializedPods)
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

func (nc *NodeController) GetNode(c *gin.Context) {
	nodeName := c.Param("name")

	node, err := nc.kh.GetNode(nodeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get node: %v", err)})
		return
	}
	serializedPods := serializer.SerializeNodeDetails(node)
	c.JSON(http.StatusOK, serializedPods)
}
