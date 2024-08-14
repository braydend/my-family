package handlers

import (
	"fmt"
	"net/http"

	"github.com/braydend/my-family/components"
	"github.com/braydend/my-family/models"
	"github.com/braydend/my-family/services"
)

type FamilyTreeHandler struct {
	http.Handler
	familyTree        models.FamilyTree
	familyTreeService *services.FamilyTreeService
}

func NewFamilyTreeHandler() *FamilyTreeHandler {
	return &FamilyTreeHandler{
		familyTree:        models.FamilyTree{},
		familyTreeService: &services.FamilyTreeService{},
	}
}

func (h *FamilyTreeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DEBUG: %s %s\n", r.Method, r.URL.Path)

	if r.Method == http.MethodPost {
		h.Post(w, r)
		return
	}
	h.Get(w, r)
}

func (h *FamilyTreeHandler) Get(w http.ResponseWriter, r *http.Request) {
	components.Page("Family Tree", components.FamilyTreePage(h.familyTree)).Render(r.Context(), w)
}

func (h *FamilyTreeHandler) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	childName := r.FormValue("child")
	parentName := r.FormValue("parent")

	fmt.Printf("DEBUG: Adding child '%s' to parent '%s'\n", childName, parentName)

	newChild := models.NewFamilyMember(childName)
	if len(parentName) > 0 {

		parent, err := h.familyTreeService.FindFamilyMember(string(parentName), &h.familyTree)

		fmt.Printf("DEBUG: Found parent: %s\n", parent.Name)
		fmt.Printf("DEBUG: Parent has %d children\n", len(parent.Children))

		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		parent.AddChild(newChild)

		fmt.Printf("DEBUG: Parent now has %d children\n", len(parent.Children))
	}

	h.familyTree.Members = append(h.familyTree.Members, newChild)

	components.Member(*newChild).Render(r.Context(), w)
}
