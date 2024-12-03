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
        "/api/e-commerce/product/createProduct": {
            "post": {
                "description": "根据提供的信息创建一个新的产品",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品服务"
                ],
                "parameters": [
                    {
                        "description": "创建产品请求体，包含产品相关信息",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.CreateProductReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "成功创建产品，返回创建后的产品信息",
                        "schema": {
                            "$ref": "#/definitions/endpoint.CreateProductResp"
                        }
                    },
                    "400": {
                        "description": "请求参数错误，例如请求体格式错误等",
                        "schema": {}
                    },
                    "404": {
                        "description": "未找到相关资源或执行创建操作失败",
                        "schema": {}
                    }
                }
            }
        },
        "/api/e-commerce/product/deleteProduct": {
            "post": {
                "description": "根据提供的产品标识删除指定的产品",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品服务"
                ],
                "summary": "删除产品",
                "parameters": [
                    {
                        "description": "删除产品请求体，包含产品标识等相关信息",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.DeleteProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功删除产品，返回成功响应信息",
                        "schema": {
                            "$ref": "#/definitions/endpoint.DeleteProductResp"
                        }
                    },
                    "400": {
                        "description": "请求参数错误，例如请求体格式错误等",
                        "schema": {}
                    },
                    "404": {
                        "description": "未找到要删除的产品或执行删除操作失败",
                        "schema": {}
                    }
                }
            }
        },
        "/api/e-commerce/product/queryProduct": {
            "post": {
                "description": "根据提供的查询条件查询单个产品的信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品服务"
                ],
                "summary": "查询单个产品",
                "parameters": [
                    {
                        "description": "查询产品请求体，包含查询条件等相关信息",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.QueryProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功查询到产品，返回产品信息",
                        "schema": {
                            "$ref": "#/definitions/vo.ProductVO"
                        }
                    },
                    "400": {
                        "description": "请求参数错误，例如请求体格式错误等",
                        "schema": {}
                    },
                    "404": {
                        "description": "未找到符合查询条件的产品",
                        "schema": {}
                    }
                }
            }
        },
        "/api/e-commerce/product/queryProductList": {
            "post": {
                "description": "根据提供的查询条件查询产品的 List信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品服务"
                ],
                "summary": "查询产品列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "此接口无需传入参数，此参数仅为占位示意，实际不会使用。",
                        "name": "noParams",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功查询到产品 List，返回产品列表信息",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vo.ProductVO"
                            }
                        }
                    },
                    "400": {
                        "description": "请求参数错误，例如请求体格式错误等",
                        "schema": {}
                    },
                    "404": {
                        "description": "未找到符合查询条件的产品列表",
                        "schema": {}
                    }
                }
            }
        },
        "/api/e-commerce/product/updateProduct": {
            "post": {
                "description": "根据提供的信息更新指定的产品",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "产品服务"
                ],
                "summary": "更新产品",
                "parameters": [
                    {
                        "description": "更新产品请求体，包含更新后的产品相关信息",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.UpdateProductReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功更新产品，返回更新后的产品信息",
                        "schema": {
                            "$ref": "#/definitions/endpoint.UpdateProductResp"
                        }
                    },
                    "400": {
                        "description": "请求参数错误，例如请求体格式错误等",
                        "schema": {}
                    },
                    "404": {
                        "description": "未找到要更新的产品或执行更新操作失败",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "ebus.Money": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "currency": {
                    "type": "string"
                }
            }
        },
        "endpoint.CreateProductReq": {
            "type": "object",
            "required": [
                "category_id",
                "product_name",
                "skus",
                "status"
            ],
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "product_name": {
                    "type": "string"
                },
                "skus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.SkuEntity"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "endpoint.CreateProductResp": {
            "type": "object"
        },
        "endpoint.DeleteProductReq": {
            "type": "object",
            "required": [
                "spu_id"
            ],
            "properties": {
                "spu_id": {
                    "type": "string"
                }
            }
        },
        "endpoint.DeleteProductResp": {
            "type": "object"
        },
        "endpoint.QueryProductReq": {
            "type": "object",
            "required": [
                "spu_id"
            ],
            "properties": {
                "spu_id": {
                    "type": "string"
                }
            }
        },
        "endpoint.UpdateProductReq": {
            "type": "object",
            "required": [
                "category_id",
                "product_name",
                "skus",
                "spu_id",
                "status"
            ],
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "product_name": {
                    "type": "string"
                },
                "skus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.SkuEntity"
                    }
                },
                "spu_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "endpoint.UpdateProductResp": {
            "type": "object"
        },
        "entity.SkuEntity": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "costAmount": {
                    "$ref": "#/definitions/ebus.Money"
                },
                "deleted": {
                    "type": "boolean"
                },
                "isDefault": {
                    "type": "boolean"
                },
                "sellAmount": {
                    "$ref": "#/definitions/ebus.Money"
                },
                "skuId": {
                    "type": "string"
                },
                "skuName": {
                    "type": "string"
                },
                "spuId": {
                    "type": "string"
                }
            }
        },
        "vo.ProductVO": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "string"
                },
                "deleted": {
                    "type": "boolean"
                },
                "icon": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "skus": {
                    "description": "sku info",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.skuVO"
                    }
                },
                "spuId": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "vo.skuVO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "costAmount": {
                    "$ref": "#/definitions/ebus.Money"
                },
                "deleted": {
                    "type": "boolean"
                },
                "isDefault": {
                    "type": "boolean"
                },
                "sellAmount": {
                    "$ref": "#/definitions/ebus.Money"
                },
                "skuId": {
                    "type": "string"
                },
                "skuName": {
                    "type": "string"
                },
                "spuId": {
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
