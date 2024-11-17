package k8s

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListNodes(t *testing.T) {
	tests := []struct {
		name           string
		initialObjects []v1.Node
		expectedCount  int
	}{
		{
			name: "Single Node",
			initialObjects: []v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-node-1",
					},
				},
			},
			expectedCount: 1,
		},
		{
			name: "Multiple Nodes",
			initialObjects: []v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-node-1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-node-2",
					},
				},
			},
			expectedCount: 2,
		},
		{
			name:           "No Nodes",
			initialObjects: []v1.Node{},
			expectedCount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset()
			for _, obj := range tt.initialObjects {
				_, _ = fakeClient.CoreV1().Nodes().Create(nil, &obj, metav1.CreateOptions{})
			}

			k8sHandler := K8sHandler{K8sClient: fakeClient}

			nodes, err := k8sHandler.ListNodes()

			if err != nil {
				t.Fatalf("ListNodes failed: %v", err)
			}

			if len(nodes.Items) != tt.expectedCount {
				t.Errorf("expected %d nodes, got %d", tt.expectedCount, len(nodes.Items))
			}
		})
	}
}

func TestGetNode(t *testing.T) {
	tests := []struct {
		name           string
		nodeName       string
		initialObjects []v1.Node
		expectError    bool
	}{
		{
			name:     "Existing Node",
			nodeName: "test-node-1",
			initialObjects: []v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-node-1",
					},
				},
			},
			expectError: false,
		},
		{
			name:     "Non-Existing Node",
			nodeName: "nonexistent-node",
			initialObjects: []v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-node-1",
					},
				},
			},
			expectError: true,
		},
		{
			name:     "Empty Node Name",
			nodeName: "",
			initialObjects: []v1.Node{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test-node-1",
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset()
			for _, obj := range tt.initialObjects {
				_, _ = fakeClient.CoreV1().Nodes().Create(nil, &obj, metav1.CreateOptions{})
			}

			k8sHandler := K8sHandler{K8sClient: fakeClient}

			node, err := k8sHandler.GetNode(tt.nodeName)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("GetNode failed: %v", err)
				}

				if node.Name != tt.nodeName {
					t.Errorf("expected node name '%s', got '%s'", tt.nodeName, node.Name)
				}
			}
		})
	}
}
