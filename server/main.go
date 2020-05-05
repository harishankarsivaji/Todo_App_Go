package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/harishankarsivaji/Todo_App_Go/server/api/middleware"
	"github.com/harishankarsivaji/Todo_App_Go/server/api/router"
)

func main() {
	r := router.SetupRouter()
	r.Use(middleware.CORS)

	const PORT = ":8080"

	log.Info("Starting server on port ", PORT)

	log.Fatal(http.ListenAndServe(PORT, r))
}
