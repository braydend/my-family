package services

import "github.com/braydend/my-family/models"

type FamilyTreeService struct {
	familyTree models.FamilyTree
}

func (f *FamilyTreeService) GetFamilyMemberCount() uint {
	return f.familyTree.CountFamilyMembers()
}
