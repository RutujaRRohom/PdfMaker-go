package service

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(employee []Employee) (err error) {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set up basic styling for the PDF
	pdf.SetFont("Arial", "", 14)
	pdf.SetMargins(20, 20, 20)
	pdf.SetTitle("Employee Report", true)

	// Generate table header
	pdf.CellFormat(40, 10, "ID", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "First Name", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Last Name", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Email", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Phone Number", "1", 1, "", false, 0, "")

	// Generate table rows with employee data
	for _, employee := range employee {
		pdf.CellFormat(40, 10, strconv.Itoa(employee.EmployeeID), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, employee.FirstName, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, employee.LastName, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, employee.Email, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, employee.PhoneNumber, "1", 1, "", false, 0, "")
	}

	// Generate the PDF file
	err = pdf.OutputFileAndClose("employeeReport.pdf")
	if err != nil {
		return err
	}

	return
}
