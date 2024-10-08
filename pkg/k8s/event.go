package k8s

import (
	"context"
	"errors"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListEvents retrieves all events in a namespace.
func (kh *K8sHandler) ListEvents(namespace string) (*v1.EventList, error) {
	var events *v1.EventList
	var err error

	if namespace == "" {
		events, err = kh.K8sClient.CoreV1().Events("").List(context.TODO(), metav1.ListOptions{})
	} else {
		events, err = kh.K8sClient.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{})
	}

	if err != nil {
		return nil, err
	}

	return events, nil
}

// GetEvent retrieves a specific event by name in a given namespace.
func (kh *K8sHandler) GetEvent(eventName, namespace string) (*v1.Event, error) {
	if eventName == "" {
		return nil, errors.New("event name must be provided")
	}

	event, err := kh.K8sClient.CoreV1().Events(namespace).Get(context.TODO(), eventName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return event, nil
}
