package models

type Servico struct {
	Id          uint `gorm:"primary_key" json:"id_servico"`
	IdVeiculo   uint `json:"id_veiculo"`
	IdMotorista uint `json:"id_motorista"`
}
