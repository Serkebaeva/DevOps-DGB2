package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Packing App into Docker file")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, )
	})
}