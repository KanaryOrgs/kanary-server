package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestListNodes(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-node-1",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	nodes, err := k8sHandler.ListNodes()
	if err != nil {
		t.Fatalf("ListNodes failed: %v", err)
	}

	if len(nodes.Items) != 1 {
		t.Errorf("expected 1 node, got %d", len(nodes.Items))
	}
	if nodes.Items[0].Name != "test-node-1" {
		t.Errorf("expected node name 'test-node-1', got %s", nodes.Items[0].Name)
	}
}

func TestGetNode(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-node-1",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	node, err := k8sHandler.GetNode("test-node-1")
	if err != nil {
		t.Fatalf("GetNode failed: %v", err)
	}

	if node.Name != "test-node-1" {
		t.Errorf("expected node name 'test-node-1', got %s", node.Name)
	}
}
