package mock

import (
	"database/sql"
	"esb-assesment/entity/request"
	"esb-assesment/entity/response"

	"github.com/gin-gonic/gin"
)

type MockInvoiceRepository struct {
	BeginTransactionfn        func() (*sql.Tx, error)
	LastIdTransactionfn       func(ctx *gin.Context) (int, error)
	InsertInvoicefn           func(ctx *gin.Context, request *request.InsertInvoiceRequest, tx *sql.Tx) error
	InsertInvoiceItemfn       func(ctx *gin.Context, idInvoice string, request *request.DetailItemRequest, amount int, tx *sql.Tx) error
	EditInvoicefn             func(ctx *gin.Context, request *request.EditInvoiceRequest, tx *sql.Tx) error
	EditInvoiceItemfn         func(ctx *gin.Context, request *request.DetailItemRequest, amount int, idInvoice string, tx *sql.Tx) error
	SelectDetailInvoicefn     func(ctx *gin.Context, idInvoice string) (*response.ResultDataInvoiceResponse, error)
	SelectDetailItemInvoicefn func(ctx *gin.Context, idInvoice string) ([]response.DetailItemResponse, error)
	SelectInvoicefn           func(ctx *gin.Context, request *request.GetAllInvoiceRequest) ([]response.ResultDataInvoiceResponse, error)
	CountInvoiceDatafn        func(ctx *gin.Context, request *request.GetAllInvoiceRequest) (string, error)
}

func (ms *MockInvoiceRepository) BeginTransaction() (*sql.Tx, error) {
	return ms.BeginTransactionfn()
}
func (ms *MockInvoiceRepository) LastIdTransaction(ctx *gin.Context) (int, error) {
	return ms.LastIdTransactionfn(ctx)
}
func (ms *MockInvoiceRepository) InsertInvoice(ctx *gin.Context, request *request.InsertInvoiceRequest, tx *sql.Tx) error {
	return ms.InsertInvoicefn(ctx, request, tx)
}
func (ms *MockInvoiceRepository) InsertInvoiceItem(ctx *gin.Context, idInvoice string, request *request.DetailItemRequest, amount int, tx *sql.Tx) error {
	return ms.InsertInvoiceItemfn(ctx, idInvoice, request, amount, tx)
}
func (ms *MockInvoiceRepository) EditInvoice(ctx *gin.Context, request *request.EditInvoiceRequest, tx *sql.Tx) error {
	return ms.EditInvoicefn(ctx, request, tx)
}
func (ms *MockInvoiceRepository) EditInvoiceItem(ctx *gin.Context, request *request.DetailItemRequest, amount int, idInvoice string, tx *sql.Tx) error {
	return ms.EditInvoiceItemfn(ctx, request, amount, idInvoice, tx)
}
func (ms *MockInvoiceRepository) SelectDetailInvoice(ctx *gin.Context, idInvoice string) (*response.ResultDataInvoiceResponse, error) {
	return ms.SelectDetailInvoicefn(ctx, idInvoice)
}
func (ms *MockInvoiceRepository) SelectDetailItemInvoice(ctx *gin.Context, idInvoice string) ([]response.DetailItemResponse, error) {
	return ms.SelectDetailItemInvoicefn(ctx, idInvoice)
}
func (ms *MockInvoiceRepository) SelectInvoice(ctx *gin.Context, request *request.GetAllInvoiceRequest) ([]response.ResultDataInvoiceResponse, error) {
	return ms.SelectInvoicefn(ctx, request)
}
func (ms *MockInvoiceRepository) CountInvoiceData(ctx *gin.Context, request *request.GetAllInvoiceRequest) (string, error) {
	return ms.CountInvoiceDatafn(ctx, request)
}
