package main

import (
	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/router"
)

var (
	k8sHandler *k8s.K8sHandler
)

func init() {
	k8sHandler = k8s.NewK8sHandler()
}

func main() {
	r := router.NewRouter(k8sHandler)
	r.Run(":" + "8080")
}
