package StudentController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otaviobernardes/api-go-gin/src/external/database"
	Models "github.com/otaviobernardes/api-go-gin/src/models"
)

func GetAll(c *gin.Context) {
	var students []Models.Student
	database.Db.Find(&students)

	c.JSON(http.StatusAccepted, students)
}

func GetOne(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	database.Db.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func Save(c *gin.Context) {

	var student Models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.Db.Create(&student)
	c.JSON(http.StatusOK, student)
}

func Delete(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	database.Db.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted",
	})
}

func Update(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	database.Db.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.Db.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}
