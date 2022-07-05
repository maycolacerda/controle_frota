package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

// Veiculo is a struct that contains the following fields:
//
// - Model gorm.Model
// - Id uint
// - Identificador string
// - Setor string
// - Tipo string
// - Marca string
// - Modelo string
// - Rastreador bool
// - Placa string
// - Disponivel bool
// - Km bool
// - TipoCarteira string
// @property  - id_veiculo - primary key
// @property {uint} Id - Primary key
// @property {string} Identificador - Unique identifier of the vehicle
// @property {string} Setor - Sector
// @property {string} Tipo - Car, Truck, Motorcycle, etc.
// @property {string} Marca - Brand
// @property {string} Modelo - Model
// @property {bool} Rastreador - true or false
// @property {string} Placa - license plate
// @property {bool} Disponivel - true or false
// @property {bool} Km - true or false
// @property {string} TipoCarteira - Type of license required to drive the vehicle
type Veiculo struct {
	gorm.Model
	Id            uint   `json:"id_veiculo" gorm:"primary_key"`
	Identificador string `json:"identificador" gorm:"unique" validate:"required" regexp:"^[a-zA-Z0-9]{1,10}$"`
	Setor         string `json:"setor" validate:"required"`
	Tipo          string `json:"tipo" validate:"required"`
	Marca         string `json:"marca" validate:"required"`
	Modelo        string `json:"modelo" validate:"required"`
	Rastreador    bool   `json:"rastreador"`
	Placa         string `json:"placa"`
	Disponivel    bool   `json:"disponivel"`
	Km            bool   `json:"km"`
	TipoCarteira  string `json:"tipoCarteira" validate:"required"`
}

// It receives a pointer to a Veiculo struct, creates a new validator, and then validates the struct.
// If there's an error, it returns it. Otherwise, it returns nil
func Validacao(veiculo *Veiculo) error {

	validate := validator.New()

	if err := validate.Struct(veiculo); err != nil {

		return err
	}
	return nil
}
