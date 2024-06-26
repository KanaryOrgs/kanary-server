// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ingresses": {
            "get": {
                "description": "get ingress list in k8s cluster.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ingresses"
                ],
                "summary": "Show ingress list.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/serializer.IngressList"
                            }
                        }
                    }
                }
            }
        },
        "/nodes": {
            "get": {
                "description": "get node list in k8s cluster.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "nodes"
                ],
                "summary": "Show node list.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/serializer.NodeList"
                            }
                        }
                    }
                }
            }
        },
        "/pods": {
            "get": {
                "description": "get pod list in k8s cluster.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pods"
                ],
                "summary": "Show pod list.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/serializer.PodList"
                            }
                        }
                    }
                }
            }
        },
        "/services": {
            "get": {
                "description": "get service list in k8s cluster.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "services"
                ],
                "summary": "Show service list.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/serializer.ServiceList"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.IngressList": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "paths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "serializer.NodeList": {
            "type": "object",
            "properties": {
                "conditions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "cpu_core": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "kubelet_version": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "ram_capacity": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "serializer.PodList": {
            "type": "object",
            "properties": {
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ip": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "restarts": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "serializer.ServiceList": {
            "type": "object",
            "properties": {
                "clusterIP": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "ports": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.ServicePort"
                    }
                },
                "selector": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "serializer.ServicePort": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "protocol": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "kanary-server API",
	Description:      "This is API document for kanary-server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
