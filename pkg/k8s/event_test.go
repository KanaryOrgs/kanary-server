package k8s

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListEvents(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.Event{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-event-1",
				Namespace: "default",
			},
		},
		&v1.Event{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-event-2",
				Namespace: "other-namespace",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	events, err := k8sHandler.ListEvents("default")
	if err != nil {
		t.Fatalf("ListEvents failed: %v", err)
	}

	if len(events.Items) != 1 {
		t.Errorf("expected 1 event in default namespace, got %d", len(events.Items))
	}

	if events.Items[0].Name != "test-event-1" {
		t.Errorf("expected event name 'test-event-1', got %s", events.Items[0].Name)
	}

	allEvents, err := k8sHandler.ListEvents("")
	if err != nil {
		t.Fatalf("ListEvents with empty namespace failed: %v", err)
	}

	if len(allEvents.Items) != 2 {
		t.Errorf("expected 2 events across all namespaces, got %d", len(allEvents.Items))
	}
}

func TestGetEvent(t *testing.T) {
	fakeClient := fake.NewSimpleClientset(
		&v1.Event{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-event-1",
				Namespace: "default",
			},
		},
	)

	k8sHandler := K8sHandler{
		K8sClient: fakeClient,
	}

	event, err := k8sHandler.GetEvent("test-event-1", "default")
	if err != nil {
		t.Fatalf("GetEvent failed: %v", err)
	}

	if event.Name != "test-event-1" {
		t.Errorf("expected event name 'test-event-1', got %s", event.Name)
	}

	_, err = k8sHandler.GetEvent("", "default")
	if err == nil || err.Error() != "event name must be provided" {
		t.Errorf("expected error 'event name must be provided', got %v", err)
	}
}
