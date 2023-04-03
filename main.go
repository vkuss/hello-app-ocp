package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		nodeName := os.Getenv("MY_NODE_NAME")
		fmt.Fprintf(w, "Hello, World! This is running on node %s\n", nodeName)
	})
	http.ListenAndServe(":8080", nil)
}
