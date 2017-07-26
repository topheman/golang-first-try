package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second) // cancel the request after 1 second
	defer cancel()

	// create a request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	// creates a new Request value with the provided context
	req = req.WithContext(ctx)

	// execute the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
