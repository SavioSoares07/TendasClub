package validate

import (
	"regexp"

	"tendasclub/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)



func ValidadeUser(u models.User) error{
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
			is.Email.Error("Email inválido"),
		),
		
		// Validaçõao de senha
		validation.Field(&u.Password,
			validation.Required.Error("Senha é obrigatória"),
			validation.Length(6, 255).Error("Senha deve ter entre no mínimo 6 catacteres"),
		),
		// Validação de número
		validation.Field(&u.Number,
			validation.Required.Error("Informe um numero valido"),
		),
	)
}