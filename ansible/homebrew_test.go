package ansible

import "testing"

func TestInstallHomeBrewFormula(t *testing.T) {
	t.Run("valid install", func(t *testing.T) {
		f, err := InstallHomeBrewFormula("Install git", "git", "", true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.TaskName != "Install git" {
			t.Errorf("expected task name 'Install git', got %s", f.TaskName)
		}
		if f.Parameters.Name != "git" {
			t.Errorf("expected name 'git', got %s", f.Parameters.Name)
		}
		if f.Parameters.State != "present" {
			t.Errorf("expected state 'present', got %s", f.Parameters.State)
		}
		if f.Become != "yes" {
			t.Errorf("expected Become='yes', got %s", f.Become)
		}
		if f.Parameters.UpdateHomeBrew != true {
			t.Error("expected UpdateHomeBrew=true")
		}
	})

	t.Run("with install options", func(t *testing.T) {
		f, err := InstallHomeBrewFormula("Install git", "git", "--HEAD", false, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.InstallOptions != "--HEAD" {
			t.Errorf("expected install options '--HEAD', got %s", f.Parameters.InstallOptions)
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := InstallHomeBrewFormula("", "git", "", false, false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty formula name", func(t *testing.T) {
		_, err := InstallHomeBrewFormula("Install", "", "", false, false)
		if err == nil {
			t.Fatal("expected error for empty formula name")
		}
	})
}

func TestUpgradeHomeBrewFormula(t *testing.T) {
	t.Run("valid upgrade", func(t *testing.T) {
		f, err := UpgradeHomeBrewFormula("Upgrade git", "git", true, false, "", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.State != "upgraded" {
			t.Errorf("expected state 'upgraded', got %s", f.Parameters.State)
		}
	})

	t.Run("upgrade all without name", func(t *testing.T) {
		f, err := UpgradeHomeBrewFormula("Upgrade all", "", true, true, "", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.UpgradeAll != true {
			t.Error("expected UpgradeAll=true")
		}
	})

	t.Run("empty name without upgrade all", func(t *testing.T) {
		_, err := UpgradeHomeBrewFormula("Upgrade", "", true, false, "", false)
		if err == nil {
			t.Fatal("expected error for empty name without upgradeAll")
		}
	})

	t.Run("with upgrade options", func(t *testing.T) {
		f, err := UpgradeHomeBrewFormula("Upgrade git", "git", false, false, "--fetch-HEAD", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.UpgradeOptions != "--fetch-HEAD" {
			t.Errorf("expected upgrade options '--fetch-HEAD', got %s", f.Parameters.UpgradeOptions)
		}
	})
}

func TestUninstallHomeBrewFormula(t *testing.T) {
	t.Run("valid uninstall", func(t *testing.T) {
		f, err := UninstallHomeBrewFormula("Remove git", "git", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.State != "absent" {
			t.Errorf("expected state 'absent', got %s", f.Parameters.State)
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := UninstallHomeBrewFormula("", "git", false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty formula name", func(t *testing.T) {
		_, err := UninstallHomeBrewFormula("Remove", "", false)
		if err == nil {
			t.Fatal("expected error for empty formula name")
		}
	})
}

func TestInstallHomeBrewCask(t *testing.T) {
	t.Run("valid install", func(t *testing.T) {
		f, err := InstallHomeBrewCask("Install Firefox", "firefox", "", true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.Name != "firefox" {
			t.Errorf("expected name 'firefox', got %s", f.Parameters.Name)
		}
		if f.Parameters.State != "present" {
			t.Errorf("expected state 'present', got %s", f.Parameters.State)
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := InstallHomeBrewCask("", "firefox", "", false, false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty cask name", func(t *testing.T) {
		_, err := InstallHomeBrewCask("Install", "", "", false, false)
		if err == nil {
			t.Fatal("expected error for empty cask name")
		}
	})
}

func TestUpgradeHomeBrewCask(t *testing.T) {
	t.Run("valid upgrade", func(t *testing.T) {
		f, err := UpgradeHomeBrewCask("Upgrade Firefox", "firefox", true, true, false, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.State != "upgraded" {
			t.Errorf("expected state 'upgraded', got %s", f.Parameters.State)
		}
		if f.Parameters.Greedy != true {
			t.Error("expected Greedy=true")
		}
	})

	t.Run("upgrade all without name", func(t *testing.T) {
		f, err := UpgradeHomeBrewCask("Upgrade all", "", false, true, true, false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.UpgradeAll != true {
			t.Error("expected UpgradeAll=true")
		}
	})

	t.Run("empty name without upgrade all", func(t *testing.T) {
		_, err := UpgradeHomeBrewCask("Upgrade", "", false, false, false, false)
		if err == nil {
			t.Fatal("expected error for empty name without upgradeAll")
		}
	})
}

func TestUninstallHomeBrewCask(t *testing.T) {
	t.Run("valid uninstall", func(t *testing.T) {
		f, err := UninstallHomeBrewCask("Remove Firefox", "firefox", false)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f.Parameters.State != "absent" {
			t.Errorf("expected state 'absent', got %s", f.Parameters.State)
		}
	})

	t.Run("empty task name", func(t *testing.T) {
		_, err := UninstallHomeBrewCask("", "firefox", false)
		if err == nil {
			t.Fatal("expected error for empty task name")
		}
	})

	t.Run("empty cask name", func(t *testing.T) {
		_, err := UninstallHomeBrewCask("Remove", "", false)
		if err == nil {
			t.Fatal("expected error for empty cask name")
		}
	})
}
