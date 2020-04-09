package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harishankarsivaji/Todo_App_Go/server/router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
