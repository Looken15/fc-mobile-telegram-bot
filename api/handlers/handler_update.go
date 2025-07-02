package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func UpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var x interface{}
		err := json.NewDecoder(r.Body).Decode(&x)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Printf("%#v\n", x)

		w.WriteHeader(http.StatusOK)
		return
	}
}
