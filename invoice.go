package harvest

import (
	"fmt"
	"time"
)

type Invoice struct {
	ID                 int64       `json:"id"`
	Client             *ClientStub `json:"client,omitempty"`
	PeriodStart        Date        `json:"period_start"`
	PeriodEnd          Date        `json:"period_end"`
	Number             string      `json:"number"`
	IssueDate          Date        `json:"issue_date"`
	DueDate            Date        `json:"due_date"`
	Amount             float64     `json:"amount"`
	Currency           string      `json:"currency"`
	State              string      `json:"state"`
	Notes              string      `json:"notes"`
	PurchaseOrder      string      `json:"purchase_order"`
	DueAmount          float64     `json:"due_amount"`
	DueAtHumanFormat   string      `json:"due_at_human_format"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	SentAt             time.Time   `json:"sent_at"`
	PaidAt             time.Time   `json:"paid_at"`
	ClosedAt           time.Time   `json:"closed_at"`
	Tax                float64     `json:"tax"`
	TaxAmount          float64     `json:"tax_amount"`
	Subject            string      `json:"subject"`
	RecurringInvoiceID int64       `json:"recurring_invoice_id"`
	Tax2               float64     `json:"tax2"`
	Tax2Amount         float64     `json:"tax2_amount"`
	ClientKey          string      `json:"client_key"`
	EstimateID         int64       `json:"estimate_id"`
	Discount           float64     `json:"discount"`
	DiscountAmount     float64     `json:"discount_amount"`
	RetainerID         int64       `json:"retainer_id"`
	CreatedByID        int64       `json:"created_by_id"`
	LineItems          []*LineItem `json:"line_items"`
}

type LineItem struct {
	ID          int64        `json:"id"`
	Kind        string       `json:"kind"`
	Description string       `json:"description"`
	Quantity    float64      `json:"quantity"`
	UnitPrice   float64      `json:"unit_price"`
	Amount      float64      `json:"amount"`
	Taxed       bool         `json:"taxed"`
	Taxed2      bool         `json:"taxed2"`
	Project     *ProjectStub `json:"project"`
}

func (a *API) GetInvoice(invoiceID int64, args Arguments) (invoice *Invoice, err error) {
	invoice = &Invoice{}
	path := fmt.Sprintf("/invoices/%d", invoiceID)
	err = a.Get(path, args, &invoice)
	return invoice, err
}

func (a *API) GetInvoices(args Arguments) (invoices []*Invoice, err error) {
	var invoicesResponse struct {
		PagedResponse
		Invoices []*Invoice `json:"invoices"`
	}
	invoices = make([]*Invoice, 0)
	err = a.GetPaginated("/invoices", args, &invoicesResponse, func() {
		for _, i := range invoicesResponse.Invoices {
			invoices = append(invoices, i)
		}
		invoicesResponse.Invoices = make([]*Invoice, 0)
	})
	return invoices, err
}
