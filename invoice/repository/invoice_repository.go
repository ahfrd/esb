package repository

import (
	"database/sql"
	"esb-assesment/entity"
	"esb-assesment/entity/request"
	"esb-assesment/entity/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

type invoiceRepository struct {
	db *sql.DB
}

// ExampleRepository implements entity.InvoiceRepository.

func NewInvoiceRepository(db *sql.DB) entity.InvoiceRepository {
	return &invoiceRepository{
		db: db,
	}
}
func (r *invoiceRepository) BeginTransaction() (*sql.Tx, error) {
	tx, err := r.db.Begin()
	return tx, err
}
func (r *invoiceRepository) LastIdTransaction(ctx *gin.Context) (int, error) {
	var idInvoice int
	q := "select id_invoice from tbl_invoice order by id_invoice desc limit 1"
	if err := r.db.QueryRow(q).Scan(&idInvoice); err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return idInvoice, nil
}
func (r *invoiceRepository) InsertInvoice(ctx *gin.Context, request *request.InsertInvoiceRequest, tx *sql.Tx) error {
	q := `INSERT INTO tbl_invoice (id_invoice,subject,issued_date,due_date,id_customer,status,total_item) values (?,?,?,?,?,?,?)`
	_, err := tx.ExecContext(ctx, q, request.IdInvoice, request.Subject, request.IssueDate, request.DueDate, request.IdCustomer, request.Status, request.TotalItem)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (r *invoiceRepository) InsertInvoiceItem(ctx *gin.Context, idInvoice string, request *request.DetailItemRequest, amount int, tx *sql.Tx) error {
	q := `INSERT INTO tbl_invoice_item (id_invoice,id_item,qty,unit_price,amount) values (?,?,?,?,?)`
	_, err := tx.ExecContext(ctx, q, idInvoice, request.IdItem, request.Qty, request.UnitPrice, amount)
	if err != nil {
		return err
	}
	return nil
}

func (r *invoiceRepository) EditInvoice(ctx *gin.Context, request *request.EditInvoiceRequest, tx *sql.Tx) error {
	q := fmt.Sprintf(`update tbl_invoice set subject = '%s', issued_date = '%s', due_date = '%s', id_customer = '%s' where id_invoice = '%s' `,
		request.Subject, request.IssueDate, request.DueDate, request.IdCustomer, request.IdInvoice)

	fmt.Println(q)
	if _, err := tx.ExecContext(ctx, q); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func (r *invoiceRepository) EditInvoiceItem(ctx *gin.Context, request *request.DetailItemRequest, amount int, idInvoice string, tx *sql.Tx) error {
	q := fmt.Sprintf(`update tbl_invoice_item set id_item = %s,qty =%d,unit_price = %d,amount = %d where id_invoice = %s`, request.IdItem, request.Qty, request.UnitPrice, amount, idInvoice)
	fmt.Println(q)
	if _, err := tx.ExecContext(ctx, q); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func (r *invoiceRepository) SelectDetailItemInvoice(ctx *gin.Context, idInvoice string) ([]response.DetailItemResponse, error) {
	var results []response.DetailItemResponse
	q := fmt.Sprintf(`select t1.item_name, t0.qty, t0.unit_price,t0.amount
					  from tbl_invoice_item t0
					  left join tbl_item t1 on t1.id = t0.id_item
					  where 
					  t0.id_invoice = %s
					  order by t0.id asc`, idInvoice)
	result, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = result.Close() }()
	for result.Next() {
		var dataItem response.DetailItemResponse
		if err = result.Scan(&dataItem.ItemName, &dataItem.Qty, &dataItem.UnitPrice, &dataItem.Amount); err != nil {
			return nil, err
		}
		results = append(results, dataItem)
	}
	return results, nil
}
func (r *invoiceRepository) SelectDetailInvoice(ctx *gin.Context, idInvoice string) (*response.ResultDataInvoiceResponse, error) {
	var resultData response.ResultDataInvoiceResponse
	var idInvoices sql.NullString
	var subject sql.NullString
	var issuDate sql.NullString
	var dueDate sql.NullString
	var customerName sql.NullString
	var status sql.NullString
	var totalItem sql.NullString
	var address sql.NullString
	q := fmt.Sprintf(`select t0.id_invoice ,t0.subject,t0.issued_date ,t0.due_date ,t2.customer_name,t0.status,t0.total_item,t2.address
					  from tbl_invoice t0
					  left join tbl_customer t2 on t2.id  = t0.id_customer
					  where t0.id_invoice = %s`, idInvoice)
	fmt.Println(q)
	if err := r.db.QueryRowContext(ctx, q).Scan(&idInvoices, &subject, &issuDate, &dueDate, &customerName, &status, &totalItem, &address); err != nil {
		return nil, err
	}
	resultData.IdInvoice = idInvoices.String
	resultData.Subject = subject.String
	resultData.IssueDate = issuDate.String
	resultData.DueDate = dueDate.String
	resultData.NamaCustomer = customerName.String
	resultData.Status = status.String
	resultData.TotalItem = totalItem.String
	resultData.Address = address.String
	return &resultData, nil

}
func (r *invoiceRepository) CountInvoiceData(ctx *gin.Context, request *request.GetAllInvoiceRequest) (string, error) {
	var count string
	filter := r.filterQuery(request)
	q := fmt.Sprintf(`select count(*) as count from tbl_invoice t0
					  left join tbl_customer t2 on t2.id  = t0.id_customer 
	 				  where 1 = 1 %s`, filter)
	fmt.Println(q)
	if err := r.db.QueryRowContext(ctx, q).Scan(&count); err != nil {
		return "", err
	}
	return count, nil
}
func (r *invoiceRepository) SelectInvoice(ctx *gin.Context, request *request.GetAllInvoiceRequest) ([]response.ResultDataInvoiceResponse, error) {
	var results []response.ResultDataInvoiceResponse
	where := r.filterQuery(request)
	q := fmt.Sprintf(`select t0.id_invoice ,t0.subject,t0.issued_date ,t0.due_date ,t0.total_item,t2.customer_name,t0.status
				      from tbl_invoice t0
					  left join tbl_customer t2 on t2.id  = t0.id_customer 
					  where 1 = 1
					  %s
					  order by t0.id_invoice desc 
					  limit %s,%s
					  `, where, request.PageNumber, request.PageSize)
	result, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = result.Close() }()
	for result.Next() {
		var dataInvoice response.ResultDataInvoiceResponse
		if err = result.Scan(&dataInvoice.IdInvoice, &dataInvoice.Subject, &dataInvoice.IssueDate,
			&dataInvoice.DueDate, &dataInvoice.TotalItem, &dataInvoice.NamaCustomer, &dataInvoice.Status); err != nil {
			return nil, err
		}
		results = append(results, dataInvoice)

	}

	return results, nil
}

func (r *invoiceRepository) filterQuery(request *request.GetAllInvoiceRequest) string {
	var where string
	if request.IdInvoice != "" {
		where += fmt.Sprintf(" and t0.id_invoice = %s", request.IdInvoice)
	}
	if request.IssueDate != "" {
		where += fmt.Sprintf(" and t0.issued_date = %s", request.IssueDate)

	}
	if request.DueDate != "" {
		where += fmt.Sprintf(" and t0.due_date = %s", request.DueDate)

	}
	if request.NamaCustomer != "" {
		where += fmt.Sprintf(" and t2.customer_name = %s", request.NamaCustomer)

	}
	if request.Status != "" {
		where += fmt.Sprintf(" and t0.status = %s", request.Status)

	}

	if request.TotalItem != "" {
		where += fmt.Sprintf(" and t0.total_item = %s", request.TotalItem)
	}
	return where
}
