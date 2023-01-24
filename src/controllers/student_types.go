package studentController

import (
	"github.com/gin-gonic/gin"
	repoStudent "github.com/otaviobernardes/api-go-gin/src/repositores/student"
)

type provider struct {
	studentRepository repoStudent.IStudentRepository
}
type StudentInterface interface {
	GetAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}
