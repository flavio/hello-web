package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	msg, msgDefined := os.LookupEnv("HELLO_MSG")
	if !msgDefined {
		msg = "world"
	}

	port, portDefined := os.LookupEnv("HELLO_PORT")
	if !portDefined {
		port = "8000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(msg))
	})

	fmt.Printf("Listening on http://0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
