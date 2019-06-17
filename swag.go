package main

const (
	// Number ...
	Number = "number"
	// String ...
	String = "string"
	// Boolean ...
	Boolean = "boolean"
	// Array ...
	Array = "array"
	// Object ...
	Object = "object"
)

// Swag ...
type Swag struct {
	Version     string                     `json:"swagger"`
	Info        map[string]string          `json:"info"`
	Paths       map[string]map[string]Path `json:"paths"`
	Definitions map[string]interface{}     `json:"definitions"`
}

// Path ...
type Path struct {
	Summary    string              `json:"summary"`
	Parameters []*Parameter        `json:"parameters"`
	Responses  map[string]Response `json:"responses"`
}

// Response ...
type Response struct {
	Description string            `json:"description"`
	Schema      map[string]string `json:"schema"`
}

// Parameter ...
type Parameter struct {
	Name     string            `json:"name"`
	In       string            `json:"in"`
	Required bool              `json:"required"`
	Schema   map[string]string `json:"schema"`
}

// Definition ...
type Definition struct {
	Type       string                 `json:"type"`
	Items      *Definition            `json:"items,omitempty"`
	Properties map[string]*Definition `json:"properties,omitempty"`
}
