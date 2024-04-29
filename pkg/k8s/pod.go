package k8s

import (
	"context"
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
