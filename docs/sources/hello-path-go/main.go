package main

import (
    "fmt"
	"log"
	"os"
    "net/http"
)

func main() {
	
	// Intanciate a logger 
	hplog := log.New(os.Stdout, "[GO] ", log.LstdFlags)

	// Define a handler for pseudo GET requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		path := r.URL.Path 

		// conditional message based on the provided path
		if path != "/" {
			// return value to the pseudo GET request
			fmt.Fprintf(w, "hello from %s\n", path)
			// log the pseudo GET request
			hplog.Println("GET ", path)
		} else {
			// return value to the pseudo GET request
			fmt.Fprintf(w, "try to add /test as path\n")
			// log the pseudo GET request
			hplog.Println("GET ", path)
		}		
	})

	// Print web service start message at the console
	hplog.Println("Creating a hello path web service with logger")
	hplog.Println("Web service accessible at localhost:8080")

	// Start the service and listen on the given port
    if err := http.ListenAndServe(":8080", nil); err != nil {
		// Log error messages
		hplog.Fatal(err)
	}
}