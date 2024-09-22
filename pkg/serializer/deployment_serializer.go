package serializer

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentList struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Replicas  int32             `json:"replicas"`
	Available int32             `json:"available"`
	Labels    map[string]string `json:"labels"`
}

type DeploymentDetails struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Replicas     int32             `json:"replicas"`
	Available    int32             `json:"available"`
	Updated      int32             `json:"updated"`
	Ready        int32             `json:"ready"`
	Labels       map[string]string `json:"labels"`
	CreationTime *metav1.Time      `json:"creation_time"`
}

func SerializeDeploymentList(deploymentList *appsv1.DeploymentList) []DeploymentList {
	if deploymentList == nil {
		return nil
	}

	serializedDeploymentList := make([]DeploymentList, len(deploymentList.Items))
	for i, deployment := range deploymentList.Items {
		serializedDeploymentList[i] = DeploymentList{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
			Replicas:  *deployment.Spec.Replicas,
			Available: deployment.Status.AvailableReplicas,
			Labels:    deployment.Labels,
		}
	}
	return serializedDeploymentList
}

func SerializeDeploymentDetails(deployment *appsv1.Deployment) DeploymentDetails {
	return DeploymentDetails{
		Name:         deployment.Name,
		Namespace:    deployment.Namespace,
		Replicas:     *deployment.Spec.Replicas,
		Available:    deployment.Status.AvailableReplicas,
		Updated:      deployment.Status.UpdatedReplicas,
		Ready:        deployment.Status.ReadyReplicas,
		Labels:       deployment.Labels,
		CreationTime: &deployment.CreationTimestamp,
	}
}
