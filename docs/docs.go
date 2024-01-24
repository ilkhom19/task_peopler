// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/people": {
            "post": {
                "description": "Create person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person CRUD"
                ],
                "summary": "Create person",
                "parameters": [
                    {
                        "description": "Person",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task_peopler_internal_models.PersonCreateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task_peopler_internal_models.Person"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            }
        },
        "/people/age": {
            "get": {
                "description": "Get people by age",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person Filter"
                ],
                "summary": "Get people by age",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person Age",
                        "name": "age",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_transport_rest.PeopleResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            }
        },
        "/people/gender": {
            "get": {
                "description": "Get people by gender",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person Filter"
                ],
                "summary": "Get people by gender",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person Gender",
                        "name": "gender",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_transport_rest.PeopleResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            }
        },
        "/people/name": {
            "get": {
                "description": "Get people by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person Filter"
                ],
                "summary": "Get people by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_transport_rest.PeopleResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            }
        },
        "/people/nationality": {
            "get": {
                "description": "Get people by nationality",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person Filter"
                ],
                "summary": "Get people by nationality",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person Nationality",
                        "name": "nationality",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_transport_rest.PeopleResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            }
        },
        "/people/{id}": {
            "get": {
                "description": "Get person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person CRUD"
                ],
                "summary": "Get person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task_peopler_internal_models.Person"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person CRUD"
                ],
                "summary": "Update person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PersonUpdateInput",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/task_peopler_internal_models.PersonUpdateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/task_peopler_internal_models.Person"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "person CRUD"
                ],
                "summary": "Delete person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_transport_rest.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal_transport_rest.PeopleResponse": {
            "type": "object",
            "properties": {
                "people": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/task_peopler_internal_models.Person"
                    }
                }
            }
        },
        "internal_transport_rest.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "task_peopler_internal_models.Person": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                }
            }
        },
        "task_peopler_internal_models.PersonCreateInput": {
            "type": "object",
            "required": [
                "first_name",
                "last_name"
            ],
            "properties": {
                "first_name": {
                    "type": "string",
                    "minLength": 2
                },
                "last_name": {
                    "type": "string",
                    "minLength": 2
                },
                "patronymic": {
                    "type": "string"
                }
            }
        },
        "task_peopler_internal_models.PersonUpdateInput": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "first_name": {
                    "type": "string",
                    "minLength": 2
                },
                "gender": {
                    "type": "string",
                    "minLength": 2
                },
                "last_name": {
                    "type": "string",
                    "minLength": 2
                },
                "nationality": {
                    "type": "string",
                    "minLength": 2
                },
                "patronymic": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.1",
	Host:             "http://localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "People-Base API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
