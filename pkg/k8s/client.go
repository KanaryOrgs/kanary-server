package k8s

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
	"os"
)

// getClientConfig tries to get in-cluster config and falls back to out-of-cluster config if it fails.
func getClientConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	k8sAPIAddr := os.Getenv("K8S_API_LISTEN_ADDR")
	k8sAPIPort := os.Getenv("K8S_API_LISTEN_PORT")
	k8sConfig := os.Getenv("K8S_API_CONFIG")

	if k8sAPIAddr == "" || k8sAPIPort == "" || k8sConfig == "" {
		return nil, fmt.Errorf("incomplete k8s API environment settings: Address=%s, Port=%s, Config=%s", k8sAPIAddr, k8sAPIPort, k8sConfig)
	}

	return clientcmd.BuildConfigFromFlags("https://"+k8sAPIAddr+":"+k8sAPIPort, k8sConfig)
}

// InitK8sClient initializes a Kubernetes client using in-cluster config, or falls back to out-of-cluster config.
func InitK8sClient() (*kubernetes.Clientset, error) {
	config, err := getClientConfig()
	if err != nil {
		log.Printf("Failed to get K8s client config: %v", err)
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to initialize K8s client: %v", err)
		return nil, err
	}

	return client, nil
}

// InitMetricK8sClient initializes a Kubernetes metrics client using in-cluster or out-of-cluster config.
func InitMetricK8sClient() (*versioned.Clientset, error) {
	config, err := getClientConfig()
	if err != nil {
		log.Printf("Failed to get K8s metrics client config: %v", err)
		return nil, err
	}

	client, err := versioned.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to initialize K8s metrics client: %v", err)
		return nil, err
	}

	return client, nil
}
