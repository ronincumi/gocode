package main

import (
	"net/http"
)


func main() {

	http.HandleFunc("/", someFunc)
	http.HandleFunc("/id", nameId)
	http.ListenAndServe(":8080", nil)

}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hallo anaku yg soleh"))
}


func nameId(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1> Hello World </h1>"))
}
