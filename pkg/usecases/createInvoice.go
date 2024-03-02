package usecases

import (
	"github.com/gin-gonic/gin"
	"invoice-service/pkg/entities"
)

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
	GetOrder(orderID string) (entities.Order, error)
	GetEmployee(employeeID string) (entities.Employee, error)
	CreateInvoice(invoice entities.Invoice) ([]byte, error)
}

func CreateInvoice(CreateInvoiceInterface CreateInvoiceInterface, CreateInvoiceInterfacePostgres CreateInvoiceInterfacePostgres) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createInvoiceRequest
		var response createInvoiceResponse
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, "error bad request")
			return
		}

		c.JSON(200, response)
	}
}
