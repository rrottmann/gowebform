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
	var html_fn string
	var json_fn string
	var server_crt string
	var server_key string

	// flags declaration using flag package
	flag.StringVar(&html_fn, "html", "gowebform.html", "Specify HTML form to serve.")
	flag.StringVar(&json_fn, "json", "gowebform.json", "Specify JSON file to store received POST form data.")
	flag.StringVar(&port, "port", ":8000", "Specify port for HTTPS enabled webserver.")
	flag.StringVar(&server_crt, "cert", "server.crt", "Specify TLS certificate file.")
	flag.StringVar(&server_key, "key", "server.key", "Specify TLS key file.")

	flag.Parse() // after declaring flags we need to call it

	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
            http.ServeFile(w, r, html_fn)
            
		case "POST":
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			// Write r.PostForm to JSON file
			json_data, err := json.MarshalIndent(r.PostForm, "", "   ")
			if err != nil {
				fmt.Fprintf(w, "json.MarshalIndent() err: %v", err)
				return
			}
			ioutil.WriteFile(json_fn, json_data, 0600)
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
	log.Fatal(http.ListenAndServeTLS(port, server_crt, server_key, nil))
}