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
            "name": "Euan Meston",
            "email": "euan@openweightlifting.org"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/events": {
            "post": {
                "description": "Fetch a single event by ID and federation.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "POST Requests"
                ],
                "summary": "Fetch a single event",
                "parameters": [
                    {
                        "description": "Federation of the event",
                        "name": "federation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "ID of the event",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/structs.Entry"
                                }
                            }
                        }
                    },
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "options": {
                "description": "Metadata shows the name, federation and date of the event along with the filename in the event_data folder.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OPTIONS Requests"
                ],
                "summary": "Fetch available event metadata within a set date range",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Start date to filter from",
                        "name": "startdate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End date to filter to",
                        "name": "enddate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/structs.EventsList"
                                }
                            }
                        }
                    },
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/history": {
            "post": {
                "description": "Pull a lifter's history by name. The name must be an exact match and can be checked using the search endpoint.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "POST Requests"
                ],
                "summary": "Retrieve a lifter's history",
                "parameters": [
                    {
                        "description": "name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.LifterHistory"
                        }
                    },
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/leaderboard": {
            "post": {
                "description": "This is the used on the index page of the website and pulls the highest single lift for a lifter within the selected filter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "POST Requests"
                ],
                "summary": "Main table on the index page",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Position to begin from within the full query",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Position to stop at within the full query",
                        "name": "stop",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by either total or sinclair",
                        "name": "sortby",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Federation or country to filter by",
                        "name": "federation",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Weightclass to filter by",
                        "name": "weightclass",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Year to filter by",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Not currently used",
                        "name": "startdate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Not currently used",
                        "name": "enddate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.LeaderboardResponse"
                        }
                    }
                }
            }
        },
        "/lifter": {
            "post": {
                "description": "This is used within the lifter page to display a lifter's record. It returns a JSON object that can be used with ChartJS without having to do any additional processing.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "POST Requests"
                ],
                "summary": "Retrieve a lifter's record for use with ChartJS",
                "parameters": [
                    {
                        "description": "name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.ChartData"
                        }
                    },
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/search": {
            "get": {
                "description": "Looks up a lifter by name and returns a list of possible matches. Requires a minimum of 3 characters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GET Requests"
                ],
                "summary": "Search through lifter names",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.NameSearchResults"
                        }
                    }
                }
            }
        },
        "/time": {
            "get": {
                "description": "Returns the current server time.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Utilities and Testing"
                ],
                "summary": "Checking the servers localtime",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.ContainerTime"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.ChartData": {
            "type": "object",
            "properties": {
                "datasets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structs.ChartSubData"
                    }
                },
                "labels": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "structs.ChartSubData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "label": {
                    "type": "string"
                }
            }
        },
        "structs.ContainerTime": {
            "type": "object",
            "properties": {
                "hour": {
                    "type": "integer"
                },
                "min": {
                    "type": "integer"
                },
                "sec": {
                    "type": "integer"
                }
            }
        },
        "structs.Entry": {
            "type": "object",
            "properties": {
                "best_cj": {
                    "type": "number"
                },
                "best_snatch": {
                    "type": "number"
                },
                "bodyweight": {
                    "type": "number"
                },
                "cj_1": {
                    "type": "number"
                },
                "cj_2": {
                    "type": "number"
                },
                "cj_3": {
                    "type": "number"
                },
                "country": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "event": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "instagram": {
                    "type": "string"
                },
                "lifter_name": {
                    "type": "string"
                },
                "sinclair": {
                    "type": "number"
                },
                "snatch_1": {
                    "type": "number"
                },
                "snatch_2": {
                    "type": "number"
                },
                "snatch_3": {
                    "type": "number"
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "structs.EventsList": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structs.SingleEventMetaData"
                    }
                }
            }
        },
        "structs.LeaderboardResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structs.Entry"
                    }
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "structs.LifterHistory": {
            "type": "object",
            "properties": {
                "graph": {
                    "$ref": "#/definitions/structs.ChartData"
                },
                "lifts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structs.Entry"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "structs.NameSearchResults": {
            "type": "object",
            "properties": {
                "names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "structs.SingleEventMetaData": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "federation": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api.openweightlifting.org",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "OpenWeightlifting API",
	Description:      "This is the API for OpenWeightlifting.org",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
