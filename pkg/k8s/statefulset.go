package k8s

import (
	"context"
	"errors"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kh *K8sHandler) ListStatefulSets(namespace string) (*appsv1.StatefulSetList, error) {
	var statefulSets *appsv1.StatefulSetList
	var err error

	if namespace == "" {
		statefulSets, err = kh.K8sClient.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{})
	} else {
		statefulSets, err = kh.K8sClient.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return statefulSets, nil
}

func (kh *K8sHandler) GetStatefulSet(statefulSetName, namespace string) (*appsv1.StatefulSet, error) {
	if statefulSetName == "" {
		return nil, errors.New("statefulset name must be provided")
	}

	statefulSet, err := kh.K8sClient.AppsV1().StatefulSets(namespace).Get(context.TODO(), statefulSetName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return statefulSet, nil
}
