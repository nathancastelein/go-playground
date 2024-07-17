package main

import (
	"fmt"
	"html"
	"log/slog"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", html.EscapeString(r.URL.Query().Get("name")))
}

func main() {
	http.HandleFunc("/hello", HelloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("fail to listen and serve", slog.String("error", err.Error()))
	}
}
