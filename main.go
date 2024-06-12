package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path != "" {
		fmt.Fprintf(w, "Hello, %s!. I'm running on %s architecture", r.URL.Path[1:], runtime.GOARCH)
	} else {
		fmt.Fprintf(w, "Hello World! I'm running on %s architecture", runtime.GOARCH)
	}
}
