package router

import (
	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	"github.com/harishankarsivaji/Todo_App_Go/server/api/middleware"
)

// SetupRouter is exported and used in main.go
func SetupRouter() *gin.Engine {

	// router := mux.NewRouter()
	router := gin.Default()

	router.GET("/api/task", middleware.GetAllTask)
	router.POST("/api/task", middleware.CreateTask)
	router.PUT("/api/task/{id}", middleware.TaskComplete)
	router.PUT("/api/undoTask/{id}", middleware.UndoTask)
	router.DELETE("/api/deleteTask/{id}", middleware.DeleteTask)
	router.DELETE("/api/deleteAllTask", middleware.DeleteAllTask)

	return router
}
