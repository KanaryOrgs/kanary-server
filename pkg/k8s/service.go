package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kh *K8sHandler) ListServices(namespace string) (*v1.ServiceList, error) {
	var services *v1.ServiceList
	var err error

	if namespace == "" {
		services, err = kh.K8sClient.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	} else {
		services, err = kh.K8sClient.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return services, nil
}
