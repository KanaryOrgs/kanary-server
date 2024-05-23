package serializer

import (
	v1 "k8s.io/api/networking/v1"
)

type IngressList struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Host      string            `json:"host"`
	Paths     []string          `json:"paths"`
	Labels    map[string]string `json:"labels"`
}

// SerializeIngressList serializes an IngressList to a slice of IngressList structures.
func SerializeIngressList(ingressList *v1.IngressList) []IngressList {
	if ingressList == nil {
		return nil
	}

	serializedIngresses := make([]IngressList, len(ingressList.Items))
	for i, ingress := range ingressList.Items {
		var paths []string
		for _, rule := range ingress.Spec.Rules {
			for _, path := range rule.HTTP.Paths {
				paths = append(paths, path.Path)
			}
		}
		host := ""
		if len(ingress.Spec.Rules) > 0 {
			host = ingress.Spec.Rules[0].Host
		}
		serializedIngresses[i] = IngressList{
			Name:      ingress.Name,
			Namespace: ingress.Namespace,
			Host:      host,
			Paths:     paths,
			Labels:    ingress.Labels,
		}
	}
	return serializedIngresses
}
