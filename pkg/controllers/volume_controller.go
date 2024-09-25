package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type VolumeController struct {
	kh *k8s.K8sHandler
}

func NewVolumeController(kh *k8s.K8sHandler) *VolumeController {
	return &VolumeController{kh: kh}
}

// GetPVs handles the GET requests to list PersistentVolumes.
func (vc *VolumeController) GetPersistentVolumes(c *gin.Context) {
	persistentVolumes, err := vc.kh.ListPersistentVolumes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get persistent volumes: %v", err)})
		return
	}
	serializedPersistentVolumes := serializer.SerializePersistentVolumeList(persistentVolumes)
	c.JSON(http.StatusOK, serializedPersistentVolumes)
}

// GetPVCs handles the GET requests to list PersistentVolumeClaims.
func (vc *VolumeController) GetPersistentVolumeClaims(c *gin.Context) {
	namespace := c.Query("namespace")
	pvcs, err := vc.kh.ListPersistentVolumeClaims(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get persistent volume claims: %v", err)})
		return
	}
	serializedPVCs := serializer.SerializePersistentVolumeClaimList(pvcs)
	c.JSON(http.StatusOK, serializedPVCs)
}

// GetPV handles the GET request to retrieve a specific PersistentVolume.
func (vc *VolumeController) GetPersistentVolume(c *gin.Context) {
	pvName := c.Param("name")

	pv, err := vc.kh.GetPersistentVolume(pvName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get persistent volume: %v", err)})
		return
	}

	serializedPV := serializer.SerializePersistentVolumeDetail(pv)
	c.JSON(http.StatusOK, serializedPV)
}

// GetPVC handles the GET request to retrieve a specific PersistentVolumeClaim.
func (vc *VolumeController) GetPersistentVolumeClaim(c *gin.Context) {
	pvcName := c.Param("name")
	namespace := c.Param("namespace")

	pvc, err := vc.kh.GetPersistentVolumeClaim(pvcName, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get persistent volume claim: %v", err)})
		return
	}

	serializedPVC := serializer.SerializePersistentVolumeClaimDetail(pvc)
	c.JSON(http.StatusOK, serializedPVC)
}
