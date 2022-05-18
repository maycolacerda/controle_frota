package controllers

import (
	"controle_frota_gin/database"
	"controle_frota_gin/models"
)

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
func ValidaVeiculo(idveiculo uint) bool {

	var veiculo models.Veiculo
	database.DB.First(&veiculo, idveiculo)
	if veiculo.ID == 0 {
		return false
	} else {
		return true
	}

}

//checa se motorista possui carteira para dividir veiculo
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
