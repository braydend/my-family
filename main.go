package main

import (
	"fmt"
	"net/http"

	"github.com/braydend/my-family/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", handlers.RootHandler())

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
