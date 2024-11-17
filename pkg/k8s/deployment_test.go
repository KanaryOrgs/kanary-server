package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestListDeployments(t *testing.T) {
	tests := []struct {
		name           string
		namespace      string
		initialObjects []appsv1.Deployment
		expectedCount  int
		expectError    bool
	}{
		{
			name:      "Single Deployment in Namespace",
			namespace: "default",
			initialObjects: []appsv1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-1",
						Namespace: "default",
					},
				},
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name:      "Multiple Deployments in Namespace",
			namespace: "default",
			initialObjects: []appsv1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-1",
						Namespace: "default",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-2",
						Namespace: "default",
					},
				},
			},
			expectedCount: 2,
			expectError:   false,
		},
		{
			name:      "No Deployments in Namespace",
			namespace: "default",
			initialObjects: []appsv1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-1",
						Namespace: "other-namespace",
					},
				},
			},
			expectedCount: 0,
			expectError:   false,
		},
		{
			name:           "Empty Namespace",
			namespace:      "",
			initialObjects: []appsv1.Deployment{},
			expectedCount:  0,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset()
			for _, obj := range tt.initialObjects {
				_, _ = fakeClient.AppsV1().Deployments(obj.Namespace).Create(
					nil,
					&obj,
					metav1.CreateOptions{},
				)
			}

			k8sHandler := K8sHandler{
				K8sClient: fakeClient,
			}

			deployments, err := k8sHandler.ListDeployments(tt.namespace)

			if (err != nil) != tt.expectError {
				t.Fatalf("expected error: %v, got: %v", tt.expectError, err)
			}

			if len(deployments.Items) != tt.expectedCount {
				t.Errorf("expected %d deployments, got %d", tt.expectedCount, len(deployments.Items))
			}
		})
	}
}

func TestGetDeployment(t *testing.T) {
	tests := []struct {
		name           string
		deploymentName string
		namespace      string
		initialObjects []appsv1.Deployment
		expectError    bool
	}{
		{
			name:           "Existing Deployment",
			deploymentName: "test-deployment-1",
			namespace:      "default",
			initialObjects: []appsv1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-1",
						Namespace: "default",
					},
				},
			},
			expectError: false,
		},
		{
			name:           "Non-Existing Deployment",
			deploymentName: "nonexistent-deployment",
			namespace:      "default",
			initialObjects: []appsv1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-1",
						Namespace: "default",
					},
				},
			},
			expectError: true,
		},
		{
			name:           "Empty Deployment Name",
			deploymentName: "",
			namespace:      "default",
			initialObjects: []appsv1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-1",
						Namespace: "default",
					},
				},
			},
			expectError: true,
		},
		{
			name:           "Wrong Namespace",
			deploymentName: "test-deployment-1",
			namespace:      "other-namespace",
			initialObjects: []appsv1.Deployment{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment-1",
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
				_, _ = fakeClient.AppsV1().Deployments(obj.Namespace).Create(
					nil,
					&obj,
					metav1.CreateOptions{},
				)
			}

			k8sHandler := K8sHandler{
				K8sClient: fakeClient,
			}

			deployment, err := k8sHandler.GetDeployment(tt.deploymentName, tt.namespace)

			if (err != nil) != tt.expectError {
				t.Fatalf("expected error: %v, got: %v", tt.expectError, err)
			}

			if err == nil && deployment.Name != tt.deploymentName {
				t.Errorf("expected deployment name: %s, got: %s", tt.deploymentName, deployment.Name)
			}
		})
	}
}
