package StudentController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otaviobernardes/api-go-gin/src/external/database"
	Models "github.com/otaviobernardes/api-go-gin/src/models"
)

func New() StudentInterface {
	return &provider{
		database: *database.New(),
	}
}

func (p *provider) GetAll(c *gin.Context) {
	var students []Models.Student
	p.database.Db.Find(&students)

	c.JSON(http.StatusAccepted, students)
}

func (p *provider) GetOne(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	p.database.Db.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (p *provider) Save(c *gin.Context) {

	var student Models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	p.database.Db.Create(&student)
	c.JSON(http.StatusOK, student)
}

func (p *provider) Delete(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	p.database.Db.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted",
	})
}

func (p *provider) Update(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")

	p.database.Db.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	p.database.Db.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}
