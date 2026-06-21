package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selam abi! Ubuntu üzerinden buluta ucan ikinci projem tıkır tıkır calısıyor!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Sunucu 8080 portunda baslatılıyor...")
	http.ListenAndServe(":8080", nil)
}
