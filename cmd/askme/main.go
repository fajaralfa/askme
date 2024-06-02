package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fajaralfa/askme/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	r := router.Create()
	port := os.Getenv("PORT")
	log.Printf("Starting server on port %v\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), r)
	if err != nil {
		log.Fatal(err)
	}

}
