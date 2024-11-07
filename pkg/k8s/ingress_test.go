package k8s

import (
	"testing"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListIngresses(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-ingress-1",
				Namespace: "default",
			},
		},
		&v1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-ingress-2",
				Namespace: "other-namespace",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	ingresses, err := k8sHandler.ListIngresses("default")
	if err != nil {
		t.Fatalf("ListIngresses failed: %v", err)
	}

	if len(ingresses.Items) != 1 {
		t.Errorf("expected 1 ingress in default namespace, got %d", len(ingresses.Items))
	}

	if ingresses.Items[0].Name != "test-ingress-1" {
		t.Errorf("expected ingress name 'test-ingress-1', got %s", ingresses.Items[0].Name)
	}

	allIngresses, err := k8sHandler.ListIngresses("")
	if err != nil {
		t.Fatalf("ListIngresses with empty namespace failed: %v", err)
	}

	if len(allIngresses.Items) != 2 {
		t.Errorf("expected 2 ingresses across all namespaces, got %d", len(allIngresses.Items))
	}
}
