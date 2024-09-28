package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type StatefulSetController struct {
	kh *k8s.K8sHandler
}

func NewStatefulSetController(kh *k8s.K8sHandler) *StatefulSetController {
	return &StatefulSetController{kh: kh}
}

func (sc *StatefulSetController) GetStatefulSets(c *gin.Context) {
	statefulSets, err := sc.kh.ListStatefulSets(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get statefulsets: %v", err)})
		return
	}
	serializedStatefulSets := serializer.SerializeStatefulSetList(statefulSets)
	c.JSON(http.StatusOK, serializedStatefulSets)
}

func (sc *StatefulSetController) GetStatefulSet(c *gin.Context) {
	statefulSetName := c.Param("name")
	namespace := c.Param("namespace")

	statefulSet, err := sc.kh.GetStatefulSet(statefulSetName, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get statefulset: %v", err)})
		return
	}

	serializedStatefulSet := serializer.SerializeStatefulSetDetails(statefulSet)
	c.JSON(http.StatusOK, serializedStatefulSet)
}
