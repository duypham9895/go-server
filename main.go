package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(respWriter http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(respWriter, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method != http.MethodGet {
		http.Error(respWriter, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(respWriter, "Hello World!")
}

func formHandler(respWriter http.ResponseWriter, req *http.Request) {
	errParsedForm := req.ParseForm()
	if errParsedForm != nil {
		fmt.Fprintln(respWriter, "Error parseForm(): ", errParsedForm.Error())
		return
	}

	fmt.Fprintln(respWriter, "POST request successful")

	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintln(respWriter, "Name = ", name)
	fmt.Fprintln(respWriter, "Address = ", address)
}
