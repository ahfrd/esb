package service

import (
	"errors"
	"esb-assesment/entity"
	"esb-assesment/entity/request"
	"esb-assesment/entity/response"
	"esb-assesment/helpers"
	"esb-assesment/helpers/constant"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type invoiceServiceImpl struct {
	InvoiceRepository entity.InvoiceRepository
}

func NewInvoiceService(invoiceRepository *entity.InvoiceRepository) entity.InvoiceService {
	return &invoiceServiceImpl{
		InvoiceRepository: *invoiceRepository,
	}
}

func (s *invoiceServiceImpl) AddInvoice(ctx *gin.Context, params *request.AddInvoiceRequest, uid string) (*response.GeneralResponse, error) {
	lastIdTrans, err := s.InvoiceRepository.LastIdTransaction(ctx)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s", err.Error()),
		}, nil
	}
	if lastIdTrans == 0 {
		lastIdTrans = 0000000
	}
	lastIdTrans++
	issueDate, _ := time.Parse(constant.DateYYYYMMDDSLASH, params.IssueDate)
	dueDate, _ := time.Parse(constant.DateYYYYMMDDSLASH, params.DueDate)
	fmt.Println(dueDate.Format(constant.DateYYYYMMDDDASH))

	paramsInsertInvoice := &request.InsertInvoiceRequest{
		IdInvoice:  fmt.Sprintf("%07d", lastIdTrans),
		Subject:    params.Subject,
		IssueDate:  issueDate.Format(constant.DateYYYYMMDDDASH),
		DueDate:    dueDate.Format(constant.DateYYYYMMDDDASH),
		IdCustomer: params.CustomerInfo.IdCustomer,
		TotalItem:  len(params.DetailItem),
		Status:     "Unpaid",
	}
	txBegin, err := s.InvoiceRepository.BeginTransaction()
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s", err.Error()),
		}, nil
	}
	insertInvoice := s.InvoiceRepository.InsertInvoice(ctx, paramsInsertInvoice, txBegin)
	if insertInvoice != nil {
		helpers.LogError(ctx, insertInvoice.Error(), uid)
		_ = txBegin.Rollback()
		return &response.GeneralResponse{
			Code: "401",
			Msg:  fmt.Sprintf("Terjadi kesalahan insert invoice %s", insertInvoice.Error()),
		}, nil

	}
	var subTotal int
	for _, item := range params.DetailItem {
		amount := item.Qty * item.UnitPrice
		subTotal += amount
		insertItemInvoice := s.InvoiceRepository.InsertInvoiceItem(ctx, fmt.Sprintf("%07d", lastIdTrans), &item, amount, txBegin)
		if insertItemInvoice != nil {
			helpers.LogError(ctx, insertItemInvoice.Error(), uid)
			_ = txBegin.Rollback()
			return &response.GeneralResponse{
				Code: "401",
				Msg:  fmt.Sprintf("Terjadi kesalahan insert item invoice %s", insertItemInvoice.Error()),
			}, nil
		}
	}
	_ = txBegin.Commit()
	presentase := float64(subTotal) * (10.0 / 100)
	grandTotal := float64(subTotal) + presentase

	var data response.SummaryInvoiceResponse
	data.GrandTotal = grandTotal
	data.Tax = presentase
	data.TotalItem = len(params.DetailItem)
	data.SubTotal = float64(subTotal)

	var resData response.GeneralResponse
	resData.Code = "200"
	resData.Msg = "sukses membuat invoice"
	resData.Data = data

	return &resData, nil
}

