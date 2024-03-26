package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		//Printing out the error
		fmt.Fprintf(w, "ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w, "Post Request is successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	//Whatever the user have written on the form we will print it here 
	fmt.Fprintf(w, "name:- %s\n", name)
	fmt.Fprintf(w, "address:- %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello")
}

func main()  {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting Server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}