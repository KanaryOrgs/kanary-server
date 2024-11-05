package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestListDeployments(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-deployment-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	deployments, err := k8sHandler.ListDeployments("default")
	if err != nil {
		t.Fatalf("ListDeployments failed: %v", err)
	}

	if len(deployments.Items) != 1 {
		t.Errorf("expected 1 deployment, got %d", len(deployments.Items))
	}
	if deployments.Items[0].Name != "test-deployment-1" {
		t.Errorf("expected deployment name 'test-deployment-1', got %s", deployments.Items[0].Name)
	}
}

func TestGetDeployment(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-deployment-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	deployment, err := k8sHandler.GetDeployment("test-deployment-1", "default")
	if err != nil {
		t.Fatalf("GetDeployment failed: %v", err)
	}

	if deployment.Name != "test-deployment-1" {
		t.Errorf("expected deployment name 'test-deployment-1', got %s", deployment.Name)
	}
}
