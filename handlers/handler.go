package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/braydend/my-family/components"
)

func RootHandler() http.Handler {
	component := components.Header("Family tree")

	return templ.Handler(component)
}
