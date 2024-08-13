package models

type FamilyTree struct {
	Members []*FamilyMember
}

func (f *FamilyTree) CountFamilyMembers() uint {
	return uint(len(f.Members))
}
