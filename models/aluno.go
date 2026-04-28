package models

import "gorm.io/gorm"

type Aluno struct {
	//insere infos padrões id, dates
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
