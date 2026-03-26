package ansible

import "testing"

func TestAddLocalGroup(t *testing.T) {
	t.Run("valid group", func(t *testing.T) {
		g, err := AddLocalGroup("Create admins", "admins", 1001, false, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if g.TaskName != "Create admins" {
			t.Errorf("expected task name 'Create admins', got %s", g.TaskName)
		}
		if g.Parameters.Name != "admins" {
			t.Errorf("expected name 'admins', got %s", g.Parameters.Name)
		}
		if g.Parameters.GID != 1001 {
			t.Errorf("expected GID=1001, got %d", g.Parameters.GID)
		}
		if g.Parameters.State != Present {
			t.Errorf("expected state '%s', got %s", Present, g.Parameters.State)
		}
	})

	t.Run("system group", func(t *testing.T) {
		g, err := AddLocalGroup("Create sys group", "sysgroup", 0, true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if g.Parameters.System != true {
			t.Error("expected System=true")
		}
		if g.Parameters.GID != 0 {
			t.Errorf("expected GID=0 (not set), got %d", g.Parameters.GID)
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := AddLocalGroup("", "admins", 1001, false, false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty group name", func(t *testing.T) {
		_, err := AddLocalGroup("Create group", "", 1001, false, false)
		if err == nil {
			t.Fatal("expected error for empty group name")
		}
	})
}

func TestRemoveLocalGroup(t *testing.T) {
	t.Run("valid remove", func(t *testing.T) {
		g, err := RemoveLocalGroup("Remove admins", "admins", false, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if g.Parameters.State != Absent {
			t.Errorf("expected state '%s', got %s", Absent, g.Parameters.State)
		}
		if g.Parameters.Force != false {
			t.Error("expected Force=false")
		}
	})

	t.Run("force remove", func(t *testing.T) {
		g, err := RemoveLocalGroup("Remove admins", "admins", true, true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if g.Parameters.Force != true {
			t.Error("expected Force=true")
		}
		if g.IgnoreErrors != true {
			t.Error("expected IgnoreErrors=true")
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := RemoveLocalGroup("", "admins", false, false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty group name", func(t *testing.T) {
		_, err := RemoveLocalGroup("Remove group", "", false, false)
		if err == nil {
			t.Fatal("expected error for empty group name")
		}
	})
}
