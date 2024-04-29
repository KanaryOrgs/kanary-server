package k8s

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
	"os"
	"path/filepath"
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

// getClientConfig tries to get in-cluster kubeconfig and falls back to out-of-cluster kubeconfig if it fails.
func getClientConfig() (*rest.Config, error) {
	kubeconfig, err := rest.InClusterConfig()
	if err == nil {
		return kubeconfig, nil
	}

	configFile := os.Getenv("KUBECONFIG")
	if configFile == "" {
		home := os.Getenv("HOME")
		configFile = filepath.Join(home, ".kube", "config")
		if _, err := os.Stat(filepath.Clean(configFile)); err != nil {
			return nil, fmt.Errorf("failed to access kubeconfig file %s: %v", configFile, err)
		}
	}

	kubeconfig, err = clientcmd.BuildConfigFromFlags("", configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to build config from file %s: %v", configFile, err)

	}
	return kubeconfig, nil
}

// initK8sClient initializes a Kubernetes client using in-cluster kubeconfig, or falls back to out-of-cluster kubeconfig.
func initK8sClient() *kubernetes.Clientset {
	config, err := getClientConfig()
	if err != nil {
		log.Printf("Failed to get K8s client kubeconfig: %v", err)
		panic(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to initialize K8s client: %v", err)
		panic(err)
	}

	return client
}

// initMetricK8sClient initializes a Kubernetes metrics client using in-cluster or out-of-cluster kubeconfig.
func initMetricK8sClient() *versioned.Clientset {
	config, err := getClientConfig()
	if err != nil {
		log.Printf("Failed to get K8s metrics client kubeconfig: %v", err)
		panic(err)
	}

	client, err := versioned.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to initialize K8s metrics client: %v", err)
		panic(err)
	}

	return client
}
