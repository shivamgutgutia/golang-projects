package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		http.Error(w, "Method not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"Hello, World!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"Error in parsing the form")
		return
	}
	fmt.Fprintf(w,"POST request successful\n")
	username:= r.FormValue("username")
	password:=r.FormValue("password")
	fmt.Fprintf(w,"Username = %s\n",username)
	fmt.Fprintf(w,"Password = %s",password)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
