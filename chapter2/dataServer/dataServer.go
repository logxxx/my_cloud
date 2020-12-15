package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("RABBITMQ_SERVER", "amqp://admin:admin@49.232.219.233:5672/")
	os.Setenv("LISTEN_ADDRESS", "127.0.0.1:8007")
	os.Setenv("STORAGE_ROOT", ".")
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}