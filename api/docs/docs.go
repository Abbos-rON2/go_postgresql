// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginModel"
                        }
                    }
                ],
                "responses": {
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorModel"
                        }
                    }
                }
            }
        },
        "/teacher": {
            "get": {
                "description": "get all teachers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teacher"
                ],
                "summary": "get all teachers",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Teacher"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorModel"
                        }
                    }
                }
            },
            "post": {
                "description": "create teacher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teacher"
                ],
                "summary": "create teacher",
                "parameters": [
                    {
                        "description": "teacher",
                        "name": "teacher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Teacher"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Teacher"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorModel"
                        }
                    }
                }
            }
        },
        "/teacher/{guid}": {
            "get": {
                "description": "get teacher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teacher"
                ],
                "summary": "get teacher",
                "parameters": [
                    {
                        "type": "string",
                        "description": "guid",
                        "name": "guid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Teacher"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorModel"
                        }
                    }
                }
            },
            "put": {
                "description": "update teacher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teacher"
                ],
                "summary": "update teacher",
                "parameters": [
                    {
                        "description": "teacher",
                        "name": "teacher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateTeacher"
                        }
                    },
                    {
                        "type": "string",
                        "description": "guid",
                        "name": "guid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Teacher"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorModel"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete teacher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teacher"
                ],
                "summary": "delete teacher",
                "parameters": [
                    {
                        "type": "string",
                        "description": "guid",
                        "name": "guid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorModel"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get all users",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.SuccessModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.User"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorModel"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateUser": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "super_admin",
                        " branch_admin"
                    ]
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ErrorModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "error": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.LoginModel": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.SuccessModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Teacher": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "branch_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.UpdateTeacher": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "birth_date": {
                    "type": "string"
                },
                "branch_id": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "super_admin",
                        " branch_admin"
                    ]
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Test Task",
	Description: "This is a sample server celler server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
