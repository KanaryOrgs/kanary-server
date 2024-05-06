package serializer

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

type NodeList struct {
	Name       string
	Status     string
	IP         string
	CPU        string
	Memory     string
	Conditions []string
	Labels     map[string]string
}

func SerializeNodeList(nodeList *v1.NodeList) []NodeList {
	if nodeList == nil {
		return nil
	}

	serializedNodeList := make([]NodeList, len(nodeList.Items))
	for i, node := range nodeList.Items {
		cpu := node.Status.Capacity[v1.ResourceCPU]
		memory := node.Status.Capacity[v1.ResourceMemory]
		ip := node.Status.Addresses[0].Address
		status := "Unknown"
		for _, condition := range node.Status.Conditions {
			if condition.Type == v1.NodeReady && condition.Status == v1.ConditionTrue {
				status = "Ready"
				break
			} else if condition.Type == v1.NodeReady && condition.Status == v1.ConditionFalse {
				status = "NotReady"
				break
			}
		}

		serializedNodeList[i] = NodeList{
			Name:       node.Name,
			Status:     status,
			IP:         ip,
			CPU:        cpu.String(),
			Memory:     memory.String(),
			Conditions: getNodeConditionString(node.Status.Conditions),
			Labels:     node.Labels,
		}
	}
	return serializedNodeList
}

func getNodeConditionString(conditions []v1.NodeCondition) []string {
	var conditionStrings []string
	for _, condition := range conditions {
		conditionStrings = append(conditionStrings, fmt.Sprintf("%s=%s", condition.Type, condition.Status))
	}
	return conditionStrings
}
