package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/DedeMarantes/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DBConnection() {
	host := os.Getenv("HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	stringConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, dbUser, dbPass, dbName)
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro ao Conectar o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
