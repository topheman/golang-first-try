package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", httpHandler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func httpHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Handler started")
	defer log.Println("Handler ended")
	ctx := req.Context()
	select {
	case <-time.After(time.Second * 5):
		fmt.Fprintln(res, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
