package stackdoc

import "testing"

type Good struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

type GoodHandler struct {
	BasePath string
	Resource Resourceable
}

func NewGoodHandler(path string, resource Resourceable) *GoodHandler {
	return &GoodHandler{path, resource}
}

func (handler *GoodHandler) GetResource() interface{} {
	return handler.Resource
}

func (handler *GoodHandler) GetPath() string {
	return handler.BasePath
}

func (handler *GoodHandler) GetSomething()   {}
func (handler *GoodHandler) PostOtherThing() {}

func setUp() *StackDoc {
	return Init()
}

func TestAddHandler(t *testing.T) {
	sd := setUp()
	good := &Good{}
	handler := NewGoodHandler("/api/v1/good", good)
	sd.AddHandler(handler)

	if len(sd.List()) == 0 {
		t.Fatalf("Failed to add handler, expected %d, got %d", 2, len(sd.List()))
	}

	if sd.List()[0].Path != "/api/v1/good" {
		t.Fatalf("Paths mismatched, expected %s got %s", "/api/v1/good", sd.List()[0])
	}
}
