package k8s

import (
	"context"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kh *K8sHandler) ListIngresses(namespace string) (*v1.IngressList, error) {
	var ingresses *v1.IngressList
	var err error

	if namespace == "" {
		ingresses, err = kh.K8sClient.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	} else {
		ingresses, err = kh.K8sClient.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return ingresses, nil
}
