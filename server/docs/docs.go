// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-09-15 19:54:18.993379341 +0300 MSK m=+0.049467782

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is POA History swagger documentation",
        "title": "Swagger History API",
        "contact": {
            "name": "API Support",
            "email": "nk@bankexfoundation.org"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "1.0"
    },
    "host": "history.bankex.team:8080",
    "basePath": "/",
    "paths": {
        "/a/new/{assetId}/{hash}": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "add hash by assetId",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new asset to assetId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "assetId",
                        "name": "assetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hash of file",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.CreateResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.CreateResponseError"
                            }
                        }
                    }
                }
            }
        },
        "/a/update/{assetId}/{hash}": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "add hash by assetId",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add new asset to assetId (if assetId exists)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "assetId",
                        "name": "assetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Hash of file",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.UpdateResponse"
                            }
                        }
                    }
                }
            }
        },
        "/get/{assetId}/{txNumber}": {
            "get": {
                "description": "Lists all assets by assetId",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return asset by assetId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "assetId",
                        "name": "assetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "txNumber",
                        "name": "txNumber",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.AssetsResponse"
                            }
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "Lists all assets",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Give info about all assets and all meta information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ListResponse"
                            }
                        }
                    }
                }
            }
        },
        "/proof/{assetId}/{txNumber}/{hash}/{timestamp}": {
            "get": {
                "description": "Merkle proof for assetId, txNumber, hash, timestamp (Actually send a JSON File with two arrays Data and Info",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get total Merkle proof",
                "parameters": [
                    {
                        "type": "string",
                        "description": "assetId",
                        "name": "assetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "txNumber",
                        "name": "txNumber",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "hash",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "timestamp",
                        "name": "timestamp",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "test it",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.AssetsResponse": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "type": "byte"
                    },
                    "example": [
                        "1c8d54df80c03a56b5470d164c49f823108f96a67d020e4c677810c9a10b1ca7"
                    ]
                }
            }
        },
        "responses.CreateResponse": {
            "type": "object",
            "properties": {
                "assetId": {
                    "type": "string",
                    "example": "a"
                },
                "hash": {
                    "type": "string",
                    "example": "96e75810b7fe519dd92f6a3f72170b00c0a8a9553f9c765a3cc681eaf7eeab38"
                },
                "merkleRoot": {
                    "type": "array",
                    "items": {
                        "type": "byte"
                    },
                    "example": [
                        "Vu14mZ91jlhkqHhjFwmgjXgxyhLjLADVQlqMSQA3Q3o="
                    ]
                },
                "timestamp": {
                    "type": "integer",
                    "example": 1536920750859
                },
                "txNumber": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "responses.CreateResponseError": {
            "type": "object",
            "properties": {
                "Answer": {
                    "type": "string",
                    "example": "This assetId is already created"
                }
            }
        },
        "responses.ListResponse": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "5b869ee5ca2985e06552a49d"
                },
                "assetId": {
                    "type": "string",
                    "example": "a"
                },
                "assets": {
                    "type": "object"
                },
                "assetsTimestamp": {
                    "type": "object"
                },
                "created_on": {
                    "type": "integer",
                    "example": 1535549157514
                },
                "data": {
                    "type": "string"
                },
                "hash": {
                    "type": "array",
                    "items": {
                        "type": "byte"
                    },
                    "example": [
                        "qNCllA0uMdgEPSVQBYzD4JESEECY2NyjbJgGjy0NP6c="
                    ]
                },
                "txNumber": {
                    "type": "integer",
                    "example": 0
                },
                "updated_on": {
                    "type": "integer",
                    "example": 1535549157514
                }
            }
        },
        "responses.UpdateResponse": {
            "type": "object",
            "properties": {
                "assetId": {
                    "type": "string",
                    "example": "a"
                },
                "timestamp": {
                    "type": "integer",
                    "example": 1536920750859
                },
                "txNumber": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
