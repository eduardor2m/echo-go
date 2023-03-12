package models

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Name     string `json:"name"`
		CPF      string `json:"cpf"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}
)
