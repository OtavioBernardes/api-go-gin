package routes

import (
	"github.com/gin-gonic/gin"
	StudentController "github.com/otaviobernardes/api-go-gin/src/controllers"
	"github.com/otaviobernardes/api-go-gin/src/external/database"
)

func HandleRequest(server *gin.Engine) {
	studentController := StudentController.New(
		*database.New(),
	)

	server.GET("/student", studentController.GetAll)
	server.GET("/student/:id", studentController.GetOne)
	server.POST("/student", studentController.Save)
	server.DELETE("/student/:id", studentController.Delete)
	server.PATCH("/student/:id", studentController.Update)
}
