package main

import (
"fmt"
"net/http"
)

func Hand(w http.ResponseWriter, r *http.Request){
    fmt.Println("Welcome to my Go server")
}

func main() {
    http.HandleFunc("/d/", Hand)
	http.ListenAndServe(":8080", nil)
	
}