package serializer

import (
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolumeList struct {
	Name          string            `json:"name"`
	Capacity      map[string]string `json:"capacity"`
	AccessModes   []string          `json:"access_modes"`
	ReclaimPolicy string            `json:"reclaim_policy"`
	Status        string            `json:"status"`
	Labels        map[string]string `json:"labels"`
}

type PersistentVolumeClaimList struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	VolumeName  string            `json:"volume_name"`
	AccessModes []string          `json:"access_modes"`
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
}

type StorageClassList struct {
	Name                 string            `json:"name"`
	Provisioner          string            `json:"provisioner"`
	ReclaimPolicy        string            `json:"reclaim_policy"`
	AllowVolumeExpansion bool              `json:"allow_volume_expansion"`
	Labels               map[string]string `json:"labels"`
}

type PersistentVolumeDetail struct {
	Name          string            `json:"name"`
	Capacity      map[string]string `json:"capacity"`
	AccessModes   []string          `json:"access_modes"`
	ReclaimPolicy string            `json:"reclaim_policy"`
	Status        string            `json:"status"`
	Labels        map[string]string `json:"labels"`
	CreationTime  *metav1.Time      `json:"creation_time"`
	StorageClass  string            `json:"storage_class"`
}

type PersistentVolumeClaimDetail struct {
	Name           string            `json:"name"`
	Namespace      string            `json:"namespace"`
	VolumeName     string            `json:"volume_name"`
	AccessModes    []string          `json:"access_modes"`
	Status         string            `json:"status"`
	Labels         map[string]string `json:"labels"`
	CreationTime   *metav1.Time      `json:"creation_time"`
	StorageRequest map[string]string `json:"storage_request"`
}

type StorageClassDetail struct {
	Name                 string            `json:"name"`
	Provisioner          string            `json:"provisioner"`
	ReclaimPolicy        string            `json:"reclaim_policy"`
	AllowVolumeExpansion bool              `json:"allow_volume_expansion"`
	Labels               map[string]string `json:"labels"`
	CreationTime         *metav1.Time      `json:"creation_time"`
	Parameters           map[string]string `json:"parameters"`
}

func SerializePersistentVolumeList(pvList *v1.PersistentVolumeList) []PersistentVolumeList {
	if pvList == nil {
		return nil
	}

	serializedPVList := make([]PersistentVolumeList, len(pvList.Items))
	for i, pv := range pvList.Items {
		accessModes := make([]string, len(pv.Spec.AccessModes))
		for j, mode := range pv.Spec.AccessModes {
			accessModes[j] = string(mode)
		}

		capacity := make(map[string]string)
		for key, val := range pv.Spec.Capacity {
			capacity[string(key)] = val.String()
		}
		serializedPVList[i] = PersistentVolumeList{
			Name:          pv.Name,
			Capacity:      capacity,
			AccessModes:   accessModes,
			ReclaimPolicy: string(pv.Spec.PersistentVolumeReclaimPolicy),
			Status:        string(pv.Status.Phase),
			Labels:        pv.Labels,
		}
	}
	return serializedPVList
}

func SerializePersistentVolumeClaimList(pvcList *v1.PersistentVolumeClaimList) []PersistentVolumeClaimList {
	if pvcList == nil {
		return nil
	}

	serializedPVCList := make([]PersistentVolumeClaimList, len(pvcList.Items))
	for i, pvc := range pvcList.Items {
		accessModes := make([]string, len(pvc.Spec.AccessModes))
		for j, mode := range pvc.Spec.AccessModes {
			accessModes[j] = string(mode)
		}
		serializedPVCList[i] = PersistentVolumeClaimList{
			Name:        pvc.Name,
			Namespace:   pvc.Namespace,
			VolumeName:  pvc.Spec.VolumeName,
			AccessModes: accessModes,
			Status:      string(pvc.Status.Phase),
			Labels:      pvc.Labels,
		}
	}
	return serializedPVCList
}

func SerializeStorageClassList(scList *storagev1.StorageClassList) []StorageClassList {
	if scList == nil {
		return nil
	}

	serializedSCList := make([]StorageClassList, len(scList.Items))
	for i, sc := range scList.Items {
		allowVolumeExpansion := false
		if sc.AllowVolumeExpansion != nil {
			allowVolumeExpansion = *sc.AllowVolumeExpansion
		}
		reclaimPolicy := ""
		if sc.ReclaimPolicy != nil {
			reclaimPolicy = string(*sc.ReclaimPolicy)
		}

		serializedSCList[i] = StorageClassList{
			Name:                 sc.Name,
			Provisioner:          sc.Provisioner,
			ReclaimPolicy:        reclaimPolicy,
			AllowVolumeExpansion: allowVolumeExpansion,
			Labels:               sc.Labels,
		}
	}
	return serializedSCList
}

func SerializePersistentVolumeDetail(pv *v1.PersistentVolume) PersistentVolumeDetail {
	accessModes := make([]string, len(pv.Spec.AccessModes))
	for i, mode := range pv.Spec.AccessModes {
		accessModes[i] = string(mode)
	}

	capacity := make(map[string]string)
	for key, val := range pv.Spec.Capacity {
		capacity[string(key)] = val.String()
	}

	return PersistentVolumeDetail{
		Name:          pv.Name,
		Capacity:      capacity,
		AccessModes:   accessModes,
		ReclaimPolicy: string(pv.Spec.PersistentVolumeReclaimPolicy),
		Status:        string(pv.Status.Phase),
		Labels:        pv.Labels,
		CreationTime:  &pv.CreationTimestamp,
		StorageClass:  pv.Spec.StorageClassName,
	}
}

func SerializePersistentVolumeClaimDetail(pvc *v1.PersistentVolumeClaim) PersistentVolumeClaimDetail {
	accessModes := make([]string, len(pvc.Spec.AccessModes))
	for i, mode := range pvc.Spec.AccessModes {
		accessModes[i] = string(mode)
	}

	storageRequest := make(map[string]string)
	storage := pvc.Spec.Resources.Requests.Storage()
	if storage != nil {
		storageRequest["storage"] = storage.String()
	}

	return PersistentVolumeClaimDetail{
		Name:           pvc.Name,
		Namespace:      pvc.Namespace,
		VolumeName:     pvc.Spec.VolumeName,
		AccessModes:    accessModes,
		Status:         string(pvc.Status.Phase),
		Labels:         pvc.Labels,
		CreationTime:   &pvc.CreationTimestamp,
		StorageRequest: storageRequest,
	}
}

func SerializeStorageClassDetail(sc *storagev1.StorageClass) StorageClassDetail {
	parameters := make(map[string]string)
	for key, value := range sc.Parameters {
		parameters[key] = value
	}

	allowVolumeExpansion := false
	if sc.AllowVolumeExpansion != nil {
		allowVolumeExpansion = *sc.AllowVolumeExpansion
	}
	reclaimPolicy := ""
	if sc.ReclaimPolicy != nil {
		reclaimPolicy = string(*sc.ReclaimPolicy)
	}

	return StorageClassDetail{
		Name:                 sc.Name,
		Provisioner:          sc.Provisioner,
		ReclaimPolicy:        reclaimPolicy,
		AllowVolumeExpansion: allowVolumeExpansion,
		Labels:               sc.Labels,
		CreationTime:         &sc.CreationTimestamp,
		Parameters:           parameters,
	}
}
