package controllers

import (
	"tendasclub/models"
	"tendasclub/services"
)

func CreateUser(user models.User) (string, error) {
    return services.CreateUser(user)
}


