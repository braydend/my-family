package services

import (
	"fmt"

	"github.com/braydend/my-family/models"
)

type FamilyTreeService struct {
}

func (f *FamilyTreeService) GetFamilyMemberCount(familyTree *models.FamilyTree) uint {
	return familyTree.CountFamilyMembers()
}

func (f *FamilyTreeService) FindFamilyMember(name string, familyTree *models.FamilyTree) (*models.FamilyMember, error) {
	for _, member := range familyTree.Members {
		if member.Name == name {
			return member, nil
		}
	}

	return &models.FamilyMember{}, fmt.Errorf("family member '%s' not found", name)
}
