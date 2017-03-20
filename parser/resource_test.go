package parser

import "testing"

type GoodResource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestListParams(t *testing.T) {
	resource := &GoodResource{}
	params := ListParams(resource)

	if len(params) == 0 {
		t.Fatalf("Exepcted %d params, got %d", 2, len(params))
	}
}
