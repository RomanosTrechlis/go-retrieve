package cli

import "testing"

func TestProfileCmds(t *testing.T) {
	profile := RegisterProfileCmd(nil)
	if profile.Name() != "profile" {
		t.Errorf("expected 'profile', got '%s'", profile.Name())
	}
	list := RegisterProfileListCmd()
	if list.Name() != "list" {
		t.Errorf("expected 'list', got '%s'", list.Name())
	}
	inspect := RegisterProfileInspectCmd()
	if inspect.Name() != "inspect" {
		t.Errorf("expected 'inspect', got '%s'", inspect.Name())
	}
}
