package serializer

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	IP        string            `json:"ip"`
	Images    []string          `json:"images"`
	Status    string            `json:"status"`
	CPUUsage  int64             `json:"cpu_usage"`
	MemUsage  int64             `json:"mem_usage"`
	Labels    map[string]string `json:"labels"`
	Restarts  int32             `json:"restarts"`
	NodeName  string            `json:"node_name"`
	StartTime *metav1.Time      `json:"start_time"`
	Volumes   []string          `json:"volumes"`
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
func SerializePodDetails(pod *v1.Pod, cpuUsage, memUsage int64) PodDetails {
	images := make([]string, len(pod.Spec.Containers))
	volumes := make([]string, len(pod.Spec.Volumes))

	for j, container := range pod.Spec.Containers {
		images[j] = container.Image
	}

	for i, volume := range pod.Spec.Volumes {
		volumes[i] = volume.Name
	}

	serializePod := PodDetails{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		IP:        pod.Status.PodIP,
		Images:    images,
		Status:    string(pod.Status.Phase),
		CPUUsage:  cpuUsage,
		MemUsage:  memUsage,
		Labels:    pod.Labels,
		Restarts:  getRestartCount(*pod),
		NodeName:  pod.Spec.NodeName,
		StartTime: pod.Status.StartTime,
		Volumes:   volumes,
	}

	return serializePod
}

func getRestartCount(pod v1.Pod) int32 {
	var restartCount int32 = 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restartCount += containerStatus.RestartCount
	}
	return restartCount
}
