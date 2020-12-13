package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/eamirgh/go-twitter-cleaner/twitter"
	"github.com/joho/godotenv"
)

func keepAlive() {
	for {
		fmt.Println("waiting to hydrate...")
		time.Sleep(15 * time.Minute)
		fmt.Println("hydrating...")
		resp, err := http.Get(os.Getenv("APP_URL"))
		if resp.StatusCode == http.StatusOK {
			fmt.Println("hydrated!")
		}
		if err != nil {
			fmt.Println(err)
		}
	}

}

func hydrate(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("I'm alive!")
	w.Header().Set("Content-Type", "plain/text;charset: utf-8;")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("working!"))
}

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	t := twitter.New()
	go t.DeleteTweets()
	mux := http.NewServeMux()
	mux.HandleFunc("/", hydrate)
	srv := &http.Server{Addr: ":8080", ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second, Handler: mux}
	go keepAlive()
	go log.Fatal(srv.ListenAndServe())
}
