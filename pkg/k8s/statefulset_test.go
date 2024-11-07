package k8s

import (
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListStatefulSets(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-statefulset-1",
				Namespace: "default",
			},
		},
		&appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-statefulset-2",
				Namespace: "other-namespace",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	statefulSets, err := k8sHandler.ListStatefulSets("default")
	if err != nil {
		t.Fatalf("ListStatefulSets failed: %v", err)
	}

	if len(statefulSets.Items) != 1 {
		t.Errorf("expected 1 StatefulSet in default namespace, got %d", len(statefulSets.Items))
	}

	if statefulSets.Items[0].Name != "test-statefulset-1" {
		t.Errorf("expected StatefulSet name 'test-statefulset-1', got %s", statefulSets.Items[0].Name)
	}

	allStatefulSets, err := k8sHandler.ListStatefulSets("")
	if err != nil {
		t.Fatalf("ListStatefulSets with empty namespace failed: %v", err)
	}

	if len(allStatefulSets.Items) != 2 {
		t.Errorf("expected 2 StatefulSets across all namespaces, got %d", len(allStatefulSets.Items))
	}
}

func TestGetStatefulSet(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&appsv1.StatefulSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-statefulset-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	statefulSet, err := k8sHandler.GetStatefulSet("test-statefulset-1", "default")
	if err != nil {
		t.Fatalf("GetStatefulSet failed: %v", err)
	}

	if statefulSet.Name != "test-statefulset-1" {
		t.Errorf("expected StatefulSet name 'test-statefulset-1', got %s", statefulSet.Name)
	}

	_, err = k8sHandler.GetStatefulSet("", "default")
	if err == nil || err.Error() != "statefulset name must be provided" {
		t.Errorf("expected error 'statefulset name must be provided', got %v", err)
	}
}
