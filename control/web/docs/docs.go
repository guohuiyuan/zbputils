// Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/getBotList": {
            "get": {
                "description": "获取机器人qq号",
                "responses": {}
            }
        },
        "/api/getGroupList": {
            "get": {
                "description": "获取群列表",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 123456,
                        "description": "机器人qq号",
                        "name": "self_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/getPluginList": {
            "get": {
                "description": "获取所有插件的状态",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "群号",
                        "name": "group_id",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/api/updatePluginStatus": {
            "post": {
                "description": "更改某一个插件状态",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "群号",
                        "name": "group_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "default": "aireply",
                        "description": "插件名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "default": true,
                        "description": "插件状态",
                        "name": "status",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "zbp api",
	Description:      "zbp restful api document",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
