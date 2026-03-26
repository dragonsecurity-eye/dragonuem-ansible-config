package ansible

import (
	"strings"
	"testing"
)

func TestAddLocalUser(t *testing.T) {
	t.Run("minimal user", func(t *testing.T) {
		u, err := AddLocalUser("Add user", false, "", false, -1, false, false, "", "", "", "testuser",
			false, "", -1, -1, -1, -1, false, "", "", -1, "", "", "", "", false, "", -1, -1, -1, "linux", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.TaskName != "Add user" {
			t.Errorf("expected task name 'Add user', got %s", u.TaskName)
		}
		if u.Parameters.Name != "testuser" {
			t.Errorf("expected name 'testuser', got %s", u.Parameters.Name)
		}
		if u.Parameters.State != "present" {
			t.Errorf("expected state 'present', got %s", u.Parameters.State)
		}
		if u.Parameters.UpdatePassword != "on_create" {
			t.Errorf("expected UpdatePassword='on_create', got %s", u.Parameters.UpdatePassword)
		}
	})

	t.Run("linux password hashing", func(t *testing.T) {
		u, err := AddLocalUser("Add user", false, "", false, -1, false, false, "", "", "", "testuser",
			false, "secret123", -1, -1, -1, -1, false, "", "", -1, "", "", "", "", false, "", -1, -1, -1, "linux", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(u.Parameters.Password, "password_hash") {
			t.Errorf("expected linux password to contain password_hash, got %s", u.Parameters.Password)
		}
		if !strings.Contains(u.Parameters.Password, "sha512") {
			t.Errorf("expected linux password to use sha512, got %s", u.Parameters.Password)
		}
	})

	t.Run("macos password", func(t *testing.T) {
		u, err := AddLocalUser("Add user", false, "", false, -1, false, false, "", "", "", "testuser",
			false, "secret123", -1, -1, -1, -1, false, "", "", -1, "", "", "", "", false, "", -1, -1, -1, "macos", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if strings.Contains(u.Parameters.Password, "password_hash") {
			t.Errorf("macos password should not contain password_hash, got %s", u.Parameters.Password)
		}
	})

	t.Run("skeleton requires create_home", func(t *testing.T) {
		_, err := AddLocalUser("Add user", false, "", false, -1, false, false, "", "", "", "testuser",
			false, "", -1, -1, -1, -1, false, "", "/etc/skel", -1, "", "", "", "", false, "", -1, -1, -1, "linux", false)
		if err == nil {
			t.Fatal("expected error when skeleton set without create_home")
		}
	})

	t.Run("with optional fields", func(t *testing.T) {
		u, err := AddLocalUser("Add user", true, "Test User", true, 1.0, false, true, "staff", "wheel,docker", "/home/test", "testuser",
			false, "", -1, -1, -1, -1, false, "/bin/zsh", "", -1, "", "", "", "", false, "0022", 1000, 2000, 500, "linux", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.Parameters.Comment != "Test User" {
			t.Errorf("expected comment 'Test User', got %s", u.Parameters.Comment)
		}
		if u.Parameters.Group != "staff" {
			t.Errorf("expected group 'staff', got %s", u.Parameters.Group)
		}
		if u.Parameters.Shell != "/bin/zsh" {
			t.Errorf("expected shell '/bin/zsh', got %s", u.Parameters.Shell)
		}
		if u.Parameters.UID != 1000 {
			t.Errorf("expected UID=1000, got %d", u.Parameters.UID)
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := AddLocalUser("", false, "", false, -1, false, false, "", "", "", "testuser",
			false, "", -1, -1, -1, -1, false, "", "", -1, "", "", "", "", false, "", -1, -1, -1, "linux", false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty username", func(t *testing.T) {
		_, err := AddLocalUser("Add user", false, "", false, -1, false, false, "", "", "", "",
			false, "", -1, -1, -1, -1, false, "", "", -1, "", "", "", "", false, "", -1, -1, -1, "linux", false)
		if err == nil {
			t.Fatal("expected error for empty username")
		}
	})
}

func TestRemoveLocalUser(t *testing.T) {
	t.Run("valid remove", func(t *testing.T) {
		u, err := RemoveLocalUser("Remove user", false, "testuser", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.Parameters.State != "absent" {
			t.Errorf("expected state 'absent', got %s", u.Parameters.State)
		}
		if u.Parameters.Name != "testuser" {
			t.Errorf("expected name 'testuser', got %s", u.Parameters.Name)
		}
	})

	t.Run("force remove", func(t *testing.T) {
		u, err := RemoveLocalUser("Remove user", true, "testuser", true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if u.Parameters.Force != true {
			t.Error("expected Force=true")
		}
		if u.IgnoreErrors != true {
			t.Error("expected IgnoreErrors=true")
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := RemoveLocalUser("", false, "testuser", false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty username", func(t *testing.T) {
		_, err := RemoveLocalUser("Remove user", false, "", false)
		if err == nil {
			t.Fatal("expected error for empty username")
		}
	})
}
