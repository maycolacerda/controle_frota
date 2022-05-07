package database

import (
	"controle_frota_gin/models"
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	db, erro := sql.Open("mysql", "root:root@tcp(localhost:5432)/controle-frota?parseTime=true")
	if erro != nil {
		log.Fatal(erro.Error())
	}
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("erro na conexão com o banco de dados: " + err.Error())
	} else {
		log.Println("conexão com o banco de dados estabelecida com sucesso")
	}
	Automigrate()
}
func Automigrate() {

	DB.AutoMigrate(&models.Veiculo{})
}
