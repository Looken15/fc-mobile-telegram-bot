package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "Hello, World!\n"); err != nil {
			fmt.Println("Ping Error:", err)
		}
	}
}
