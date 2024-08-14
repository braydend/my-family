package services_test

import (
	"testing"

	"github.com/braydend/my-family/models"
	"github.com/braydend/my-family/services"
)

func TestFamilyTreeService_GetFamilyMemberCount(t *testing.T) {
	familyTree := &models.FamilyTree{
		Members: []*models.FamilyMember{{}, {}},
	}

	service := &services.FamilyTreeService{}

	result := service.GetFamilyMemberCount(familyTree)

	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
}

func TestFamilyTreeService_FindFamilyMember(t *testing.T) {
	bob := models.NewFamilyMember("Bob")

	familyTree := &models.FamilyTree{
		Members: []*models.FamilyMember{
			models.NewFamilyMember("Alice"),
			bob,
		},
	}

	service := &services.FamilyTreeService{}

	result, err := service.FindFamilyMember("Bob", familyTree)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if result != bob {
		t.Errorf("expected '%v', got '%v'", bob, result)
	}
}

func TestFamilyTreeService_FindFamilyMember_NotFound(t *testing.T) {
	familyTree := &models.FamilyTree{
		Members: []*models.FamilyMember{{}},
	}

	service := &services.FamilyTreeService{}

	_, err := service.FindFamilyMember("Bob", familyTree)

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
