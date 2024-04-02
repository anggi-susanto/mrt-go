// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "MRT WW Support",
            "email": "antscpk06@gmail.com"
        },
        "license": {
            "name": "GPL-3.0 License",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/waste-water": {
            "get": {
                "description": "get all waste water data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "waste water"
                ],
                "summary": "get all waste water data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Waste water data",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.WasteWaterData"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "create waste water data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "waste water"
                ],
                "summary": "create waste water data",
                "parameters": [
                    {
                        "description": "waste water data",
                        "name": "waste_water",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.WasteWaterData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.WasteWaterData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    }
                }
            }
        },
        "/waste-water/{id}": {
            "get": {
                "description": "get waste water data by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "waste water"
                ],
                "summary": "get waste water data by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Waste water data ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.WasteWaterData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "update waste water data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "waste water"
                ],
                "summary": "update waste water data",
                "parameters": [
                    {
                        "description": "waste water data",
                        "name": "waste_water",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.WasteWaterData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.WasteWaterData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete waste water data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "waste water"
                ],
                "summary": "delete waste water data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Waste water data ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ColiformsData": {
            "type": "object",
            "properties": {
                "E_coli": {
                    "type": "number"
                },
                "fecal": {
                    "type": "number"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "domain.WasteWaterData": {
            "type": "object",
            "properties": {
                "Ammonium": {
                    "type": "number"
                },
                "BOD": {
                    "type": "number"
                },
                "CDOM": {
                    "type": "number"
                },
                "COD": {
                    "type": "number"
                },
                "Chloride": {
                    "type": "number"
                },
                "Coliforms": {
                    "$ref": "#/definitions/domain.ColiformsData"
                },
                "Crude_Oils": {
                    "type": "number"
                },
                "DOC": {
                    "type": "number"
                },
                "Dissolved_Oxygen": {
                    "type": "number"
                },
                "EC_Salinity_TDS": {
                    "type": "number"
                },
                "Nitrate": {
                    "type": "number"
                },
                "ORP_REDOX": {
                    "type": "number"
                },
                "Optical_Brighteners": {
                    "type": "number"
                },
                "Pressure": {
                    "type": "number"
                },
                "Refined_Oils": {
                    "type": "number"
                },
                "TOC": {
                    "type": "number"
                },
                "Temperature": {
                    "type": "number"
                },
                "Tryptophan": {
                    "type": "number"
                },
                "Turbidity": {
                    "type": "number"
                },
                "_id": {
                    "type": "string"
                },
                "pH": {
                    "type": "number"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "rest.ResponseError": {
            "type": "object",
            "properties": {
                "message": {
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
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "MRT Waste Water API",
	Description: "This is an API Document for MRT Waste Water",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
