package models

import (
	"fmt"

	"github.com/google/uuid"
)

type FamilyMember struct {
	id       string
	Name     string
	Parents  []FamilyMember
	Children []FamilyMember
}

func NewFamilyMember(name string) *FamilyMember {
	id := uuid.New().String()
	return &FamilyMember{id: id, Name: name, Parents: []FamilyMember{}, Children: []FamilyMember{}}
}

func (f *FamilyMember) GetId() string {
	return f.id
}

func (f *FamilyMember) AddParent(parent *FamilyMember) error {
	for _, child := range f.Children {
		if child.GetId() == parent.GetId() {
			return fmt.Errorf("cannot add parent that is its own child")
		}
	}

	f.Parents = append(f.Parents, *parent)
	parent.Children = append(parent.Children, *f)

	return nil
}

func (f *FamilyMember) AddChild(child *FamilyMember) error {
	for _, parent := range f.Parents {
		if parent.GetId() == child.GetId() {
			return fmt.Errorf("cannot add child that is its own parent")
		}
	}

	f.Children = append(f.Children, *child)
	child.Parents = append(child.Parents, *f)

	return nil
}
