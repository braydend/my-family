package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/braydend/my-family/components"
	"github.com/braydend/my-family/models"
)

func RootHandler() http.Handler {
	stubFamilyTree := models.FamilyTree{[]models.FamilyMember{{"John"}, {"Jack"}, {"Jane"}}}

	component := components.Page("Family Tree", components.FamilyTree(stubFamilyTree))

	return templ.Handler(component)
}
