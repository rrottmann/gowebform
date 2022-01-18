package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// variables read from cmdline
	var port string
	var htmlFn string
	var jsonFn string
	var serverCrt string
	var serverKey string

	// flags declaration using flag package
	flag.StringVar(&htmlFn, "html", "gowebform.html", "Specify HTML form to serve.")
	flag.StringVar(&jsonFn, "json", "gowebform.json", "Specify JSON file to store received POST form data.")
	flag.StringVar(&port, "port", ":8000", "Specify port for HTTPS enabled webserver.")
	flag.StringVar(&serverCrt, "cert", "server.crt", "Specify TLS certificate file.")
	flag.StringVar(&serverKey, "key", "server.key", "Specify TLS key file.")

	flag.Parse() // after declaring flags we need to call it

	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, htmlFn)

		case "POST":
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			// Write r.PostForm to JSON file
			jsonData, err := json.MarshalIndent(r.PostForm, "", "   ")
			if err != nil {
				fmt.Fprintf(w, "json.MarshalIndent() err: %v", err)
				return
			}
			ioutil.WriteFile(jsonFn, jsonData, 0600)
			fmt.Fprintf(w, "{\"code\": 200, \"message\": \"accepted\"}")

		default:
			w.WriteHeader(400)
			fmt.Fprintf(w, "{\"code\": 400, \"message\": \"Sorry, only GET and POST methods are supported.\"")
		}
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	//log.Fatal(http.ListenAndServe(port, nil))
	log.Fatal(http.ListenAndServeTLS(port, serverCrt, serverKey, nil))
}
