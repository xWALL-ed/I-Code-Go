package main

import (
	"fmt"
	"net/http"
	"github.com/xWALL-ed/I-Code-Go/JsonLinkVault/handler"
)

func main() {

	jsonFile := "JsonLinkVault\\url.json"

	fallbackMux := fallbackMux()
	defaultMux := http.HandlerFunc(handler.JsonHandler(jsonFile, fallbackMux))


	fmt.Print("Server Starts at :8080")
	http.ListenAndServe(":8080", defaultMux)

}

func fallbackMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the JsonLinkVault"))
	})
	mux.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Put your URL and what you want to shortly call it inside the JSON file.\nThats it, Now Run the server and use the / endpoint to access your links."))
	})

	return mux
}
