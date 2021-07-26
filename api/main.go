package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	controllers "my_library_app/controllers"
	helpers "my_library_app/helpers"
)

var connection = helpers.ConnectDB()

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/books", controllers.ListBooks(connection)).Methods("GET")
	r.HandleFunc("/api/books/{id}", controllers.GetBook(connection)).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook(connection)).Methods("POST")
	r.HandleFunc("/api/books/{id}", controllers.UpdateBook(connection)).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controllers.DeleteBook(connection)).Methods("DELETE")

	var port = os.Getenv("PORT")

	fmt.Println("Server ready at http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, r))

}
