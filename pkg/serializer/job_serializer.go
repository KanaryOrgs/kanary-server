package serializer

import (
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobList struct {
	Name        string            `json:"name"`
	Namespace   string            `json:"namespace"`
	Completions int32             `json:"completions"`
	Labels      map[string]string `json:"labels"`
}

type JobDetails struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Completions  int32             `json:"completions"`
	Active       int32             `json:"active"`
	Failed       int32             `json:"failed"`
	Labels       map[string]string `json:"labels"`
	CreationTime *metav1.Time      `json:"creation_time"`
}

func SerializeJobList(jobList *batchv1.JobList) []JobList {
	if jobList == nil {
		return nil
	}

	serializedJobList := make([]JobList, len(jobList.Items))
	for i, job := range jobList.Items {
		serializedJobList[i] = JobList{
			Name:        job.Name,
			Namespace:   job.Namespace,
			Completions: *job.Spec.Completions,
			Labels:      job.Labels,
		}
	}
	return serializedJobList
}

func SerializeJobDetails(job *batchv1.Job) JobDetails {
	return JobDetails{
		Name:         job.Name,
		Namespace:    job.Namespace,
		Completions:  *job.Spec.Completions,
		Active:       job.Status.Active,
		Failed:       job.Status.Failed,
		Labels:       job.Labels,
		CreationTime: &job.CreationTimestamp,
	}
}
