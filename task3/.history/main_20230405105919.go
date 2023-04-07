package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Packing App into Docker file")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Dockerized Hello World App")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}