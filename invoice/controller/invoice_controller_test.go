package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"esb-assesment/entity/mock"
	"esb-assesment/entity/request"
	"esb-assesment/entity/response"
	"esb-assesment/invoice/controller"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceController_Add(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.InvoiceController{}

	router := gin.Default()
	router.POST("/add", controller.AddInvoice)

	t.Run("Successful Invoice", func(t *testing.T) {
		bodyReq := request.AddInvoiceRequest{
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
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewReader(requestData))

		controller.InvoiceService = &mock.MockInvoiceService{
			AddInvoicefn: func(ctx *gin.Context, req *request.AddInvoiceRequest, requestId string) (*response.GeneralResponse, error) {
				response := &response.GeneralResponse{
					Code: "200",
					Msg:  "Sukses",
				}
				return response, nil
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response response.GeneralResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Invoice - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Invoice - Service Error", func(t *testing.T) {
		bodyReq := request.AddInvoiceRequest{
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
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewReader(requestData))

		controller.InvoiceService = &mock.MockInvoiceService{
			AddInvoicefn: func(ctx *gin.Context, req *request.AddInvoiceRequest, requestId string) (*response.GeneralResponse, error) {
				return nil, errors.New("service error")
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})
}

func TestInvoiceController_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.InvoiceController{}

	router := gin.Default()
	router.GET("/list", controller.GetAllInvoice)

	t.Run("Successful Invoice", func(t *testing.T) {
		queryParams := url.Values{
			"pageNumber":   {"1"},
			"pageSize":     {"10"},
			"idInvoice":    {"12345"},
			"subject":      {"Sample subject"},
			"issueDate":    {"2023-01-01"},
			"dueDate":      {"2023-01-10"},
			"namaCustomer": {"John Doe"},
			"totalItem":    {"5"},
			"status":       {"Paid"},
		}

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/list?"+queryParams.Encode(), nil)

		controller.InvoiceService = &mock.MockInvoiceService{
			GetAllInvoicefn: func(ctx *gin.Context, req *request.GetAllInvoiceRequest, requestId string) (*response.GeneralResponse, error) {
				response := &response.GeneralResponse{
					Code: "200",
					Msg:  "Success",
				}
				return response, nil
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response response.GeneralResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)
	})

	t.Run("Failed Invoice - Service Error", func(t *testing.T) {
		queryParams := url.Values{
			"pageNumber":   {"1"},
			"pageSize":     {"10"},
			"idInvoice":    {"12345"},
			"subject":      {"Sample subject"},
			"issueDate":    {"2023-01-01"},
			"dueDate":      {"2023-01-10"},
			"namaCustomer": {"John Doe"},
			"totalItem":    {"5"},
			"status":       {"Paid"},
		}

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/list?"+queryParams.Encode(), nil)

		controller.InvoiceService = &mock.MockInvoiceService{
			DetailInvoicefn: func(ctx *gin.Context, req *request.GetInvoiceDetailRequest, requestId string) (*response.GeneralResponse, error) {
				return nil, errors.New("service error")
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestInvoiceController_GetDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.InvoiceController{}

	router := gin.Default()
	router.GET("/detail", controller.GetDetailInvoice)

	t.Run("Successful Invoice", func(t *testing.T) {
		queryParams := url.Values{
			"idInvoice": {"1"},
		}

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/detail?"+queryParams.Encode(), nil)

		controller.InvoiceService = &mock.MockInvoiceService{
			DetailInvoicefn: func(ctx *gin.Context, req *request.GetInvoiceDetailRequest, requestId string) (*response.GeneralResponse, error) {
				response := &response.GeneralResponse{
					Code: "200",
					Msg:  "Success",
				}
				return response, nil
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response response.GeneralResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)
	})

	t.Run("Failed Invoice - Service Error", func(t *testing.T) {
		queryParams := url.Values{
			"idInvoice": {"1"},
		}

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/list?"+queryParams.Encode(), nil)

		controller.InvoiceService = &mock.MockInvoiceService{
			GetAllInvoicefn: func(ctx *gin.Context, req *request.GetAllInvoiceRequest, requestId string) (*response.GeneralResponse, error) {
				return nil, errors.New("service error")
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestInvoiceController_Edit(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.InvoiceController{}

	router := gin.Default()
	router.POST("/edit", controller.EditInvoice)

	t.Run("Successful Invoice", func(t *testing.T) {
		queryParam := url.Values{
			"idInvoice": {"1"},
		}
		bodyReq := request.EditInvoiceRequest{
			IdInvoice:  queryParam.Encode(),
			Subject:    "Pembelian alat berat",
			IssueDate:  "02/09/2007",
			DueDate:    "09/09/2007",
			IdCustomer: "1",
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
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/edit", bytes.NewReader(requestData))

		controller.InvoiceService = &mock.MockInvoiceService{
			EditInvoicefn: func(ctx *gin.Context, req *request.EditInvoiceRequest, requestId string) (*response.GeneralResponse, error) {
				response := &response.GeneralResponse{
					Code: "200",
					Msg:  "Sukses",
				}
				return response, nil
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response response.GeneralResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Invoice - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/edit", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Invoice - Service Error", func(t *testing.T) {
		queryParam := url.Values{
			"idInvoice": {"0000004"},
		}
		bodyReq := request.EditInvoiceRequest{
			IdInvoice:  queryParam.Encode(),
			Subject:    "Pembelian alat berat",
			IssueDate:  "02/09/2007",
			DueDate:    "09/09/2007",
			IdCustomer: "1",
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
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/edit", bytes.NewReader(requestData))

		controller.InvoiceService = &mock.MockInvoiceService{
			EditInvoicefn: func(ctx *gin.Context, req *request.EditInvoiceRequest, requestId string) (*response.GeneralResponse, error) {
				return nil, errors.New("service error")
			},
		}

		router.ServeHTTP(w, ctx.Request)
		assert.Equal(t, http.StatusBadRequest, w.Code)

	})
}
