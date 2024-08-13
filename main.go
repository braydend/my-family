package main

import (
	"fmt"
	"net/http"

	"github.com/braydend/my-family/components"

	"github.com/a-h/templ"
)

func main() {
	component := components.Hello("Bray")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
