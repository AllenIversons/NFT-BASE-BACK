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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/collections": {
            "get": {
                "description": "get all collections under the current account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collection"
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/collections",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/collections/{collection-id}": {
            "get": {
                "description": "get single collection by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collection"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "collection id",
                        "name": "collection-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/collections/XXXX",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/collections/{collection-id}/items": {
            "get": {
                "description": "all items in collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collection"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "collection id",
                        "name": "collection-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/collections/XXXXX/items",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/create": {
            "post": {
                "description": "create new collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "collection"
                ],
                "responses": {
                    "200": {
                        "description": "POST /api/v1/collections/create",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event-banner": {
            "get": {
                "description": "event-banner",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tour"
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/event-banner",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/events": {
            "get": {
                "description": "all events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/events",
                        "schema": {
                            "$ref": "#/definitions/v1.EventsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/events/{event-id}": {
            "get": {
                "description": "single event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "event-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/events/balala",
                        "schema": {
                            "$ref": "#/definitions/v1.EventResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/events/{event-id}/items": {
            "get": {
                "description": "items in event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "event-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/events/balala/items",
                        "schema": {
                            "$ref": "#/definitions/v1.EventItemsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/events/{event-id}/join": {
            "post": {
                "description": "user join event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "event-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/events/balala/join",
                        "schema": {
                            "$ref": "#/definitions/v1.JoinEventResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/events/{event-id}/likes": {
            "post": {
                "description": "user join event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "event id",
                        "name": "event-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "item id",
                        "name": "item-id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/events/balala/likes",
                        "schema": {
                            "$ref": "#/definitions/v1.JoinEventResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/events/{event-id}/ranks": {
            "get": {
                "description": "item ranks in event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "event-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/events/balala/ranks",
                        "schema": {
                            "$ref": "#/definitions/v1.EventItemsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/events/{event-id}/submit-item": {
            "post": {
                "description": "user submit item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "event"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "event id",
                        "name": "event-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "item id",
                        "name": "item-id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/events/balala/submit-item",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/items": {
            "get": {
                "description": "all items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "favorites,popular,newest",
                        "name": "sort-by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "image,video,audio",
                        "name": "filter",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/items",
                        "schema": {
                            "$ref": "#/definitions/v1.AllItemsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/items/{item-id}": {
            "get": {
                "description": "single item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "collection id",
                        "name": "item-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/items/yiiiiiii",
                        "schema": {
                            "$ref": "#/definitions/v1.Item"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Error"
                        }
                    }
                }
            }
        },
        "/tour": {
            "get": {
                "description": "tour",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tour"
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/tour",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tr/articles": {
            "get": {
                "description": "tutorials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tour"
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/tr/articles",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tr/articles/:articles-id": {
            "post": {
                "description": "view articles by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tour"
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/tr/articles/XXXXXX",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/users",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "campus",
                        "name": "Campus",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "Email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "passwd",
                        "name": "Passwd",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "bannerimage",
                        "name": "BannerImage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "avatarimage",
                        "name": "AvatarImage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "UserName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/users",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "description": "get user:if username == id, return all information includes password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/Anna",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/change-profile": {
            "post": {
                "description": "change profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user Name",
                        "name": "userName",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/Anna/change-profile",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/collected": {
            "get": {
                "description": "collected",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/Anna/collected",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/create-collection": {
            "post": {
                "description": "create-collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "collection Name",
                        "name": "collection-Name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/Anna/create-collection",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/create-item": {
            "post": {
                "description": "create-item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "item Name",
                        "name": "item-Name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/Anna/create-item",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/creation": {
            "get": {
                "description": "creation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/Anna/creation",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/edit-collection": {
            "post": {
                "description": "edit-collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "collection Name",
                        "name": "collection-Name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/Anna/edit-collection",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/edit-item": {
            "post": {
                "description": "edit-item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "item Name",
                        "name": "item-Name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "POST/api/v1/Anna/edit-item",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/{id}/favorites": {
            "get": {
                "description": "favorites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GET/api/v1/Anna/favorites",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "utils.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "USER_NAME_EXIST"
                },
                "message": {
                    "type": "string",
                    "example": "user name exist"
                }
            }
        },
        "v1.AllItemsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "success"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Item"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "v1.Event": {
            "type": "object",
            "properties": {
                "event_description": {
                    "type": "string"
                },
                "event_id": {
                    "type": "integer"
                },
                "event_name": {
                    "type": "string",
                    "example": "helloworld"
                },
                "start_time": {
                    "type": "string"
                },
                "time_duration": {
                    "type": "integer"
                }
            }
        },
        "v1.EventItem": {
            "type": "object",
            "properties": {
                "collection_id": {
                    "type": "integer",
                    "example": 5
                },
                "image": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer",
                    "example": 123455
                },
                "localFavorites": {
                    "type": "integer",
                    "example": 1
                },
                "ownerId": {
                    "type": "string",
                    "example": "mazhengwang-ust-hk"
                }
            }
        },
        "v1.EventItemsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "SUCCESS"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.EventItem"
                    }
                },
                "totel": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "v1.EventResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "event": {
                    "$ref": "#/definitions/v1.Event"
                }
            }
        },
        "v1.EventsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Event"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "v1.Item": {
            "type": "object",
            "properties": {
                "collection_id": {
                    "type": "integer",
                    "example": 5
                },
                "createrId": {
                    "type": "string",
                    "example": "mazhengwang-ust-hk"
                },
                "favorites": {
                    "type": "integer",
                    "example": 1
                },
                "image": {
                    "type": "string"
                },
                "item_id": {
                    "type": "integer",
                    "example": 123455
                },
                "ownerId": {
                    "type": "string",
                    "example": "mazhengwang-ust-hk"
                }
            }
        },
        "v1.JoinEventResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "SUCCESS,USER_JOINED"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "HKUST-NFT",
	Description:      "HKUST-NFT Server API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
