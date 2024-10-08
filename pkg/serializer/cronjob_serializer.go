package serializer

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CronJobList struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Schedule  string            `json:"schedule"`
	Labels    map[string]string `json:"labels"`
}

type CronJobDetails struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Schedule     string            `json:"schedule"`
	Labels       map[string]string `json:"labels"`
	CreationTime *metav1.Time      `json:"creation_time"`
}

func SerializeCronJobList(cronJobList *batchv1.CronJobList) []CronJobList {
	if cronJobList == nil {
		return nil
	}

	serializedCronJobList := make([]CronJobList, len(cronJobList.Items))
	for i, cj := range cronJobList.Items {
		serializedCronJobList[i] = CronJobList{
			Name:      cj.Name,
			Namespace: cj.Namespace,
			Schedule:  cj.Spec.Schedule,
			Labels:    cj.Labels,
		}
	}
	return serializedCronJobList
}

func SerializeCronJobDetails(cj *batchv1.CronJob) CronJobDetails {
	return CronJobDetails{
		Name:         cj.Name,
		Namespace:    cj.Namespace,
		Schedule:     cj.Spec.Schedule,
		Labels:       cj.Labels,
		CreationTime: &cj.CreationTimestamp,
	}
}
