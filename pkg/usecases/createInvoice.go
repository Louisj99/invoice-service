package usecases

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"invoice-service/pkg/entities"
	"strconv"
	"time"
)

type createInvoiceRequest struct {
	OrderID         string
	EmployeeID      string
	PaymentAmount   float64
	CustomerEmail   string
	CustomerName    string
	CustomerAddress string
	PaymentMethod   string
}

type createInvoiceResponse struct {
	PDFData []byte
}

type EmailSender interface {
	SendEmail(to string, subject string, body string, pdf []byte) error
}

type CreateInvoiceInterfacePostgres interface {
	GetOrder(orderID string) (entities.Order, error)
	GetEmployee(employeeID string) (entities.Employee, error)
}

func CreateInvoice(CreateInvoiceInterfacePostgres CreateInvoiceInterfacePostgres) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createInvoiceRequest
		var response createInvoiceResponse
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(400, "error bad request")
			return
		}

		order, err := CreateInvoiceInterfacePostgres.GetOrder(request.OrderID)
		if err != nil {
			c.JSON(500, "error getting order")
			return
		}
		employee, err := CreateInvoiceInterfacePostgres.GetEmployee(request.EmployeeID)
		if err != nil {
			c.JSON(500, "error getting employee")
			return
		}

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)

		pdf.Cell(40, 10, "Purchase Invoice")
		pdf.Ln(20)

		pdf.SetFont("Arial", "", 12)
		pdf.Cell(40, 10, "Company Name")
		pdf.Cell(40, 10, "Company Address")
		pdf.Cell(40, 10, "Employee Name"+employee.Name)
		pdf.Cell(40, 10, "Employee Email"+employee.Email)
		pdf.Ln(20)

		pdf.Cell(40, 10, "Customer Name: "+order.UserID)
		pdf.Cell(40, 10, "Customer Address: "+request.CustomerAddress)
		pdf.Cell(40, 10, "Customer Email: "+request.CustomerEmail)
		pdf.Ln(20)

		pdf.Cell(40, 10, "Invoice Date: "+time.Now().Format("2006-01-02"))
		pdf.Cell(40, 10, "Invoice ID: "+strconv.Itoa(order.OrderID))
		pdf.Ln(20)

		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(40, 10, "Item")
		pdf.Cell(40, 10, "Quantity")
		pdf.Cell(40, 10, "Price per unit")
		pdf.Ln(10)

		pdf.SetFont("Arial", "", 12)
		for _, item := range order.OrderItems {
			pdf.Cell(40, 10, item.ItemID)
			pdf.Cell(40, 10, fmt.Sprintf("%d", item.Quantity))
			pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.PricePerUnit))
			pdf.Ln(10)
		}

		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(40, 10, "Total Amount: "+fmt.Sprintf("%.2f", request.PaymentAmount))
		pdf.Ln(20)

		var buf bytes.Buffer
		err = pdf.Output(&buf)
		if err != nil {
			c.JSON(500, "error creating PDF")
			return
		}

		// Set the PDF data in the response
		response.PDFData = buf.Bytes()

		c.JSON(200, response)
	}
}
