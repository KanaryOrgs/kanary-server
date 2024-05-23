package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kh *K8sHandler) ListNodes() (*v1.NodeList, error) {
	nodeList, err := kh.K8sClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return nodeList, nil
}

func (kh *K8sHandler) GetNode(nodeName string) (*v1.Node, error) {
	node, err := kh.K8sClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return node, nil
}
