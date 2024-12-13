package main

import (
	"fmt"
	"net/http"
)

var s []string = []string{"a", "b", "c"}

func main() {
	http.HandleFunc("/search", Search)
	http.ListenAndServe(":8086", nil)
}

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	userName := r.URL.Query().Get("user_name")
	for _, name := range s {
		if name == userName {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello " + userName))
			return
		}
	}

	http.Error(w, "user not found in db", http.StatusNotFound)
	return
}
