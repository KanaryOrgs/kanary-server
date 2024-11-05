package k8s

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListPods(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-pod-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	pods, err := k8sHandler.ListPods("default")
	if err != nil {
		t.Fatalf("ListPods failed: %v", err)
	}

	if len(pods.Items) != 1 {
		t.Errorf("expected 1 pod, got %d", len(pods.Items))
	}
	if pods.Items[0].Name != "test-pod-1" {
		t.Errorf("expected pod name 'test-pod-1', got %s", pods.Items[0].Name)
	}
}
