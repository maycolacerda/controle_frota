package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
)

// It returns true if the motorista exists in the database, and false if it doesn't
func ValidaMotorista(idmotorista uint) bool {
	var motorista models.Motorista
	database.DB.First(&motorista, idmotorista)
	if motorista.ID == 0 {
		return false
	} else {
		return true
	}

}

//checa se veiculo existe
// It returns true if the vehicle exists in the database, otherwise it returns false
func ValidaVeiculo(idveiculo uint) bool {

	var veiculo models.Veiculo
	database.DB.First(&veiculo, idveiculo)
	if veiculo.ID == 0 {
		return false
	} else {
		return true
	}

}

// It returns true if the driver's license is equal to the vehicle's license and the vehicle's license
// is not equal to "N/A"
func ValidaAptidao(idmotorista uint, idveiculo uint) bool {

	var motorista models.Motorista
	var veiculo models.Veiculo
	database.DB.First(&motorista, idmotorista)
	database.DB.First(&veiculo, idveiculo)
	if motorista.Carteira == veiculo.TipoCarteira && veiculo.TipoCarteira != "N/A" {
		return true
	} else {
		return false
	}

}

// ValidaServico validates a driver and a vehicle and returns true if both are valid
func ValidaServico(idmotorista uint, idveiculo uint) bool {

	//chnl := make (chan bool)
	vm := ValidaMotorista(idmotorista)
	vv := ValidaVeiculo(idveiculo)
	va := ValidaAptidao(idmotorista, idveiculo)
	if vm && vv && va {
		return true
	} else {
		return false
	}

}
