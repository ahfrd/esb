package service_test

import (
	"database/sql"
	"esb-assesment/entity/mock"
	"esb-assesment/entity/request"
	"esb-assesment/invoice/repository"
	"esb-assesment/invoice/service"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddInvoice(t *testing.T) {

	os.Setenv("DB", "mysql")
	os.Setenv("DBURL", "root:root@tcp(127.0.0.1:3306)/esb")
	db, err := sql.Open(os.Getenv("DB"), os.Getenv("DBURL"))
	if err != nil {
		fmt.Println(err)
	}
	mockRepo := &mock.MockInvoiceRepository{}
	realRepo := repository.NewInvoiceRepository(db)
	service := service.NewInvoiceService(&realRepo)

	t.Run("Success", func(t *testing.T) {
		mockRepo.LastIdTransactionfn = func(ctx *gin.Context) (int, error) {
			return 1, nil
		}
		mockRepo.BeginTransactionfn = func() (*sql.Tx, error) {
			return &sql.Tx{}, nil
		}
		mockRepo.InsertInvoicefn = func(ctx *gin.Context, params *request.InsertInvoiceRequest, tx *sql.Tx) error {
			return nil
		}
		mockRepo.InsertInvoiceItemfn = func(ctx *gin.Context, idInvoice string, item *request.DetailItemRequest, amount int, tx *sql.Tx) error {
			return nil
		}

		params := request.AddInvoiceRequest{
			Subject:   "Pembelian alat berat",
			IssueDate: "02/09/2007",
			DueDate:   "09/09/2007",
			CustomerInfo: request.DetailCustomerRequest{
				IdCustomer: "1",
				Address:    "aaaa",
			},
			DetailItem: []request.DetailItemRequest{
				{
					IdItem:    "1",
					Qty:       2,
					UnitPrice: 50,
				},
				{
					IdItem:    "1",
					Qty:       2,
					UnitPrice: 50,
				},
			},
		}
		ctx := &gin.Context{}

		response, err := service.AddInvoice(ctx, &params, "uid")

		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "200", response.Code)
		assert.Equal(t, "sukses membuat invoice", response.Msg)
	})

}
