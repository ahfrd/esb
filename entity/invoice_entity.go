package entity

import (
	"database/sql"
	"esb-assesment/entity/request"
	"esb-assesment/entity/response"

	"github.com/gin-gonic/gin"
)

type InvoiceService interface {
	AddInvoice(ctx *gin.Context, request *request.AddInvoiceRequest, uid string) (*response.GeneralResponse, error)
	GetAllInvoice(ctx *gin.Context, request *request.GetAllInvoiceRequest, uid string) (*response.GeneralResponse, error)
	DetailInvoice(ctx *gin.Context, request *request.GetInvoiceDetailRequest, uid string) (*response.GeneralResponse, error)
	EditInvoice(ctx *gin.Context, request *request.EditInvoiceRequest, uid string) (*response.GeneralResponse, error)
}

type InvoiceRepository interface {
	BeginTransaction() (*sql.Tx, error)
	LastIdTransaction(ctx *gin.Context) (int, error)
	InsertInvoice(ctx *gin.Context, request *request.InsertInvoiceRequest, tx *sql.Tx) error
	InsertInvoiceItem(ctx *gin.Context, idInvoice string, request *request.DetailItemRequest, amount int, tx *sql.Tx) error
	EditInvoice(ctx *gin.Context, request *request.EditInvoiceRequest, tx *sql.Tx) error
	EditInvoiceItem(ctx *gin.Context, request *request.DetailItemRequest, amount int, idInvoice string, tx *sql.Tx) error
	SelectDetailInvoice(ctx *gin.Context, idInvoice string) (*response.ResultDataInvoiceResponse, error)
	SelectDetailItemInvoice(ctx *gin.Context, idInvoice string) ([]response.DetailItemResponse, error)
	SelectInvoice(ctx *gin.Context, request *request.GetAllInvoiceRequest) ([]response.ResultDataInvoiceResponse, error)
	CountInvoiceData(ctx *gin.Context, request *request.GetAllInvoiceRequest) (string, error)
}
