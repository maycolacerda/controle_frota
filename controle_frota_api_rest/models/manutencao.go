package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type Manutencao struct {
	gorm.Model
	Id         uint   `json:"id_manutencao" gorm:"primary_key"`
	IdVeiculo  uint   `json:"id_veiculo" validate:"required"`
	DataInicio string `json:"data_inicio" validate:"required"`
	DataFim    string `json:"data_fim"`
	Descricao  string `json:"descricao" validate:"required"`
	Concluido  bool   `json:"concluido"`
}

func ValidacaoManutencao(manutencao *Manutencao) error {

	validate := validator.New()

	if err := validate.Struct(manutencao); err != nil {

		return err
	}
	return nil
}
