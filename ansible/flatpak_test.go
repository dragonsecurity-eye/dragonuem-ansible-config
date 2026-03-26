package ansible

import "testing"

func TestInstallFlatpakPackage(t *testing.T) {
	t.Run("valid install", func(t *testing.T) {
		f, err := InstallFlatpakPackage("Install Firefox", "org.mozilla.firefox", false, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.TaskName != "Install Firefox" {
			t.Errorf("expected task name 'Install Firefox', got %s", f.TaskName)
		}
		if f.Parameters.Name != "org.mozilla.firefox" {
			t.Errorf("expected package name 'org.mozilla.firefox', got %s", f.Parameters.Name)
		}
		if f.Parameters.State != "present" {
			t.Errorf("expected state 'present', got %s", f.Parameters.State)
		}
		if f.IgnoreErrors != false {
			t.Error("expected IgnoreErrors=false")
		}
	})

	t.Run("latest state", func(t *testing.T) {
		f, err := InstallFlatpakPackage("Install Firefox", "org.mozilla.firefox", true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.State != "latest" {
			t.Errorf("expected state 'latest', got %s", f.Parameters.State)
		}
	})

	t.Run("ignore errors", func(t *testing.T) {
		f, err := InstallFlatpakPackage("Install Firefox", "org.mozilla.firefox", false, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.IgnoreErrors != true {
			t.Error("expected IgnoreErrors=true")
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := InstallFlatpakPackage("", "org.mozilla.firefox", false, false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty package name", func(t *testing.T) {
		_, err := InstallFlatpakPackage("Install", "", false, false)
		if err == nil {
			t.Fatal("expected error for empty package name")
		}
	})
}

func TestUninstallFlatpakPackage(t *testing.T) {
	t.Run("valid uninstall", func(t *testing.T) {
		f, err := UninstallFlatpakPackage("Remove Firefox", "org.mozilla.firefox", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.State != "absent" {
			t.Errorf("expected state 'absent', got %s", f.Parameters.State)
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := UninstallFlatpakPackage("", "org.mozilla.firefox", false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty package name", func(t *testing.T) {
		_, err := UninstallFlatpakPackage("Remove", "", false)
		if err == nil {
			t.Fatal("expected error for empty package name")
		}
	})
}
