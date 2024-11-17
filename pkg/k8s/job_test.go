package k8s

import (
	"testing"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListJobs(t *testing.T) {
	tests := []struct {
		name           string
		namespace      string
		initialObjects []batchv1.Job
		expectedCount  int
	}{
		{
			name:      "Single Job in Namespace",
			namespace: "default",
			initialObjects: []batchv1.Job{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-1",
						Namespace: "default",
					},
				},
			},
			expectedCount: 1,
		},
		{
			name:      "Multiple Jobs in Namespace",
			namespace: "default",
			initialObjects: []batchv1.Job{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-1",
						Namespace: "default",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-2",
						Namespace: "default",
					},
				},
			},
			expectedCount: 2,
		},
		{
			name:           "No Jobs in Namespace",
			namespace:      "default",
			initialObjects: []batchv1.Job{},
			expectedCount:  0,
		},
		{
			name:      "Jobs in All Namespaces",
			namespace: "",
			initialObjects: []batchv1.Job{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-1",
						Namespace: "default",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-2",
						Namespace: "kube-system",
					},
				},
			},
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeClient := fake.NewSimpleClientset()
			for _, obj := range tt.initialObjects {
				_, _ = fakeClient.BatchV1().Jobs(obj.Namespace).Create(nil, &obj, metav1.CreateOptions{})
			}

			k8sHandler := K8sHandler{K8sClient: fakeClient}

			jobs, err := k8sHandler.ListJobs(tt.namespace)

			if err != nil {
				t.Fatalf("ListJobs failed: %v", err)
			}

			if len(jobs.Items) != tt.expectedCount {
				t.Errorf("expected %d jobs, got %d", tt.expectedCount, len(jobs.Items))
			}
		})
	}
}

func TestGetJob(t *testing.T) {
	tests := []struct {
		name           string
		jobName        string
		namespace      string
		initialObjects []batchv1.Job
		expectError    bool
	}{
		{
			name:      "Existing Job",
			jobName:   "test-job-1",
			namespace: "default",
			initialObjects: []batchv1.Job{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-1",
						Namespace: "default",
					},
				},
			},
			expectError: false,
		},
		{
			name:      "Non-Existing Job",
			jobName:   "nonexistent-job",
			namespace: "default",
			initialObjects: []batchv1.Job{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-1",
						Namespace: "default",
					},
				},
			},
			expectError: true,
		},
		{
			name:      "Empty Job Name",
			jobName:   "",
			namespace: "default",
			initialObjects: []batchv1.Job{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-1",
						Namespace: "default",
					},
				},
			},
			expectError: true,
		},
		{
			name:      "Wrong Namespace",
			jobName:   "test-job-1",
			namespace: "other-namespace",
			initialObjects: []batchv1.Job{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-job-1",
						Namespace: "default",
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
				_, _ = fakeClient.BatchV1().Jobs(obj.Namespace).Create(nil, &obj, metav1.CreateOptions{})
			}

			k8sHandler := K8sHandler{K8sClient: fakeClient}

			job, err := k8sHandler.GetJob(tt.jobName, tt.namespace)

			if tt.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("GetJob failed: %v", err)
				}

				if job.Name != tt.jobName {
					t.Errorf("expected job name '%s', got '%s'", tt.jobName, job.Name)
				}
			}
		})
	}
}
