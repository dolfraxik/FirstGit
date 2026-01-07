package main

import (
"fmt"
"net/http"
)

func Hand(w http.ResponseWriter, r *http.Request){
    fmt.Println("Gosha, i am not  your subscriber anymore!")
}

func main() {
    http.HandleFunc("/d/", Hand)
	http.ListenAndServe(":8080", nil)
	
}