package routes

import (
	"esb-assesment/invoice/controller"

	"github.com/gin-gonic/gin"
)

func SetUpInvoiceRoute(router *gin.Engine, InvoiceController *controller.InvoiceController) {
	router.POST("/add", InvoiceController.AddInvoice)
	router.GET("/list", InvoiceController.GetAllInvoice)
	router.GET("/detail", InvoiceController.GetDetailInvoice)
	router.PUT("/edit", InvoiceController.EditInvoice)

}
