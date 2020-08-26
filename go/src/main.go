package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,greeting("Code.education Rocks!"))
}

func main(){

	http.HandleFunc("/",handler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatal(err)
    }
	
}