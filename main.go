package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Hello Team, Thank you for showing interest in Rancher! "

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
