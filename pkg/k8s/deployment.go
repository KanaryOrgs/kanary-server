package k8s

import (
	"context"
	"errors"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListDeployments lists all Deployments in a namespace.
func (kh *K8sHandler) ListDeployments(namespace string) (*appsv1.DeploymentList, error) {
	var deployments *appsv1.DeploymentList
	var err error

	if namespace == "" {
		deployments, err = kh.K8sClient.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	} else {
		deployments, err = kh.K8sClient.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return deployments, nil
}

// GetDeployment retrieves a single Deployment by name within a given namespace.
func (kh *K8sHandler) GetDeployment(deploymentName, namespace string) (*appsv1.Deployment, error) {
	if deploymentName == "" {
		return nil, errors.New("deployment name must be provided")
	}

	deployment, err := kh.K8sClient.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}
