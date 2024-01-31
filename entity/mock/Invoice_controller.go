package mock

import (
	"esb-assesment/entity/request"
	"esb-assesment/entity/response"

	"github.com/gin-gonic/gin"
)

type MockInvoiceService struct {
	AddInvoicefn    func(ctx *gin.Context, request *request.AddInvoiceRequest, uid string) (*response.GeneralResponse, error)
	GetAllInvoicefn func(ctx *gin.Context, request *request.GetAllInvoiceRequest, uid string) (*response.GeneralResponse, error)
	DetailInvoicefn func(ctx *gin.Context, request *request.GetInvoiceDetailRequest, uid string) (*response.GeneralResponse, error)
	EditInvoicefn   func(ctx *gin.Context, request *request.EditInvoiceRequest, uid string) (*response.GeneralResponse, error)
}

func (m *MockInvoiceService) AddInvoice(ctx *gin.Context, request *request.AddInvoiceRequest, uid string) (*response.GeneralResponse, error) {
	if m.AddInvoicefn != nil {
		return m.AddInvoicefn(ctx, request, uid)
	}
	return nil, nil
}

func (m *MockInvoiceService) GetAllInvoice(ctx *gin.Context, request *request.GetAllInvoiceRequest, uid string) (*response.GeneralResponse, error) {
	if m.GetAllInvoicefn != nil {
		return m.GetAllInvoicefn(ctx, request, uid)
	}
	return nil, nil
}

func (m *MockInvoiceService) DetailInvoice(ctx *gin.Context, request *request.GetInvoiceDetailRequest, uid string) (*response.GeneralResponse, error) {
	if m.DetailInvoicefn != nil {
		return m.DetailInvoicefn(ctx, request, uid)
	}
	return nil, nil
}
func (m *MockInvoiceService) EditInvoice(ctx *gin.Context, request *request.EditInvoiceRequest, uid string) (*response.GeneralResponse, error) {
	if m.EditInvoicefn != nil {
		return m.EditInvoicefn(ctx, request, uid)
	}
	return nil, nil
}
