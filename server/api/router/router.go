package router

import (
	"github.com/gin-gonic/gin"
	"github.com/harishankarsivaji/Todo_App_Go/server/api/middleware"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/api/task", middleware.GetAllTask)
	router.POST("/api/task", middleware.CreateTask)
	router.PUT("/api/taskComplete/:id", middleware.TaskComplete)
	router.PUT("/api/undoTask/:id", middleware.UndoTask)
	router.DELETE("/api/deleteTask/:id", middleware.DeleteTask)
	router.DELETE("/api/deleteAllTask", middleware.DeleteAllTask)

	return router
}
