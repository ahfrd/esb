package request

type AddInvoiceRequest struct {
	Subject      string                `json:"subject"`
	IssueDate    string                `json:"issueDate"`
	DueDate      string                `json:"dueDate"`
	CustomerInfo DetailCustomerRequest `json:"customerInfo"`
	DetailItem   []DetailItemRequest   `json:"detailItem"`
}

type DetailCustomerRequest struct {
	IdCustomer string `json:"idCustomer"`
	Address    string `json:"address"`
}

type DetailItemRequest struct {
	IdItem    string `json:"idItem"`
	Qty       int    `json:"qty"`
	UnitPrice int    `json:"unitPrice"`
}

type InsertInvoiceRequest struct {
	IdInvoice  string `json:"idInvoice"`
	Subject    string `json:"subject"`
	IssueDate  string `json:"issueDate"`
	DueDate    string `json:"dueDate"`
	IdCustomer string `json:"idCustomer"`
	Status     string `json:"status"`
	TotalItem  int    `json:"totalItem"`
}
type GetAllInvoiceRequest struct {
	PageNumber   string `json:"pageNumber" binding:"required"`
	PageSize     string `json:"pageSize" binding:"required"`
	IdInvoice    string `json:"idInvoice"`
	Subject      string `json:"subject"`
	IssueDate    string `json:"issueDate"`
	DueDate      string `json:"dueDate"`
	NamaCustomer string `json:"namaCustomer"`
	TotalItem    string `json:"totalItem"`
	Status       string `json:"status"`
}

type GetInvoiceDetailRequest struct {
	IdInvoice string `json:"idInvoice"`
}

type EditInvoiceRequest struct {
	IdInvoice  string              `json:"idInvoice"`
	Subject    string              `json:"subject"`
	IssueDate  string              `json:"issueDate"`
	DueDate    string              `json:"dueDate"`
	IdCustomer string              `json:"idCustomer"`
	DetailItem []DetailItemRequest `json:"detailItem"`
}
