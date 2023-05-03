package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	skeletoncss "github.com/monasticacadedemy/go-skeleton-css"
)

//go:embed index.html
var html []byte

func main() {
	port := ":8000"

	http.Handle("/css/", http.StripPrefix("/css/", skeletoncss.FileServer))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "index.html", time.Time{}, bytes.NewReader(html))
	})

	fmt.Println("listening on", port)
	http.ListenAndServe(port, nil)
}
