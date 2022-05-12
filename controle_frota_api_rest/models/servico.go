package models

import "gopkg.in/go-playground/validator.v9"

type Servico struct {
	Id          uint   `gorm:"primary_key" json:"id_servico"`
	IdVeiculo   uint   `json:"id_veiculo" validate:"required"`
	IdMotorista uint   `json:"id_motorista" validate:"required"`
	DataInicio  string `json:"data_inicio" validate:"required"`
	DataFim     string `json:"data_fim"`
}

func ValidacaoServico(servico *Servico) error {

	validate := validator.New()

	if err := validate.Struct(servico); err != nil {

		return err
	}
	return nil
}
