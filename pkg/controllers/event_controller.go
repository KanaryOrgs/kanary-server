package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/serializer"
)

type EventController struct {
	kh *k8s.K8sHandler
}

func NewEventController(kh *k8s.K8sHandler) *EventController {
	return &EventController{kh: kh}
}

func (ec *EventController) GetEvents(c *gin.Context) {
	events, err := ec.kh.ListEvents(c.Query("namespace"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get events: %v", err)})
		return
	}
	serializedEvents := serializer.SerializeEventList(events)
	c.JSON(http.StatusOK, serializedEvents)
}

func (ec *EventController) GetEvent(c *gin.Context) {
	eventName := c.Param("name")
	namespace := c.Param("namespace")

	event, err := ec.kh.GetEvent(eventName, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get event: %v", err)})
		return
	}

	serializedEvent := serializer.SerializeEventDetails(event)
	c.JSON(http.StatusOK, serializedEvent)
}
