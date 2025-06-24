package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Number   string `json:"number"`
}


func (u User) Validade() error{
	return validation.ValidateStruct(&u,	
		// Validação de nome
		validation.Field(&u.Name,
			 validation.Required.Error("Nome é obrigatório"),
			 validation.Match(regexp.MustCompile(`^[a-zA-Z\s]+$`)).Error("Nome inválido, deve conter apenas letras"),
			 validation.Length(3, 255).Error("Nome deve ter entre 3 e 255 caracteres"),
			),
		
		// Validação de email
		validation.Field(&u.Email,
		 	validation.Required.Error("Email é obrigatório"),
		),
		
		// Validaçõao de senha
		validation.Field(&u.Password,
			validation.Required.Error("Senha é obrigatória"),
			validation.Length(6, 255).Error("Senha deve ter entre no mínimo 6 catacteres"),
		),
		// Validação de número
		validation.Field(&u.Number,
			validation.Required.Error("Informe um numero valido"),
			validation.Match(regexp.MustCompile(`^\+?[0-9]$`)).Error("Número inválido"),
		),
	)
}