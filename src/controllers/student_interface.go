package StudentController

import (
	"github.com/gin-gonic/gin"
	"github.com/otaviobernardes/api-go-gin/src/external/database"
)

type provider struct {
	database database.Provider
}

type StudentInterface interface {
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Save(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}
