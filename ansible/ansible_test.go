package ansible

import "testing"

func TestNewAnsiblePlaybook(t *testing.T) {
	pb := NewAnsiblePlaybook()

	if pb.Hosts != "localhost" {
		t.Errorf("expected Hosts=localhost, got %s", pb.Hosts)
	}
	if pb.Connection != "local" {
		t.Errorf("expected Connection=local, got %s", pb.Connection)
	}
	if pb.GatherFacts != false {
		t.Error("expected GatherFacts=false")
	}
	if len(pb.Tasks) != 0 {
		t.Errorf("expected empty Tasks, got %d", len(pb.Tasks))
	}
}

func TestAddAnsibleTask(t *testing.T) {
	pb := NewAnsiblePlaybook()
	pb.AddAnsibleTask("task1")
	pb.AddAnsibleTask("task2")

	if len(pb.Tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(pb.Tasks))
	}
}
