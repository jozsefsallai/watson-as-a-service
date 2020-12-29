// This file is not being used on the production website, as everything inside
// of the `api` directory is a serverless function that Vercel can handle. The
// mock server is merely for development purposes.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jozsefsallai/watson-as-a-service/api"
)

// PORT specifies the port on which the app will run.
const PORT int = 3000

func rootHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./index.html")
}

func main() {
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/api/encode", api.EncodeHandler)
	http.HandleFunc("/api/decode", api.DecodeHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Printf("Starting WATSON-as-a-Service on port :%d\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", PORT), nil)
	if err != nil {
		log.Fatal(err)
	}
}
