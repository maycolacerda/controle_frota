package models

import "gopkg.in/go-playground/validator.v9"

// Servico is a struct with fields Id, IdVeiculo, IdMotorista, DataInicio and DataFim.
//
// The first field, Id, is of type uint and has a tag json:"id_servico".
//
// The second field, IdVeiculo, is of type uint and has a tag json:"id_veiculo".
//
// The third field, IdMotorista, is of type uint and has a tag json:"id_motorista".
//
// It validates the struct Servico using the validator package
// The fourth field, DataInicio, is of type string and has a tag
// @property {uint} Id - The primary key of the table
// @property {uint} IdVeiculo - The vehicle id
// @property {uint} IdMotorista - The driver's ID
// @property {string} DataInicio - 2019-01-01
// @property {string} DataFim - date
type Servico struct {
	Id          uint   `gorm:"primary_key" json:"id_servico"`
	IdVeiculo   uint   `json:"id_veiculo" validate:"required"`
	IdMotorista uint   `json:"id_motorista" validate:"required"`
	DataInicio  string `json:"data_inicio" validate:"required"`
	DataFim     string `json:"data_fim"`
}

// It validates the struct Servico using the validator package
func ValidacaoServico(servico *Servico) error {

	validate := validator.New()

	if err := validate.Struct(servico); err != nil {

		return err
	}
	return nil
}
