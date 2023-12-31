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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/a-items": {
            "get": {
                "description": "Return all stored items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemA"
                ],
                "summary": "Show all items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.ItemA"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "creates an item with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemA"
                ],
                "summary": "Creates an item",
                "parameters": [
                    {
                        "description": "Item Properties",
                        "name": "itemA",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ItemA"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ItemA"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/a-items/{id}": {
            "get": {
                "description": "get item by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemA"
                ],
                "summary": "Show an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ItemA"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Updates an item with given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemA"
                ],
                "summary": "Updates an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Item Properties",
                        "name": "itemA",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ItemA"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Deletes an item with given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemA"
                ],
                "summary": "Deletes an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "string",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ItemA"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/b-items": {
            "get": {
                "description": "Return all stored items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemB"
                ],
                "summary": "Show all items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.ItemB"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "creates an item with given data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemB"
                ],
                "summary": "Creates an item",
                "parameters": [
                    {
                        "description": "Item Properties",
                        "name": "itemB",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ItemB"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ItemB"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/b-items/{id}": {
            "get": {
                "description": "get item by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemB"
                ],
                "summary": "Show an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ItemB"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Updates an item with given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemB"
                ],
                "summary": "Updates an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Item Properties",
                        "name": "itemB",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ItemB"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Deletes an item with given ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "itemB"
                ],
                "summary": "Deletes an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "string",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ItemB"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ItemA": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "domain.ItemB": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8085",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Service A Swagger Example API",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
