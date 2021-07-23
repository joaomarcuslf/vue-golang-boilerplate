package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println("Server listening on port 5000")
	log.Panic(
		http.ListenAndServe(":5000", nil),
	)
}
