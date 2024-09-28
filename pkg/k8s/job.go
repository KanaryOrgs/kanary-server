package k8s

import (
	"context"
	"errors"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kh *K8sHandler) ListJobs(namespace string) (*batchv1.JobList, error) {
	var jobs *batchv1.JobList
	var err error

	if namespace == "" {
		jobs, err = kh.K8sClient.BatchV1().Jobs("").List(context.TODO(), metav1.ListOptions{})
	} else {
		jobs, err = kh.K8sClient.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (kh *K8sHandler) GetJob(jobName, namespace string) (*batchv1.Job, error) {
	if jobName == "" {
		return nil, errors.New("job name must be provided")
	}

	job, err := kh.K8sClient.BatchV1().Jobs(namespace).Get(context.TODO(), jobName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return job, nil
}
