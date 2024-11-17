package k8s

import (
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetDaemonSets(t *testing.T) {
	tests := []struct {
		name           string
		initialObjects []appsv1.DaemonSet
		expectedCount  int
	}{
		{
			name: "Single DaemonSet in Namespace",
			initialObjects: []appsv1.DaemonSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-daemonset-1",
						Namespace: "default",
					},
				},
			},
			expectedCount: 1,
		},
		{
			name: "Multiple DaemonSets in Namespace",
			initialObjects: []appsv1.DaemonSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-daemonset-1",
						Namespace: "default",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-daemonset-2",
						Namespace: "default",
					},
				},
			},
			expectedCount: 2,
		},
		{
			name:           "No DaemonSets in Namespace",
			initialObjects: []appsv1.DaemonSet{},
			expectedCount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset()
			for _, obj := range tt.initialObjects {
				_, _ = fakeClient.AppsV1().DaemonSets(obj.Namespace).Create(nil, &obj, metav1.CreateOptions{})
			}

			k8sHandler := K8sHandler{K8sClient: fakeClient}

			daemonSets, err := k8sHandler.GetDaemonSets()

			if err != nil {
				t.Fatalf("GetDaemonSets failed: %v", err)
			}

			if len(daemonSets.Items) != tt.expectedCount {
				t.Errorf("expected %d daemonsets, got %d", tt.expectedCount, len(daemonSets.Items))
			}
		})
	}
}

func TestGetDaemonSet(t *testing.T) {
	tests := []struct {
		name           string
		namespace      string
		daemonSetName  string
		initialObjects []appsv1.DaemonSet
		expectError    bool
	}{
		{
			name:          "Existing DaemonSet",
			namespace:     "default",
			daemonSetName: "test-daemonset-1",
			initialObjects: []appsv1.DaemonSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-daemonset-1",
						Namespace: "default",
					},
				},
			},
			expectError: false,
		},
		{
			name:          "Non-Existing DaemonSet",
			namespace:     "default",
			daemonSetName: "nonexistent-daemonset",
			initialObjects: []appsv1.DaemonSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-daemonset-1",
						Namespace: "default",
					},
				},
			},
			expectError: true,
		},
		{
			name:          "Wrong Namespace",
			namespace:     "other-namespace",
			daemonSetName: "test-daemonset-1",
			initialObjects: []appsv1.DaemonSet{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-daemonset-1",
						Namespace: "default",
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			fakeClient := fake.NewSimpleClientset()
			for _, obj := range tt.initialObjects {
				_, _ = fakeClient.AppsV1().DaemonSets(obj.Namespace).Create(nil, &obj, metav1.CreateOptions{})
			}

			k8sHandler := K8sHandler{K8sClient: fakeClient}

			daemonSet, err := k8sHandler.GetDaemonSet(tt.namespace, tt.daemonSetName)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("GetDaemonSet failed: %v", err)
				}

				if daemonSet.Name != tt.daemonSetName {
					t.Errorf("expected daemonset name '%s', got '%s'", tt.daemonSetName, daemonSet.Name)
				}
			}
		})
	}
}
