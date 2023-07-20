package main

import (
	"fmt"
	"net/http"
)

func main() {
	//! this is the basic server which will serve the static files in the static folders
	fileServer := http.FileServer(http.Dir("./static"))

	//? this will serve all the files in the static folder
	//? to the root path of the server generally index.html is served at localhost:5005
	http.Handle("/", fileServer)

	//* this will serve the files in the static folder to the /form path
	http.HandleFunc("/form", formHandler)

	//* this will serve the files in the static folder to the /hello path
	http.HandleFunc("/hello", helloHandler)

	//* print the message to the console
	fmt.Println("Starting a server at port 5005")

	//* start the server at port 5005
	if err := http.ListenAndServe(":5005", nil); err != nil {
		panic(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// it will check PrseForm and if there is an error it will print the error
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error: %v", err)
	}
	fmt.Fprintf(w, "POST Request is Successful!\n")

	// take the name and email from the form request
	name := r.FormValue("name")
	email := r.FormValue("email")

	// print the name and email to the web page
	fmt.Fprintf(w, "Name is %s\n", name)
	fmt.Fprintf(w, "Email is %s\n", email)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// check if the path is /hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Path doesn't exists", http.StatusNotFound)
		return
	}

	// check if the method is GET
	if r.Method != "GET" {
		http.Error(w, "Wrong Method Called", http.StatusNotFound)
		return
	}

	// print the message to the web page
	fmt.Fprintf(w, "hello from Golang!")
}
