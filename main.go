package main

import (
	"esb-assesment/config"
	"esb-assesment/helpers"
	esbController "esb-assesment/invoice/controller"
	esbRepository "esb-assesment/invoice/repository"
	esbService "esb-assesment/invoice/service"
	"esb-assesment/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	env := helpers.Env{}
	env.StartingCheck()

}
func main() {
	router := gin.Default()

	initDb, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		log.Panic(err)
	}
	EsbRepository := esbRepository.NewInvoiceRepository(initDb)
	//Service
	EsbService := esbService.NewInvoiceService(&EsbRepository)
	//Controller
	EsbController := esbController.NewInvoiceController(&EsbService)
	routes.SetUpInvoiceRoute(router, &EsbController)
	router.Run(":8080")
}
