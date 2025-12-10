package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)
	log.Println("Serving frontend at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
