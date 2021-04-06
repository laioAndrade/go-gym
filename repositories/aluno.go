package repositories

import (
	"github.com/laioandrade/go-gym/database"
	"github.com/laioandrade/go-gym/models"
)

func CreateAluno(aluno *models.Aluno) error {
	if err := database.Db.Create(&aluno).Error; err != nil {
		return err
	}
	return nil
}

func GetAlunos() []models.Aluno {
	alunos := []models.Aluno{}
	database.Db.Find(&alunos)
	return alunos
}

func GetAluno(aluno *models.Aluno, id string) error{
	if err := database.Db.Where("id = ?", id).First(&aluno).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAluno(aluno *models.Aluno, input models.Aluno) error{
	if err := database.Db.Model(&aluno).Updates(input).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAluno(aluno *models.Aluno) error {
	if err := database.Db.Delete(&aluno).Error; err != nil {
		return err
	}
	return nil
}