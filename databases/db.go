package databases

import (
	"log"

	"github.com/DedeMarantes/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DBConnection() {
	stringConexao := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao Conectar o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
