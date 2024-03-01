package usecases

type createInvoiceRequest struct {
	OrderID       string
	PaymentAmount float64
	CustomerEmail string
	EmployeeID    string
}

type createInvoiceResponse struct {
	PDFData []byte
}

type CreateInvoiceInterface interface {
	CreateInvoice(request createInvoiceRequest) ([]byte, error)
}
type EmailSender interface {
	SendEmail(to string, subject string, body string, pdf []byte) error
}

type CreateInvoiceInterfacePostgres interface {
	GetOrder(orderID string) (Order, error)
	GetEmployee(employeeID string) (Employee, error)
	CreateInvoice(orderID string, paymentAmount float64, customerEmail string, employeeID string) error
}
