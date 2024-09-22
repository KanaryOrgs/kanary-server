package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type DeploymentController struct {
	kh *k8s.K8sHandler
}

func NewDeploymentController(kh *k8s.K8sHandler) *DeploymentController {
	return &DeploymentController{kh: kh}
}

// GetDeployments handles the GET requests to list Deployments.
func (dc *DeploymentController) GetDeployments(c *gin.Context) {
	deployments, err := dc.kh.ListDeployments(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get deployments: %v", err)})
		return
	}
	serializedDeployments := serializer.SerializeDeploymentList(deployments)
	c.JSON(http.StatusOK, serializedDeployments)
}

// GetDeployment handles the GET request to retrieve a specific Deployment.
func (dc *DeploymentController) GetDeployment(c *gin.Context) {
	deploymentName := c.Param("name")
	namespace := c.Param("namespace")

	deployment, err := dc.kh.GetDeployment(deploymentName, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get deployment: %v", err)})
		return
	}

	serializedDeployment := serializer.SerializeDeploymentDetails(deployment)
	c.JSON(http.StatusOK, serializedDeployment)
}
