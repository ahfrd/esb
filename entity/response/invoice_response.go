package response

type LoginResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GeneralResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type SummaryInvoiceResponse struct {
	TotalItem  int     `json:"totalItem"`
	SubTotal   float64 `json:"subTotal"`
	Tax        float64 `json:"tax"`
	GrandTotal float64 `json:"grandTotal"`
}

type ResultDataInvoiceResponse struct {
	IdInvoice    string `json:"idInvoice"`
	Subject      string `json:"subject"`
	IssueDate    string `json:"issueDate"`
	DueDate      string `json:"dueDate"`
	NamaCustomer string `json:"namaCustomer"`
	TotalItem    string `json:"totalItem"`
	Status       string `json:"status"`
	Address      string `json:"address"`
}

type CountData struct {
	TotalRecord   string `json:"total_record"`
	TotalPage     string `json:"total_page"`
	RecordPerPage string `json:"record_per_page"`
	CurrentPage   string `json:"current_page"`
	StartRecord   string `json:"start_record"`
	FirstRecord   string `json:"first_record"`
}

type ListInvoiceResponse struct {
	ListData   []ResultDataInvoiceResponse `json:"listData"`
	Pagination CountData                   `json:"pagination"`
}

type DetailItemResponse struct {
	ItemName  string `json:"itemName"`
	Qty       int    `json:"qty"`
	UnitPrice int    `json:"unitPrice"`
	Amount    int    `json:"amount"`
}

type InvoiceDetailEntity struct {
	IdInvoice string `json:"idInvoice"`
	Subject   string `json:"subject"`
	IssueDate string `json:"issueDate"`
	DueDate   string `json:"dueDate"`
}

type CustomerInfoEntity struct {
	CustomerName string `json:"customerName"`
	Address      string `json:"address"`
}

type ResultDetailInvoiceResponse struct {
	InvoiceDetail  InvoiceDetailEntity    `json:"invoiceDetail"`
	CustomerInfo   CustomerInfoEntity     `json:"customerInfo"`
	DetailItem     []DetailItemResponse   `json:"detailItem"`
	InvoiceSummary SummaryInvoiceResponse `json:"invoiceSummary"`
	Status         string                 `json:"status"`
}
