package k8s

import (
	"context"
	batchv1 "k8s.io/api/batch/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kh *K8sHandler) GetCronJobs() (*batchv1.CronJobList, error) {
	return kh.K8sClient.BatchV1().CronJobs("").List(context.TODO(), metav1.ListOptions{})
}

func (kh *K8sHandler) GetCronJob(namespace, name string) (*batchv1.CronJob, error) {
	return kh.K8sClient.BatchV1().CronJobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}
