package parser

import "reflect"

// Param - Parameter object
type Param struct {
	Type string
	Name string
	Tag  string
}

// ListParams - List the parameters on a struct
func ListParams(i interface{}) []Param {
	var params []Param

	// Get struct type
	structType := reflect.TypeOf(i).Elem()

	// For each field in the struct
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		newParam := Param{
			Type: field.Type.Name(),
			Name: field.Name,
			Tag:  field.Tag.Get("json"),
		}
		params = append(params, newParam)
	}
	return params
}
