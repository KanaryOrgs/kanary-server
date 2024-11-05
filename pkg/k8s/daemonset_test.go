package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestGetDaemonSets(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&appsv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-daemonset-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	daemonSets, err := k8sHandler.GetDaemonSets()
	if err != nil {
		t.Fatalf("GetDaemonSets failed: %v", err)
	}

	if len(daemonSets.Items) != 1 {
		t.Errorf("expected 1 daemonset, got %d", len(daemonSets.Items))
	}
	if daemonSets.Items[0].Name != "test-daemonset-1" {
		t.Errorf("expected daemonset name 'test-daemonset-1', got %s", daemonSets.Items[0].Name)
	}
}

func TestGetDaemonSet(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&appsv1.DaemonSet{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-daemonset-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	daemonSet, err := k8sHandler.GetDaemonSet("default", "test-daemonset-1")
	if err != nil {
		t.Fatalf("GetDaemonSet failed: %v", err)
	}

	if daemonSet.Name != "test-daemonset-1" {
		t.Errorf("expected daemonset name 'test-daemonset-1', got %s", daemonSet.Name)
	}
}
