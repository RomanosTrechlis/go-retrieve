package cli

import "testing"

func TestRegistryCmds(t *testing.T) {
	registry := RegisterRegistryCmd()
	if registry.Name() != "registry" {
		t.Errorf("expected 'registry', got '%s'", registry.Name())
	}
	add := RegisterRegistryAddCmd()
	if add.Name() != "add" {
		t.Errorf("expected 'add', got '%s'", add.Name())
	}
	def := RegisterRegistryDefineCmd()
	if def.Name() != "define" {
		t.Errorf("expected 'define', got '%s'", def.Name())
	}
}
