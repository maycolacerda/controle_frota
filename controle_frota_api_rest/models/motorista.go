package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

// Motorista is a struct with fields Id, Nome, Carteira, Disponivel and Ativo.
// @property  - Model: This is a special property that will be used by GORM to map the struct to a
// database table.
// @property {uint} Id - The primary key of the table.
// @property {string} Nome - Name of the driver
// @property {string} Carteira - Driver's license
// @property {bool} Disponivel - true if the driver is available to work
// @property {bool} Ativo - If the driver is active or not
type Motorista struct {
	gorm.Model
	Id         uint   `json:"id_motorista" gorm:"primary_key"`
	Nome       string `json:"nome" validate:"required"`
	Carteira   string `json:"carteira" validate:"required"`
	Disponivel bool   `json:"disponivel"`
	Ativo      bool   `json:"ativo"`
}

// It validates the struct Motorista using the validator package
func ValidacaoMotorista(motorista *Motorista) error {

	validate := validator.New()

	if err := validate.Struct(motorista); err != nil {

		return err
	}
	return nil
}
