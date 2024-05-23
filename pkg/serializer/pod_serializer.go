package serializer

import (
	v1 "k8s.io/api/core/v1"
)

type PodList struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Images    []string          `json:"images"`
	IP        string            `json:"ip"`
	Status    string            `json:"status"`
	Labels    map[string]string `json:"labels"`
	Restarts  int32             `json:"restarts"`
}

// PodDetails represents detailed information about a single pod.
type PodDetails struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Containers   []Container       `json:"containers"`
	RestartCount int32             `json:"restartCount"`
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
}

// Container represents information about a single container within a pod.
type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// SerializePods serializes a PodList to a slice of PodList structures.
func SerializePodList(podList *v1.PodList) []PodList {
	if podList == nil {
		return nil
	}

	serializedPodList := make([]PodList, len(podList.Items))
	for i, pod := range podList.Items {
		images := make([]string, len(pod.Spec.Containers))
		for j, container := range pod.Spec.Containers {
			images[j] = container.Image
		}
		serializedPodList[i] = PodList{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			IP:        pod.Status.PodIP,
			Images:    images,
			Status:    string(pod.Status.Phase),
			Labels:    pod.Labels,
			Restarts:  getRestartCount(pod),
		}
	}
	return serializedPodList
}

// SerializePodDetails serializes a PodList to a slice of PodDetails structures.
func SerializePodDetails(podList *v1.PodList) []PodDetails {
	if podList == nil {
		return nil
	}

	serializedPods := make([]PodDetails, len(podList.Items))
	for i, pod := range podList.Items {
		containers := make([]Container, len(pod.Spec.Containers))
		for j, container := range pod.Spec.Containers {
			containers[j] = Container{
				Name:  container.Name,
				Image: container.Image,
			}
		}
		serializedPods[i] = PodDetails{
			Name:         pod.Name,
			Namespace:    pod.Namespace,
			Containers:   containers,
			RestartCount: pod.Status.ContainerStatuses[0].RestartCount,
			Status:       string(pod.Status.Phase),
			Labels:       pod.Labels,
		}
	}
	return serializedPods
}

func getRestartCount(pod v1.Pod) int32 {
	var restartCount int32 = 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restartCount += containerStatus.RestartCount
	}
	return restartCount
}
