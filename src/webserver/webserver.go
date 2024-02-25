//Constantes
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("https://www.google.com/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\nwww.google.com", r.URL.Path)
}
