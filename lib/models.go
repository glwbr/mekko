package juice

import (
	"time"
)

// Invoice represents a complete electronic invoice
type Invoice struct {
	AccessKey     string
	IssueDate     time.Time
	TotalValue    float64
	Merchant      Merchant
	Customer      *Customer // NOTE: optional customer info
	Items         []InvoiceItem
	PaymentMethod string
	Taxes         Taxes
	Raw           *RawInvoiceData
}

type Merchant struct {
	Name         string
	CNPJ         string
	Address      Address
	StateRegCode string
}

type Customer struct {
	Name    string
	CPF     string
	Address *Address
}

type InvoiceItem struct {
	Code        string
	Description string
	Quantity    float64
	UnitPrice   float64
	TotalPrice  float64
	TaxCode     string
}

type Address struct {
	Street     string
	Number     string
	District   string
	City       string
	State      string
	PostalCode string
}

type Taxes struct {
	ICMS   float64
	PIS    float64
	COFINS float64
	Total  float64
}

// NOTE: RawInvoiceData stores the unparsed data for debugging
type RawInvoiceData struct {
	HTML string
}
