package main

import sd "github.com/ewanvalentine/stack-doc"

func main() {

	// Create instance of a resource or model for instance
	resource := &Resource{}

	// Initialise StackDoc
	stackdoc := sd.Init()

	// Create an instance of your handler
	handler := NewResourceHandler("/api/v1/resource", resource)

	// Add your handler
	stackdoc.AddHandler(handler)

	// Serves API/UI for StackDoc on ports, 9090/7070
	stackdoc.Serve()
}

// Resource - Test resource
type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

// ResourceHandler -
type ResourceHandler struct {
	BasePath string
	Resource sd.Resourceable
}

// NewResourceHandler -
func NewResourceHandler(path string, resource sd.Resourceable) *ResourceHandler {
	return &ResourceHandler{path, resource}
}

// GetResource - This is required for StackDoc
func (handler *ResourceHandler) GetResource() interface{} {
	return handler.Resource
}

// GetPath - This is required for StackDoc
func (handler *ResourceHandler) GetPath() string {
	return handler.BasePath
}

func (handler *ResourceHandler) Get()                          {}
func (handler *ResourceHandler) GetAll()                       {}
func (handler *ResourceHandler) NoneVerbedMethod()             {}
func (handler *ResourceHandler) UpdateResource()               {}
func (handler *ResourceHandler) DeleteResource()               {}
func (handler *ResourceHandler) UpdateResourceWhereSomething() {}
