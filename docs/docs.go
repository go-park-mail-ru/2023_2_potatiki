// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Dima",
            "url": "http://t.me/belozerovmsk"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/check_auth": {
            "get": {
                "security": [
                    {
                        "AuthKey": []
                    }
                ],
                "description": "Check user is logged in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "CheckAuth",
                "parameters": [
                    {
                        "type": "string",
                        "default": "zuzu-t=xxx",
                        "description": "Token",
                        "name": "Cookie",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile",
                        "schema": {
                            "$ref": "#/definitions/models.Profile"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/auth/logout": {
            "get": {
                "description": "Logout from Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/api/auth/signin": {
            "post": {
                "description": "Login to Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignIn",
                "parameters": [
                    {
                        "description": "SignInPayload",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile",
                        "schema": {
                            "$ref": "#/definitions/models.Profile"
                        }
                    },
                    "400": {
                        "description": "error messege",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/auth/signup": {
            "post": {
                "description": "Create Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignUp",
                "parameters": [
                    {
                        "description": "SignUpPayload",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile",
                        "schema": {
                            "$ref": "#/definitions/models.Profile"
                        }
                    },
                    "400": {
                        "description": "error messege",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/cart/add_product": {
            "post": {
                "description": "add product to cart or change its number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "AddProduct",
                "parameters": [
                    {
                        "description": "product info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CartProductUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "cart info",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/cart/delete_product": {
            "post": {
                "description": "delete product from cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "DeleteProduct",
                "parameters": [
                    {
                        "description": "product info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CartProductDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "cart info",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/cart/summary": {
            "get": {
                "description": "Get cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "GetCart",
                "responses": {
                    "200": {
                        "description": "Cart info",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/cart/update": {
            "post": {
                "description": "Update cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "UpdateCart",
                "parameters": [
                    {
                        "description": "cart info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CartUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "cart info",
                        "schema": {
                            "$ref": "#/definitions/models.Cart"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/category/get_all": {
            "get": {
                "description": "Get category tree",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Category",
                "responses": {
                    "200": {
                        "description": "Category tree",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/order/create": {
            "post": {
                "description": "Create Order using profile ID from cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "CreateOrder",
                "responses": {
                    "200": {
                        "description": "New order info",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "401": {
                        "description": "User unauthorized"
                    },
                    "404": {
                        "description": "Cart not found"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/order/get_all": {
            "get": {
                "description": "Get all Orders using profile ID from cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "GetOrders",
                "responses": {
                    "200": {
                        "description": "All orders info",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Order"
                            }
                        }
                    },
                    "401": {
                        "description": "User unauthorized"
                    },
                    "404": {
                        "description": "Orders not found"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/order/get_current": {
            "get": {
                "description": "Get Current Order using profile ID from cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "GetCurrentOrder",
                "responses": {
                    "200": {
                        "description": "Current order info",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "401": {
                        "description": "User unauthorized"
                    },
                    "404": {
                        "description": "Order not found"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/products/category": {
            "get": {
                "description": "Get products by category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Products",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "category_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Skip number of products",
                        "name": "paging",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Display number of products",
                        "name": "count",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Products by category id",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "error messege",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/products/get_all": {
            "get": {
                "description": "Get products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Products",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Skip number of products",
                        "name": "paging",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Display number of products",
                        "name": "count",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Products array",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "error messege",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/products/{id}": {
            "get": {
                "description": "Get product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product UUID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product info",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "error messege",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/profile/update-data": {
            "post": {
                "description": "Update profile data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "UpdateProfileData",
                "parameters": [
                    {
                        "description": "UpdateProfileDataPayload",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProfileDataPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile",
                        "schema": {
                            "$ref": "#/definitions/models.Profile"
                        }
                    },
                    "400": {
                        "description": "error messege",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/profile/update-photo/{id}": {
            "post": {
                "description": "Update profile photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "UpdatePhoto",
                "parameters": [
                    {
                        "description": "photo",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile",
                        "schema": {
                            "$ref": "#/definitions/models.Profile"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "413": {
                        "description": "Request Entity Too Large"
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        },
        "/api/profile/{id}": {
            "get": {
                "description": "Get profile by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "GetProfile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Profile UUID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile",
                        "schema": {
                            "$ref": "#/definitions/models.Profile"
                        }
                    },
                    "400": {
                        "description": "error messege",
                        "schema": {
                            "$ref": "#/definitions/responser.Response"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Cart": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CartProduct"
                    }
                }
            }
        },
        "models.CartProduct": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "categoryName": {
                    "type": "string"
                },
                "categoryParent": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "img": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "productId": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "models.CartProductDelete": {
            "type": "object",
            "properties": {
                "productId": {
                    "type": "string"
                }
            }
        },
        "models.CartProductUpdate": {
            "type": "object",
            "properties": {
                "productId": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "models.CartUpdate": {
            "type": "object",
            "properties": {
                "productsInfo": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CartProductUpdate"
                    }
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "categoryName": {
                    "type": "string"
                },
                "categoryParent": {
                    "type": "integer"
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderProduct"
                    }
                },
                "statusId": {
                    "type": "integer"
                }
            }
        },
        "models.OrderProduct": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "categoryName": {
                    "type": "string"
                },
                "categoryParent": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "img": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "productId": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "categoryName": {
                    "type": "string"
                },
                "categoryParent": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "img": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "productId": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "models.Profile": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "img": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.SignInPayload": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                }
            }
        },
        "models.SignUpPayload": {
            "type": "object",
            "required": [
                "login",
                "password",
                "phone"
            ],
            "properties": {
                "login": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.UpdateProfileDataPayload": {
            "type": "object",
            "properties": {
                "passwords": {
                    "type": "object",
                    "properties": {
                        "newPass": {
                            "type": "string",
                            "maxLength": 32,
                            "minLength": 8
                        },
                        "oldPass": {
                            "type": "string",
                            "maxLength": 32,
                            "minLength": 8
                        }
                    }
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "responser.Response": {
            "type": "object",
            "properties": {
                "error": {},
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ZuZu Backend API",
	Description:      "API server for ZuZu.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
