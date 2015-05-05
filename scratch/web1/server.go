package main

import (
	"fmt"
	"net/http"
	"os"
	"math/rand"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println("Current working directory:", pwd)
	http.HandleFunc("/", handler)
	http.HandleFunc("/r", randomNumberGenerator)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(pwd))))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ranji's first server in Go!\n")
}

func randomNumberGenerator(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", rand.Int())
}

func addIntegers(a int, b int) int {
	return a + b
}
