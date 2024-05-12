package k8s

import (
	"context"
	"errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListPods handles the listing of all pods.
func (kh *K8sHandler) ListPods(namespace string) (*v1.PodList, error) {
	var pods *v1.PodList
	var err error

	if namespace == "" {
		pods, err = kh.K8sClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	} else {
		pods, err = kh.K8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return pods, nil
}

// GetPod retrieves a single pod by name within a given namespace.
func (kh *K8sHandler) GetPod(podName, namespace string) (*v1.Pod, error) {
	if podName == "" {
		return nil, errors.New("pod name must be provided")
	}

	pod, err := kh.K8sClient.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return pod, nil
}

func (kh *K8sHandler) GetPodUsage(podName, namespace string) (int64, int64, error) {
	// Get the current CPU and memory usage of the pod
	podMetrics, err := kh.MetricK8sClient.MetricsV1beta1().PodMetricses(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		return 0.0, 0.0, err
	}

	var totalCpuUsage int64
	var totalMemUsage int64
	for _, container := range podMetrics.Containers {
		totalCpuUsage += container.Usage.Cpu().MilliValue()
		totalMemUsage += container.Usage.Memory().Value() / 1048576
	}

	return totalCpuUsage, totalMemUsage, nil
}
