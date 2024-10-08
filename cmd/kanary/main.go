package main

import (
	"net/http"

	"github.com/kanaryorgs/kanary-server/pkg/k8s"
	"github.com/kanaryorgs/kanary-server/pkg/router"
	"github.com/rs/cors"
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

// @host		localhost:8080
// @BasePath	/v1
func main() {
	r := router.NewRouter(k8sHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Frontend origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	// Wrap the router with CORS
	handler := c.Handler(r)

	// Start the server with CORS handler
	http.ListenAndServe(":8080", handler)
}
