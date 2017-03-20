package parser

import "testing"

type Thing struct{}

func (thing Thing) GetPath() string { return "test" }
func (thing Thing) GetResource() interface{} {
	return &Thing{}
}

func (thing Thing) GetTest()             {}
func (thing Thing) GetOtherTest()        {}
func (thing Thing) CreateThing()         {}
func (thing Thing) UpdateSomething()     {}
func (thing Thing) DeleteSomethingElse() {}

func setUp() {

}

func TestListFunctions(t *testing.T) {
	thing := &Thing{}
	endpoints := ListFunctions(thing)
	if len(endpoints) != 5 {
		t.Fatalf("Wrong amount of functions, expected %d got %d", 5, len(endpoints))
	}

	if endpoints[0].Method != "POST" {
		t.Fatalf("Expected method %s got %s", "POST", endpoints[0].Method)
	}

	if endpoints[1].Method != "DELETE" {
		t.Fatalf("Expected method %s got %s", "DELETE", endpoints[1].Method)
	}

	if endpoints[2].Method != "GET" {
		t.Fatalf("Expected method %s got %s", "GET", endpoints[2].Method)
	}

	if endpoints[3].Method != "GET" {
		t.Fatalf("Expected method %s got %s", "GET", endpoints[3].Method)
	}

	if endpoints[4].Method != "UPDATE" {
		t.Fatalf("Expected method %s got %s", "UPDATE", endpoints[4].Method)
	}
}
