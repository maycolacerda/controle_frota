package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type Veiculo struct {
	gorm.Model
	Id            uint   `json:"id_veiculo" gorm:"primary_key"`
	Identificador string `json:"identificador" gorm:"unique" validate:"required" regexp:"^[a-zA-Z0-9]{1,10}$"`
	Tipo          string `json:"tipo" validate:"required"`
	Marca         string `json:"marca" validate:"required"`
	Modelo        string `json:"modelo" validate:"required"`
	Rastreador    bool   `json:"rastreador"`
	Placa         string `json:"placa"`
	Disponivel    bool   `json:"disponivel"`
	Km            bool   `json:"km"`
	TipoCarteira  string `json:"tipoCarteira" validate:"required"`
}

func Validacao(veiculo *Veiculo) error {

	validate := validator.New()

	if err := validate.Struct(veiculo); err != nil {

		return err
	}
	return nil
}
