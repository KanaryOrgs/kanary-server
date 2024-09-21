package k8s

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kh *K8sHandler) GetDaemonSets() (*appsv1.DaemonSetList, error) {
	return kh.K8sClient.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})
}

func (kh *K8sHandler) GetDaemonSet(namespace, name string) (*appsv1.DaemonSet, error) {
	return kh.K8sClient.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}
