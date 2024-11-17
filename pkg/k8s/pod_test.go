package k8s

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListPods(t *testing.T) {
	tests := []struct {
		name           string
		namespace      string
		initialObjects []v1.Pod
		expectedCount  int
	}{
		{
			name:      "Single Pod in Namespace",
			namespace: "default",
			initialObjects: []v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod-1",
						Namespace: "default",
					},
				},
			},
			expectedCount: 1,
		},
		{
			name:      "Multiple Pods in Namespace",
			namespace: "default",
			initialObjects: []v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod-1",
						Namespace: "default",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod-2",
						Namespace: "default",
					},
				},
			},
			expectedCount: 2,
		},
		{
			name:           "No Pods in Namespace",
			namespace:      "default",
			initialObjects: []v1.Pod{},
			expectedCount:  0,
		},
		{
			name:      "Pods in Different Namespace",
			namespace: "other-namespace",
			initialObjects: []v1.Pod{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-pod-1",
						Namespace: "default",
					},
				},
			},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset()
			for _, obj := range tt.initialObjects {
				_, _ = fakeClient.CoreV1().Pods(obj.Namespace).Create(nil, &obj, metav1.CreateOptions{})
			}

			k8sHandler := K8sHandler{K8sClient: fakeClient}

			pods, err := k8sHandler.ListPods(tt.namespace)

			if err != nil {
				t.Fatalf("ListPods failed: %v", err)
			}

			if len(pods.Items) != tt.expectedCount {
				t.Errorf("expected %d pods, got %d", tt.expectedCount, len(pods.Items))
			}
		})
	}
}
