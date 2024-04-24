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

type K8sHandler struct {
	K8sClient       *kubernetes.Clientset
	MetricK8sClient *versioned.Clientset
}

func NewK8sHandler() *K8sHandler {
	kh := &K8sHandler{
		K8sClient:       initK8sClient(),
		MetricK8sClient: initMetricK8sClient(),
	}
	return kh
}

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

// initK8sClient initializes a Kubernetes client using in-cluster config, or falls back to out-of-cluster config.
func initK8sClient() *kubernetes.Clientset {
	config, err := getClientConfig()
	if err != nil {
		log.Printf("Failed to get K8s client config: %v", err)
		panic(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to initialize K8s client: %v", err)
		panic(err)
	}

	return client
}

// initMetricK8sClient initializes a Kubernetes metrics client using in-cluster or out-of-cluster config.
func initMetricK8sClient() *versioned.Clientset {
	config, err := getClientConfig()
	if err != nil {
		log.Printf("Failed to get K8s metrics client config: %v", err)
		panic(err)
	}

	client, err := versioned.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to initialize K8s metrics client: %v", err)
		panic(err)
	}

	return client
}
