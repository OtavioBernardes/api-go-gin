package studentController

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otaviobernardes/api-go-gin/src/models"
	repoStudent "github.com/otaviobernardes/api-go-gin/src/repositores/student"
)

func New(d repoStudent.IStudentRepository) StudentInterface {
	return &provider{
		studentRepository: d,
	}
}

func (p *provider) GetAll(c *gin.Context) {
	students, _ := p.studentRepository.GetAll()

	c.JSON(http.StatusAccepted, students)
}

func (p *provider) Save(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	p.studentRepository.Save(&student)
	c.JSON(http.StatusOK, student)
}

func (p *provider) GetOne(c *gin.Context) {
	id := c.Params.ByName("id")
	student, _ := p.studentRepository.GetOne(id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (p *provider) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id)
	p.studentRepository.Delete(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted",
	})
}

func (p *provider) Update(c *gin.Context) {
	id := c.Params.ByName("id")

	student, _ := p.studentRepository.GetOne(id)
	fmt.Println(student)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	p.studentRepository.Update(&student)
	c.JSON(http.StatusOK, student)
}
