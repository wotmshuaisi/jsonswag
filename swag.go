package main

type Swag struct {
	Version     string                     `json:"swagger"`
	Info        map[string]string          `json:"info"`
	Paths       map[string]map[string]Path `json:"paths"`
	Definitions map[string]interface{}     `json:"definitions"`
}

type Path struct {
	Summary    string              `json:"summary"`
	Parameters []Parameter         `json:"parameters"`
	Responses  map[string]Response `json:"responses"`
}

type Response struct {
	Description string            `json:"description"`
	Schema      map[string]string `json:"schema"`
}

type Parameter struct {
	Name     string            `json:"name"`
	In       string            `json:"in"`
	Required bool              `json:"required"`
	Schema   map[string]string `json:"schema"`
}

type Definition struct {
	Type       string      `json:"type"`
	Properties interface{} `json:"properties"`
}
