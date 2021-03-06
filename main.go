package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading config file(.env)")
	}

	port := os.Getenv("PORT")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Heroku, my name is Ivan\n")
	})
	log.Printf("Server run on port -> %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
