package models_test

import (
	"testing"

	"github.com/braydend/my-family/models"
)

func TestFamilyMember_AddChild(t *testing.T) {
	alice := models.NewFamilyMember("Alice")
	bob := models.NewFamilyMember("Bob")
	chris := models.NewFamilyMember("Chris")

	alice.AddChild(bob)
	alice.AddChild(chris)

	if len(alice.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(alice.Children))
	}

	if len(bob.Parents) != 1 {
		t.Errorf("expected 1 parent, got %d", len(bob.Parents))
	}

	if len(chris.Parents) != 1 {
		t.Errorf("expected 1 parent, got %d", len(chris.Parents))
	}
}

func TestFamilyMember_AddChild_FailsIfChildIsParent(t *testing.T) {
	alice := models.NewFamilyMember("Alice")
	bob := models.NewFamilyMember("Bob")

	alice.AddParent(bob)

	err := alice.AddChild(bob)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestFamilyMember_AddParent(t *testing.T) {
	alice := models.NewFamilyMember("Alice")
	bob := models.NewFamilyMember("Bob")
	chris := models.NewFamilyMember("Chris")
	darcy := models.NewFamilyMember("Darcy")

	alice.AddParent(bob)
	bob.AddParent(chris)
	alice.AddParent(darcy)

	if len(alice.Parents) != 2 {
		t.Errorf("expected 2 parents, got %d", len(alice.Parents))
	}

	if len(bob.Parents) != 1 {
		t.Errorf("expected 1 parent, got %d", len(bob.Parents))
	}

	if len(bob.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(bob.Children))
	}

	if len(chris.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(chris.Children))
	}

	if len(darcy.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(darcy.Children))
	}
}

func TestFamilyMember_AddParent_FailsIfParentIsChild(t *testing.T) {
	alice := models.NewFamilyMember("Alice")
	bob := models.NewFamilyMember("Bob")

	alice.AddChild(bob)

	err := alice.AddParent(bob)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
