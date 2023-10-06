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
        "/user-basic/add": {
            "post": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Show an account",
                "parameters": [
                    {
                        "description": "用户信息参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserBasic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "添加成功",
                        "schema": {
                            "$ref": "#/definitions/model.UserBasic"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.UserBasic": {
            "type": "object",
            "properties": {
                "clientIp": {
                    "description": "客户端ip",
                    "type": "string"
                },
                "clientPort": {
                    "description": "客户端接口",
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deviceInfo": {
                    "description": "设备信息",
                    "type": "string"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string"
                },
                "heartbeatTime": {
                    "description": "心跳时间",
                    "type": "string"
                },
                "identity": {
                    "description": "唯一标识",
                    "type": "string"
                },
                "isLogOut": {
                    "description": "登录状态",
                    "type": "integer"
                },
                "loginOutTime": {
                    "description": "心跳时间",
                    "type": "string"
                },
                "loginTime": {
                    "description": "登录时间",
                    "type": "string"
                },
                "name": {
                    "description": "用户名字",
                    "type": "string"
                },
                "password": {
                    "description": "用户密码",
                    "type": "string"
                },
                "phone": {
                    "description": "用户电话",
                    "type": "string"
                },
                "updatedAt": {
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
