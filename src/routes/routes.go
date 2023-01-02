package routes

import (
	"github.com/gin-gonic/gin"
	StudentController "github.com/otaviobernardes/api-go-gin/src/controllers"
)

func HandleRequest(server *gin.Engine) {
	server.GET("/student", StudentController.GetAll)
	server.GET("/student/:id", StudentController.GetOne)
	server.POST("/student", StudentController.Save)
	server.DELETE("/student/:id", StudentController.Delete)
	server.PATCH("/student/:id", StudentController.Update)
}
