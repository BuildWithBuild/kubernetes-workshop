package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Message string
}

func getHello(w http.ResponseWriter, r *http.Request) {
	msg := &Message{
		Message: "hello world!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	for index := 1; index <= 10; index++ {
		num := 10 / index
		log.Printf("index is %d and num is %d \n", index, num)
	}

	w.Write(bytes)
}

func main() {
	http.HandleFunc("/", getHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
