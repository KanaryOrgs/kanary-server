package serializer

import (
	v1 "k8s.io/api/core/v1"
)

type ServiceList struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	ClusterIP string            `json:"clusterIP"`
	Ports     []ServicePort     `json:"ports"`
	Labels    map[string]string `json:"labels"`
	Selector  map[string]string `json:"selector"`
}

type ServicePort struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     int32  `json:"port"`
}

// SerializeServiceList serializes a ServiceList to a slice of ServiceList structures.
func SerializeServiceList(serviceList *v1.ServiceList) []ServiceList {
	if serviceList == nil {
		return nil
	}

	serializedServices := make([]ServiceList, len(serviceList.Items))
	for i, svc := range serviceList.Items {
		ports := make([]ServicePort, len(svc.Spec.Ports))
		for j, port := range svc.Spec.Ports {
			ports[j] = ServicePort{
				Name:     port.Name,
				Protocol: string(port.Protocol),
				Port:     port.Port,
			}
		}
		serializedServices[i] = ServiceList{
			Name:      svc.Name,
			Namespace: svc.Namespace,
			ClusterIP: svc.Spec.ClusterIP,
			Ports:     ports,
			Labels:    svc.Labels,
			Selector:  svc.Spec.Selector,
		}
	}
	return serializedServices
}
