package serializer

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSetList struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Images       []string          `json:"images"`
	Labels       map[string]string `json:"labels"`
	Current      int32             `json:"current"`
	Ready        int32             `json:"ready"`
	Available    int32             `json:"available"`
	NodeSelector map[string]string `json:"node_selector"`
}

type DaemonSetDetails struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Images       []string          `json:"images"`
	Labels       map[string]string `json:"labels"`
	Desired      int32             `json:"desired"`
	Current      int32             `json:"current"`
	Ready        int32             `json:"ready"`
	Available    int32             `json:"available"`
	NodeSelector map[string]string `json:"node_selector"`
	CreationTime *metav1.Time      `json:"creation_time"`
}

func SerializeDaemonSetList(daemonSetList *appsv1.DaemonSetList) []DaemonSetList {
	if daemonSetList == nil {
		return nil
	}

	serializedDaemonSetList := make([]DaemonSetList, len(daemonSetList.Items))
	for i, ds := range daemonSetList.Items {
		images := make([]string, len(ds.Spec.Template.Spec.Containers))
		for j, container := range ds.Spec.Template.Spec.Containers {
			images[j] = container.Image
		}
		serializedDaemonSetList[i] = DaemonSetList{
			Name:         ds.Name,
			Namespace:    ds.Namespace,
			Images:       images,
			Labels:       ds.Labels,
			Current:      ds.Status.CurrentNumberScheduled,
			Ready:        ds.Status.NumberReady,
			Available:    ds.Status.NumberAvailable,
			NodeSelector: ds.Spec.Template.Spec.NodeSelector,
		}
	}
	return serializedDaemonSetList
}

func SerializeDaemonSetDetails(ds *appsv1.DaemonSet) DaemonSetDetails {
	images := make([]string, len(ds.Spec.Template.Spec.Containers))
	for i, container := range ds.Spec.Template.Spec.Containers {
		images[i] = container.Image
	}

	return DaemonSetDetails{
		Name:         ds.Name,
		Namespace:    ds.Namespace,
		Images:       images,
		Labels:       ds.Labels,
		Current:      ds.Status.CurrentNumberScheduled,
		Ready:        ds.Status.NumberReady,
		Available:    ds.Status.NumberAvailable,
		NodeSelector: ds.Spec.Template.Spec.NodeSelector,
		CreationTime: &ds.CreationTimestamp,
	}
}
