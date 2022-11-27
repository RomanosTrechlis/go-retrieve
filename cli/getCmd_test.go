package cli

import "testing"

func TestGetCmds(t *testing.T) {
	get := RegisterGetCmd()
	if get.Name() != "get" {
		t.Errorf("expected 'get', got '%s'", get.Name())
	}
	list := RegisterGetListCmd()
	if list.Name() != "list" {
		t.Errorf("expected 'list', got '%s'", list.Name())
	}
}
