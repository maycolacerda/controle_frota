package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

// // Manutencao is a struct that represents a maintenance record.
// // Manutencao
// @property  - Model: This is a struct that contains the following fields: ID, CreatedAt, UpdatedAt,
// DeletedAt.
// @property {uint} Id - The primary key of the table.
// @property {uint} IdVeiculo - The vehicle ID
// @property {string} DataInicio - Start date
// @property {string} DataFim - Date of the end of the maintenance
// @property {string} Descricao - Description of the maintenance
// @property {bool} Concluido - true or false
type Manutencao struct {
	gorm.Model
	Id         uint   `json:"id_manutencao" gorm:"primary_key"`
	IdVeiculo  uint   `json:"id_veiculo" validate:"required"`
	DataInicio string `json:"data_inicio" validate:"required"`
	DataFim    string `json:"data_fim"`
	Descricao  string `json:"descricao" validate:"required"`
	Concluido  bool   `json:"concluido"`
}

// It validates the struct Manutencao using the validator package
func ValidacaoManutencao(manutencao *Manutencao) error {

	validate := validator.New()

	if err := validate.Struct(manutencao); err != nil {

		return err
	}
	return nil
}
