# StackDoc

StackDoc is an automated API documentation framework which integrates directly into your API's code to produce automated, easy to read documentation for your API. StackDoc uses reflection and annotation tags to understand how your API works.

## Example
```go
package main

import sd "github.com/ewanvalentine/stack-doc"

func main() {

	// Create instance of a resource or model for instance
	resource := &Resource{}

	// Initialise StackDoc
	stackdoc := sd.Init()

	// Create an instance of your handler
	handler := NewResourceHandler("/api/v1/resource", resource, stackdoc)

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
	BasePath  string
	Resource  sd.Resourceable
	StackDock *sd.StackDoc
}

// NewResourceHandler -
func NewResourceHandler(path string, resource sd.Resourceable, stackdoc *sd.StackDoc) *ResourceHandler {
	return &ResourceHandler{path, resource, stackdoc}
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

```

You can quickly run this by doing `$ go run example/main.go`. The route is relative, so ensure to run it from this directory. Then heading to `http://localhost:7070` in your browser.

Above you can see we're creating a handler for an api called 'ResourceHandler', this takes a base path for that resource, i.e `/api/v1/resources`, then it takes an instance of the resource itself. This is important as it allows StackDoc to understand the relationship between your model, and your endpoints.

You will notice that all the methods for our api endpoints begin with HTTP verbs, this allows StackDoc to infer the correct HTTP method for each endpoint. So be sure to prefix your methods with a HTTP verb, for example `PutSomething` or `GetAllTheThings`, this will translate to `PUT` and `GET`.

Currently there's no way to correctly assume handler methods HTTP verb without prefixing the verb to the method name, luckily, this is pretty good practice anyway.

## How to run
Register StackDoc in your code as shown above, then ensure you're running `stackdoc.Serve()`, you may need to run it as a go-routine `go stackdoc.Server()` if you have another webserver for your API called after this point.

### Next big steps
- Be able to customise routes without prefixing the verb to the method name.
- Consolidate `Serve()` to run on one port instead of two.
- Include filter params per endpoint.
- Allow users to define specific resources, or modified resources per endpoint. I.e you may only pass certain fields of a model into a certain endpoint.
- Make the doc layouts and css customisable.
- Customisable logo.
- Split docs by resource or group of endpoints.