func (s *invoiceServiceImpl) GetAllInvoice(ctx *gin.Context, request *request.GetAllInvoiceRequest, uid string) (*response.GeneralResponse, error) {
	var listDataInvoice response.ListInvoiceResponse
	var resData response.GeneralResponse
	if request.IssueDate != "" {
		issueDate, _ := time.Parse(constant.DateYYYYMMDDSLASH, request.IssueDate)
		request.IssueDate = issueDate.Format(constant.DateYYYYMMDDDASH)
	}
	if request.DueDate != "" {
		dueDate, _ := time.Parse(constant.DateYYYYMMDDSLASH, request.DueDate)
		request.DueDate = dueDate.Format(constant.DateYYYYMMDDDASH)
	}

	countDataInvoice, err := s.InvoiceRepository.CountInvoiceData(ctx, request)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan count data invoice %s", err.Error()),
		}, nil
	}
	if countDataInvoice == "0" {
		errMsg := errors.New("data invoice tidak ditemukan")
		helpers.LogError(ctx, errMsg.Error(), uid)

		return &response.GeneralResponse{
			Code: "404",
			Msg:  errMsg.Error(),
		}, nil
	}

	floTotalRow, _ := strconv.ParseFloat(countDataInvoice, 64)
	floRecordPerPage, _ := strconv.ParseFloat(request.PageSize, 64)
	floCurrentPage, _ := strconv.ParseFloat(request.PageNumber, 64)
	totalPage := math.Ceil(floTotalRow / floRecordPerPage)
	currentPage := request.PageNumber
	firstRecord := (floCurrentPage - 1) * floRecordPerPage
	startRecord := firstRecord + 1
	countData := response.CountData{
		TotalRecord:   countDataInvoice,
		TotalPage:     strconv.FormatFloat(totalPage, 'f', 0, 64),
		RecordPerPage: request.PageSize,
		CurrentPage:   currentPage,
		StartRecord:   strconv.FormatFloat(startRecord, 'f', 0, 64),
		FirstRecord:   strconv.FormatFloat(firstRecord, 'f', 0, 64),
	}
	request.PageNumber = strconv.FormatFloat(firstRecord, 'f', 0, 64)
	selectDataInvoice, err := s.InvoiceRepository.SelectInvoice(ctx, request)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan  %s", err.Error()),
		}, nil
	}
	listDataInvoice.ListData = selectDataInvoice
	listDataInvoice.Pagination = countData

	resData.Code = "200"
	resData.Msg = "sukses menampilkan data invoice"
	resData.Data = listDataInvoice
	return &resData, nil
}
func (s *invoiceServiceImpl) DetailInvoice(ctx *gin.Context, request *request.GetInvoiceDetailRequest, uid string) (*response.GeneralResponse, error) {
	var resultData response.GeneralResponse
	var dataDetail response.ResultDetailInvoiceResponse
	selectDataInvoice, err := s.InvoiceRepository.SelectDetailInvoice(ctx, request.IdInvoice)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan  %s", err.Error()),
		}, nil
	}
	detailItemInvoice, err := s.InvoiceRepository.SelectDetailItemInvoice(ctx, request.IdInvoice)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan  %s", err.Error()),
		}, nil
	}
	var subTotal int
	for _, item := range detailItemInvoice {
		amount := item.Qty * item.UnitPrice
		subTotal += amount
	}
	presentase := float64(subTotal) * (10.0 / 100)
	grandTotal := float64(subTotal) + presentase

	dataDetail.InvoiceDetail.IdInvoice = selectDataInvoice.IdInvoice
	dataDetail.InvoiceDetail.Subject = selectDataInvoice.Subject
	dataDetail.InvoiceDetail.IssueDate = selectDataInvoice.IssueDate
	dataDetail.InvoiceDetail.DueDate = selectDataInvoice.DueDate
	dataDetail.CustomerInfo.CustomerName = selectDataInvoice.NamaCustomer
	dataDetail.CustomerInfo.Address = selectDataInvoice.Address
	dataDetail.DetailItem = detailItemInvoice
	dataDetail.InvoiceSummary.GrandTotal = grandTotal
	dataDetail.InvoiceSummary.SubTotal = float64(subTotal)
	dataDetail.InvoiceSummary.Tax = presentase
	dataDetail.InvoiceSummary.TotalItem = len(detailItemInvoice)
	dataDetail.Status = selectDataInvoice.Status

	resultData.Code = "200"
	resultData.Msg = "sukses menamplkan data detail invoice "
	resultData.Data = dataDetail
	return &resultData, nil
}

func (s *invoiceServiceImpl) EditInvoice(ctx *gin.Context, request *request.EditInvoiceRequest, uid string) (*response.GeneralResponse, error) {
	var resData response.GeneralResponse

	issueDate, _ := time.Parse(constant.DateYYYYMMDDSLASH, request.IssueDate)
	request.IssueDate = issueDate.Format(constant.DateYYYYMMDDDASH)
	dueDate, _ := time.Parse(constant.DateYYYYMMDDSLASH, request.DueDate)
	request.DueDate = dueDate.Format(constant.DateYYYYMMDDDASH)

	txBegin, err := s.InvoiceRepository.BeginTransaction()
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GeneralResponse{
			Code: "400",
			Msg:  fmt.Sprintf("Terjadi kesalahan %s", err.Error()),
		}, nil
	}
	for _, item := range request.DetailItem {
		amount := item.Qty * item.UnitPrice
		updateInvoiceItem := s.InvoiceRepository.EditInvoiceItem(ctx, &item, amount, request.IdInvoice, txBegin)
		if updateInvoiceItem != nil {
			helpers.LogError(ctx, updateInvoiceItem.Error(), uid)
			_ = txBegin.Rollback()
			return &response.GeneralResponse{
				Code: "401",
				Msg:  fmt.Sprintf("Terjadi kesalahan update item invoice %s", updateInvoiceItem.Error()),
			}, nil
		}
	}
	updateInvoice := s.InvoiceRepository.EditInvoice(ctx, request, txBegin)
	if updateInvoice != nil {
		helpers.LogError(ctx, updateInvoice.Error(), uid)
		_ = txBegin.Rollback()
		return &response.GeneralResponse{
			Code: "401",
			Msg:  fmt.Sprintf("Terjadi kesalahan update item invoice %s", updateInvoice.Error()),
		}, nil
	}
	_ = txBegin.Commit()
	resData.Code = "200"
	resData.Msg = "sukses edit invoice"
	return &resData, nil
}
