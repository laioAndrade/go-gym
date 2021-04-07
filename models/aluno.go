package models

import (
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"nome"`
	UserName string `json:"username" gorm:"unique;not null"`
	Password string `json:"senha" gorm:"not null"`
	Age uint8	`json:"idade" gorm:"not null"`
	Gender string `json:"sexo" gorm:"not null"`
	// CreatedAt    time.Time
  // UpdatedAt    time.Time
}