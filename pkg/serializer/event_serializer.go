package serializer

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EventList struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Reason    string            `json:"reason"`
	Message   string            `json:"message"`
	Labels    map[string]string `json:"labels"`
}

type EventDetails struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Reason       string            `json:"reason"`
	Message      string            `json:"message"`
	Source       string            `json:"source"`
	Labels       map[string]string `json:"labels"`
	CreationTime *metav1.Time      `json:"creation_time"`
}

func SerializeEventList(eventList *v1.EventList) []EventList {
	if eventList == nil {
		return nil
	}

	serializedEventList := make([]EventList, len(eventList.Items))
	for i, event := range eventList.Items {
		serializedEventList[i] = EventList{
			Name:      event.Name,
			Namespace: event.Namespace,
			Reason:    event.Reason,
			Message:   event.Message,
			Labels:    event.Labels,
		}
	}
	return serializedEventList
}

func SerializeEventDetails(event *v1.Event) EventDetails {
	return EventDetails{
		Name:         event.Name,
		Namespace:    event.Namespace,
		Reason:       event.Reason,
		Message:      event.Message,
		Source:       event.Source.Component,
		Labels:       event.Labels,
		CreationTime: &event.CreationTimestamp,
	}
}
