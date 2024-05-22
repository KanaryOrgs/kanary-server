package serializer

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type NodeList struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	IP             string `json:"ip"`
	CpuCore        int64  `json:"cpu_core"`
	RamCapacity    int64  `json:"ram_capacity"`
	OS             string `json:"os"`
	KubeletVersion string `json:"kubelet_version"`
}

type NodeDetails struct {
	Name           string            `json:"name"`
	Status         string            `json:"status"`
	IP             string            `json:"ip"`
	CpuCore        int64             `json:"cpu_core"`
	RamCapacity    int64             `json:"ram_capacity"`
	OS             string            `json:"os"`
	KubeletVersion string            `json:"kubelet_version"`
	Conditions     []string          `json:"conditions"`
	Labels         map[string]string `json:"labels"`
	Allocatable    map[string]string `json:"allocatable"`
	Capacity       map[string]string `json:"capacity"`
}

func SerializeNodeList(nodeList *v1.NodeList) []NodeList {
	if nodeList == nil {
		return nil
	}

	serializedNodeList := make([]NodeList, len(nodeList.Items))
	for i, node := range nodeList.Items {
		serializedNodeList[i] = NodeList{
			Name:           node.Name,
			Status:         getNodeStatus(node.Status.Conditions),
			IP:             getNodeIP(node.Status.Addresses),
			CpuCore:        getQuantityValue(node.Status.Capacity[v1.ResourceCPU]),
			RamCapacity:    getQuantityValue(node.Status.Capacity[v1.ResourceMemory]) / (1024 * 1024 * 1024),
			OS:             node.Status.NodeInfo.OSImage,
			KubeletVersion: node.Status.NodeInfo.KubeletVersion,
		}
	}
	return serializedNodeList
}

func SerializeNodeDetails(node *v1.Node) NodeDetails {
	if node == nil {
		return NodeDetails{}
	}

	allocatable := make(map[string]string)
	for k, v := range node.Status.Allocatable {
		allocatable[string(k)] = v.String()
	}

	capacity := make(map[string]string)
	for k, v := range node.Status.Capacity {
		capacity[string(k)] = v.String()
	}

	return NodeDetails{
		Name:           node.Name,
		Status:         getNodeStatus(node.Status.Conditions),
		IP:             getNodeIP(node.Status.Addresses),
		CpuCore:        getQuantityValue(node.Status.Capacity[v1.ResourceCPU]),
		RamCapacity:    getQuantityValue(node.Status.Capacity[v1.ResourceMemory]) / (1024 * 1024 * 1024),
		OS:             node.Status.NodeInfo.OSImage,
		KubeletVersion: node.Status.NodeInfo.KubeletVersion,
		Conditions:     getNodeConditionString(node.Status.Conditions),
		Labels:         node.Labels,
		Allocatable:    allocatable,
		Capacity:       capacity,
	}
}

func getNodeStatus(conditions []v1.NodeCondition) string {
	for _, condition := range conditions {
		if condition.Type == v1.NodeReady {
			if condition.Status == v1.ConditionTrue {
				return "Ready"
			}
			return "NotReady"
		}
	}
	return "Unknown"
}

func getNodeIP(addresses []v1.NodeAddress) string {
	for _, addr := range addresses {
		if addr.Type == v1.NodeInternalIP {
			return addr.Address
		}
	}
	return ""
}

func getNodeConditionString(conditions []v1.NodeCondition) []string {
	var conditionStrings []string
	for _, condition := range conditions {
		conditionStrings = append(conditionStrings, fmt.Sprintf("%s=%s", condition.Type, condition.Status))
	}
	return conditionStrings
}

func getQuantityValue(quantity resource.Quantity) int64 {
	return quantity.Value()
}
