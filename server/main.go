package main

import (
	"github.com/harishankarsivaji/Todo_App_Go/server/api/middleware"
	"github.com/harishankarsivaji/Todo_App_Go/server/api/router"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	r := router.SetupRouter()
	r.Use(middleware.CORS)

	var PORT = ":9090"

	log.Info("Starting server on port ", PORT)

	log.Fatal(http.ListenAndServe(PORT, r))
}
