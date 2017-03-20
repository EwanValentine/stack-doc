package parser

import (
	"reflect"
	"strings"
)

// Handler -
type Handler interface {
	GetResource() interface{}
	GetPath() string
}

// Endpoint -
type Endpoint struct {

	// GET
	Method string

	// /api/v1/resource
	Path string

	// { id: string, name: string }
	Params []Param

	// FindAll
	Handler string
}

// ListFunctions - List the functions found on a struct
func ListFunctions(item interface{}) []Endpoint {
	var methods []Endpoint
	structType := reflect.TypeOf(item)

	// Foreach method on a handler
	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		handler := item.(Handler)
		resource := handler.GetResource()
		path := handler.GetPath()

		httpVerb := "ANY"
		var params []Param

		if strings.HasPrefix(method.Name, "Get") {
			httpVerb = "GET"
		} else if strings.HasPrefix(method.Name, "Create") {
			params = ListParams(resource)
			httpVerb = "POST"
		} else if strings.HasPrefix(method.Name, "Update") {
			params = ListParams(resource)
			httpVerb = "UPDATE"
			path = path + "/:id"
		} else if strings.HasPrefix(method.Name, "Delete") {
			httpVerb = "DELETE"
			path = path + "/:id"
		}

		// If the method is GetResource or GetPath, don't include these
		// As they're utility methods only.
		if method.Name == "GetResource" || method.Name == "GetPath" {
			continue
		}

		endpoint := Endpoint{
			Method:  httpVerb,
			Path:    path,
			Params:  params,
			Handler: method.Name,
		}
		methods = append(methods, endpoint)
	}
	return methods
}
