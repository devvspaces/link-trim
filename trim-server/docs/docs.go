// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/:link_uuid": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "redirect to the original link from the uuid",
                "operationId": "redirect-to-original-link",
                "parameters": [
                    {
                        "type": "string",
                        "description": "short uuid",
                        "name": "link_uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "300": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/user/all": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get all the users in the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.TrimmedLink": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "trimmed": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "trimmedLinks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TrimmedLink"
                    }
                },
                "username": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Trim-Sever",
	Description:      "This is a simple link shortner api. You can visit the GitHub repository at https://github.com/Xavier577/link-trim",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}