package main

func typeDetection(v interface{}) string {
	if _, ok := v.(uint); ok {
		return "number"
	}
	if _, ok := v.(int); ok {
		return "number"
	}
	if _, ok := v.(float64); ok {
		return "number"
	}
	if _, ok := v.(string); ok {
		return "string"
	}
	if _, ok := v.(bool); ok {
		return "boolean"
	}
	if _, ok := v.(map[string]interface{}); ok {
		return "object"
	}
	if _, ok := v.([]interface{}); ok {
		return "array"
	}
	panic(v)
}
