package serializer

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSetList struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Replicas  int32             `json:"replicas"`
	Ready     int32             `json:"ready"`
	Labels    map[string]string `json:"labels"`
}

type StatefulSetDetails struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Replicas     int32             `json:"replicas"`
	Ready        int32             `json:"ready"`
	Labels       map[string]string `json:"labels"`
	CreationTime *metav1.Time      `json:"creation_time"`
}

func SerializeStatefulSetList(statefulSetList *appsv1.StatefulSetList) []StatefulSetList {
	if statefulSetList == nil {
		return nil
	}

	serializedStatefulSetList := make([]StatefulSetList, len(statefulSetList.Items))
	for i, statefulSet := range statefulSetList.Items {
		serializedStatefulSetList[i] = StatefulSetList{
			Name:      statefulSet.Name,
			Namespace: statefulSet.Namespace,
			Replicas:  *statefulSet.Spec.Replicas,
			Ready:     statefulSet.Status.ReadyReplicas,
			Labels:    statefulSet.Labels,
		}
	}
	return serializedStatefulSetList
}

func SerializeStatefulSetDetails(statefulSet *appsv1.StatefulSet) StatefulSetDetails {
	return StatefulSetDetails{
		Name:         statefulSet.Name,
		Namespace:    statefulSet.Namespace,
		Replicas:     *statefulSet.Spec.Replicas,
		Ready:        statefulSet.Status.ReadyReplicas,
		Labels:       statefulSet.Labels,
		CreationTime: &statefulSet.CreationTimestamp,
	}
}
