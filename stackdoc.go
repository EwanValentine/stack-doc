package stackdoc

import (
	"net/http"

	api "github.com/ewanvalentine/stack-api"
	"github.com/ewanvalentine/stack-doc/parser"
	registrar "github.com/ewanvalentine/stack-registrar"
)

// API -
type API struct {
	Endpoint []parser.Endpoint
}

// Resourceable - API resource
type Resourceable interface{}

// StackDoc -
type StackDoc struct {
	Handlers []parser.Handler
}

// Init -
func Init() *StackDoc {
	return &StackDoc{}
}

// AddHandler -
func (stackdoc *StackDoc) AddHandler(handler parser.Handler) {
	stackdoc.Handlers = append(stackdoc.Handlers, handler)
}

// List - List all endpoints
func (stackdoc *StackDoc) List() []parser.Endpoint {
	var allEndpoints []parser.Endpoint
	for _, handler := range stackdoc.Handlers {
		endpoints := parser.ListFunctions(handler)
		for _, endpoint := range endpoints {
			allEndpoints = append(allEndpoints, endpoint)
		}
	}
	return allEndpoints
}

// Serve - Serves as API
func (stackdoc *StackDoc) Serve() {
	registry := registrar.Init(registrar.SetHost("http://localhost:8080"))
	app := api.Init(registry)

	app.Get("/api/docs", func(c *api.Context) {
		c.JSON(stackdoc.List(), http.StatusOK)
	})

	// Static assets, i.e react build folder
	http.Handle("/", http.FileServer(http.Dir("./ui/build")))

	// Run API
	go app.Run(api.SetPort(":9090"))

	// Run UI
	http.ListenAndServe(":7070", nil)
}
