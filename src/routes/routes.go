package routes

import (
	"github.com/gin-gonic/gin"
	StudentController "github.com/otaviobernardes/api-go-gin/src/controllers"
	repoStudent "github.com/otaviobernardes/api-go-gin/src/repositores/student"
)

func HandleRequest(server *gin.Engine) {
	studentRepository := repoStudent.New()

	studentController := StudentController.New(
		studentRepository,
	)

	server.GET("/student", studentController.GetAll)
	server.POST("/student", studentController.Save)
	server.GET("/student/:id", studentController.GetOne)
	server.DELETE("/student/:id", studentController.Delete)
	server.PATCH("/student/:id", studentController.Update)
}
