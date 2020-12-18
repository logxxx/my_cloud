package objects

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		log.Println("dataServer.PUT")
		put(w, r)
		return
	}
	if m == http.MethodGet {
		log.Println("dataServer.GET")
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}