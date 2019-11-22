package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Everyone, Welcome to github Actions. Happy CI/CD with GitHub.")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)

}
