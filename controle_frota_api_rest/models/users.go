package models

import "gopkg.in/go-playground/validator.v9"

type Users struct {
	Id        uint   `json:"id_user" gorm:"primary_key"`
	Username  string `json:"username" validate:"required" gorm:"unique"`
	Email     string `json:"email" validate:"required" gorm:"unique"`
	Password  string `json:"password" validate:"required"`
	Permissao string `json:"permissao" validate:"required"`
}

func ValidacaoUsers(user *Users) error {

	validate := validator.New()

	if err := validate.Struct(user); err != nil {

		return err
	}
	return nil
}
