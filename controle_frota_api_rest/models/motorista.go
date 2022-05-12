package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type Motorista struct {
	gorm.Model
	Id         uint   `json:"id_motorista" gorm:"primary_key"`
	Nome       string `json:"nome" validate:"required"`
	Carteira   string `json:"carteira" validate:"required"`
	Disponivel bool   `json:"disponivel"`
	Ativo      bool   `json:"ativo"`
}

func ValidacaoMotorista(motorista *Motorista) error {

	validate := validator.New()

	if err := validate.Struct(motorista); err != nil {

		return err
	}
	return nil
}
