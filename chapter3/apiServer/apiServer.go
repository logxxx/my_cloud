package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./versions"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("RABBITMQ_SERVER", "amqp://admin:admin@49.232.219.233:5672/")
	os.Setenv("LISTEN_ADDRESS", "127.0.0.1:8008")
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Println("apiServer Listening:", os.Getenv("LISTEN_ADDRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}