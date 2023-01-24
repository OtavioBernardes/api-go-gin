package repoStudent

import (
	"github.com/otaviobernardes/api-go-gin/src/external/database"
	Models "github.com/otaviobernardes/api-go-gin/src/models"
)

func New() IStudentRepository {
	return &provider{
		db: database.New(),
	}
}

func (p *provider) GetAll() ([]Models.Student, error) {
	students := []Models.Student{}
	p.db.Find(&students)
	return students, nil
}

func (p *provider) Save(student *Models.Student) {
	p.db.Save(student)
}

func (p *provider) GetOne(id string) (Models.Student, error) {
	student := Models.Student{}
	p.db.First(&student, id)

	return student, nil
}

func (p *provider) Delete(id string) {
	student := Models.Student{}
	p.db.Delete(&student, id)
}

func (p *provider) Update(student *Models.Student) {
	p.db.Model(&student).UpdateColumns(student)
}
