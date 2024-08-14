package models

type FamilyMember struct {
	Name     string
	Parents  []FamilyMember
	Children []FamilyMember
}

func NewFamilyMember(name string) *FamilyMember {
	return &FamilyMember{Name: name, Parents: []FamilyMember{}, Children: []FamilyMember{}}
}

func (f *FamilyMember) AddParent(parent FamilyMember) {
	f.Parents = append(f.Parents, parent)
}

func (f *FamilyMember) AddChild(child FamilyMember) {
	f.Children = append(f.Children, child)
}
