package main

import (
	"fmt"
	"net/http"

	"github.com/braydend/my-family/handlers"
)

func main() {
	familyTreeHandler := handlers.NewFamilyTreeHandler()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// http.HandleFunc("/count", countHandler.ServeHTTP)
	http.HandleFunc("/", familyTreeHandler.ServeHTTP)
	// http.Handle("/", handlers.RootHandler())

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
