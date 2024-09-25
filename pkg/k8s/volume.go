package k8s

import (
	"context"
	"errors"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListPersistentVolumes lists all PersistentVolumes in a namespace.
func (kh *K8sHandler) ListPersistentVolumes() (*v1.PersistentVolumeList, error) {
	persistentVolumes, err := kh.K8sClient.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return persistentVolumes, nil
}

// ListPersistentVolumeClaims lists all PersistentVolumeClaims in a namespace.
func (kh *K8sHandler) ListPersistentVolumeClaims(namespace string) (*v1.PersistentVolumeClaimList, error) {
	var persistentVolumeClaims *v1.PersistentVolumeClaimList
	var err error

	if namespace == "" {
		persistentVolumeClaims, err = kh.K8sClient.CoreV1().PersistentVolumeClaims("").List(context.TODO(), metav1.ListOptions{})
	} else {
		persistentVolumeClaims, err = kh.K8sClient.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return persistentVolumeClaims, nil
}

// GetPersistentVolume retrieves a single PersistentVolume by name.
func (kh *K8sHandler) GetPersistentVolume(pvName string) (*v1.PersistentVolume, error) {
	if pvName == "" {
		return nil, errors.New("persistent volume name must be provided")
	}

	persistentVolume, err := kh.K8sClient.CoreV1().PersistentVolumes().Get(context.TODO(), pvName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return persistentVolume, nil
}

// GetPersistentVolumeClaim retrieves a single PersistentVolumeClaim by name in a given namespace.
func (kh *K8sHandler) GetPersistentVolumeClaim(pvcName, namespace string) (*v1.PersistentVolumeClaim, error) {
	if pvcName == "" {
		return nil, errors.New("persistent volume claim name must be provided")
	}

	persistentVolumeClaim, err := kh.K8sClient.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return persistentVolumeClaim, nil
}
