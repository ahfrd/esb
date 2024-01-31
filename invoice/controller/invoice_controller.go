package controller

import (
	"encoding/json"
	"esb-assesment/entity"
	"esb-assesment/entity/request"
	"esb-assesment/entity/response"
	"esb-assesment/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type InvoiceController struct {
	InvoiceService entity.InvoiceService
}

func NewInvoiceController(invoiceService *entity.InvoiceService) InvoiceController {
	return InvoiceController{InvoiceService: *invoiceService}
}

func (c *InvoiceController) AddInvoice(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.AddInvoiceRequest
	if err := ctx.BindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.InvoiceService.AddInvoice(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
func (c *InvoiceController) GetAllInvoice(ctx *gin.Context) {
	requestId := guuid.New()

	bodyReq := &request.GetAllInvoiceRequest{
		PageNumber:   ctx.Query("pageNumber"),
		PageSize:     ctx.Query("pageSize"),
		IdInvoice:    ctx.Query("idInvoice"),
		Subject:      ctx.Query("subject"),
		IssueDate:    ctx.Query("issueDate"),
		DueDate:      ctx.Query("dueDate"),
		NamaCustomer: ctx.Query("namaCustomer"),
		TotalItem:    ctx.Query("totalItem"),
		Status:       ctx.Query("status"),
	}
	if err := ctx.ShouldBindQuery(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}
	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)
	response, err := c.InvoiceService.GetAllInvoice(ctx, bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
func (c *InvoiceController) GetDetailInvoice(ctx *gin.Context) {
	requestId := guuid.New()

	bodyReq := &request.GetInvoiceDetailRequest{
		IdInvoice: ctx.Query("idInvoice"),
	}
	if err := ctx.ShouldBindQuery(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}
	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)
	response, err := c.InvoiceService.DetailInvoice(ctx, bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
func (c *InvoiceController) EditInvoice(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.EditInvoiceRequest
	bodyReq.IdInvoice = ctx.Query("idInvoice")
	fmt.Println(bodyReq)
	if err := ctx.BindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s ", err.Error()),
		})
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.InvoiceService.EditInvoice(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
