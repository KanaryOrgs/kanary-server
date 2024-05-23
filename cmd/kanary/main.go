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

//	@title			kanary-server API
//	@version		1.0
//	@description	This is API document for kanary-server
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1
func main() {
	r := router.NewRouter(k8sHandler)
	r.Run(":" + "8080")
}
