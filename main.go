package main

import (
	"github.com/DedeMarantes/gin-api-rest/databases"
	"github.com/DedeMarantes/gin-api-rest/routes"
)

func main() {
	databases.DBConnection()
	routes.HandleRequests()
}
