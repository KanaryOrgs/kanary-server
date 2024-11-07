package k8s

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListServices(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-service-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	services, err := k8sHandler.ListServices("default")
	if err != nil {
		t.Fatalf("ListServices failed: %v", err)
	}

	if len(services.Items) != 1 {
		t.Errorf("expected 1 service, got %d", len(services.Items))
	}
	if services.Items[0].Name != "test-service-1" {
		t.Errorf("expected service name 'test-service-1', got %s", services.Items[0].Name)
	}
}
