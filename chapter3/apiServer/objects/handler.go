package objects

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		log.Println("apiServer/objects/Put")
		put(w, r)
		return
	}
	if m == http.MethodGet {
		log.Println("apiServer/objects/Get")
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		log.Println("apiServer/objects/Delete")
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}