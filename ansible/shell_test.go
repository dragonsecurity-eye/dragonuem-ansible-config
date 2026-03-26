package ansible

import "testing"

func TestExecuteScript(t *testing.T) {
	t.Run("valid script", func(t *testing.T) {
		s, err := ExecuteScript("Run setup", "echo hello", "", "", "linux", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.TaskName != "Run setup" {
			t.Errorf("expected task name 'Run setup', got %s", s.TaskName)
		}
		if s.Shell != "echo hello" {
			t.Errorf("expected shell 'echo hello', got %s", s.Shell)
		}
	})

	t.Run("with executable", func(t *testing.T) {
		s, err := ExecuteScript("Run setup", "echo hello", "/bin/bash", "", "linux", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.Args.Executable != "/bin/bash" {
			t.Errorf("expected executable '/bin/bash', got %s", s.Args.Executable)
		}
	})

	t.Run("ignore errors", func(t *testing.T) {
		s, err := ExecuteScript("Run setup", "echo hello", "", "", "linux", true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if s.IgnoreErrors != true {
			t.Error("expected IgnoreErrors=true")
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := ExecuteScript("", "echo hello", "", "", "linux", false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty shell", func(t *testing.T) {
		_, err := ExecuteScript("Run setup", "", "", "", "linux", false)
		if err == nil {
			t.Fatal("expected error for empty shell")
		}
	})
}
