package repoStudent

import (
	Models "github.com/otaviobernardes/api-go-gin/src/models"
	"gorm.io/gorm"
)

type provider struct {
	db *gorm.DB
}

type IStudentRepository interface {
	GetAll() ([]Models.Student, error)
	GetOne(id string) (Models.Student, error)
	Save(*Models.Student)
	Delete(id string)
	Update(*Models.Student)
}
