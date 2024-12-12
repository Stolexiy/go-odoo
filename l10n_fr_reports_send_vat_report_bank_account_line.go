package odoo

// L10NFrReportsSendVatReportBankAccountLine represents l10n_fr_reports.send.vat.report.bank.account.line model.
type L10NFrReportsSendVatReportBankAccountLine struct {
	AccountNumber         *String   `xmlrpc:"account_number,omitempty"`
	BankBic               *String   `xmlrpc:"bank_bic,omitempty"`
	BankId                *Many2One `xmlrpc:"bank_id,omitempty"`
	BankPartnerId         *Many2One `xmlrpc:"bank_partner_id,omitempty"`
	CompanyPartnerId      *Many2One `xmlrpc:"company_partner_id,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId            *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	IsWronglyConfigured   *Bool     `xmlrpc:"is_wrongly_configured,omitempty"`
	L10NFrSendVatReportId *Many2One `xmlrpc:"l10n_fr_send_vat_report_id,omitempty"`
	VatAmount             *Float    `xmlrpc:"vat_amount,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// L10NFrReportsSendVatReportBankAccountLines represents array of l10n_fr_reports.send.vat.report.bank.account.line model.
type L10NFrReportsSendVatReportBankAccountLines []L10NFrReportsSendVatReportBankAccountLine

// L10NFrReportsSendVatReportBankAccountLineModel is the odoo model name.
const L10NFrReportsSendVatReportBankAccountLineModel = "l10n_fr_reports.send.vat.report.bank.account.line"

// Many2One convert L10NFrReportsSendVatReportBankAccountLine to *Many2One.
func (lsvrbal *L10NFrReportsSendVatReportBankAccountLine) Many2One() *Many2One {
	return NewMany2One(lsvrbal.Id.Get(), "")
}

// CreateL10NFrReportsSendVatReportBankAccountLine creates a new l10n_fr_reports.send.vat.report.bank.account.line model and returns its id.
func (c *Client) CreateL10NFrReportsSendVatReportBankAccountLine(lsvrbal *L10NFrReportsSendVatReportBankAccountLine) (int64, error) {
	ids, err := c.CreateL10NFrReportsSendVatReportBankAccountLines([]*L10NFrReportsSendVatReportBankAccountLine{lsvrbal})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NFrReportsSendVatReportBankAccountLine creates a new l10n_fr_reports.send.vat.report.bank.account.line model and returns its id.
func (c *Client) CreateL10NFrReportsSendVatReportBankAccountLines(lsvrbals []*L10NFrReportsSendVatReportBankAccountLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range lsvrbals {
		vv = append(vv, v)
	}
	return c.Create(L10NFrReportsSendVatReportBankAccountLineModel, vv, nil)
}

// UpdateL10NFrReportsSendVatReportBankAccountLine updates an existing l10n_fr_reports.send.vat.report.bank.account.line record.
func (c *Client) UpdateL10NFrReportsSendVatReportBankAccountLine(lsvrbal *L10NFrReportsSendVatReportBankAccountLine) error {
	return c.UpdateL10NFrReportsSendVatReportBankAccountLines([]int64{lsvrbal.Id.Get()}, lsvrbal)
}

// UpdateL10NFrReportsSendVatReportBankAccountLines updates existing l10n_fr_reports.send.vat.report.bank.account.line records.
// All records (represented by ids) will be updated by lsvrbal values.
func (c *Client) UpdateL10NFrReportsSendVatReportBankAccountLines(ids []int64, lsvrbal *L10NFrReportsSendVatReportBankAccountLine) error {
	return c.Update(L10NFrReportsSendVatReportBankAccountLineModel, ids, lsvrbal, nil)
}

// DeleteL10NFrReportsSendVatReportBankAccountLine deletes an existing l10n_fr_reports.send.vat.report.bank.account.line record.
func (c *Client) DeleteL10NFrReportsSendVatReportBankAccountLine(id int64) error {
	return c.DeleteL10NFrReportsSendVatReportBankAccountLines([]int64{id})
}

// DeleteL10NFrReportsSendVatReportBankAccountLines deletes existing l10n_fr_reports.send.vat.report.bank.account.line records.
func (c *Client) DeleteL10NFrReportsSendVatReportBankAccountLines(ids []int64) error {
	return c.Delete(L10NFrReportsSendVatReportBankAccountLineModel, ids)
}

// GetL10NFrReportsSendVatReportBankAccountLine gets l10n_fr_reports.send.vat.report.bank.account.line existing record.
func (c *Client) GetL10NFrReportsSendVatReportBankAccountLine(id int64) (*L10NFrReportsSendVatReportBankAccountLine, error) {
	lsvrbals, err := c.GetL10NFrReportsSendVatReportBankAccountLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lsvrbals)[0]), nil
}

// GetL10NFrReportsSendVatReportBankAccountLines gets l10n_fr_reports.send.vat.report.bank.account.line existing records.
func (c *Client) GetL10NFrReportsSendVatReportBankAccountLines(ids []int64) (*L10NFrReportsSendVatReportBankAccountLines, error) {
	lsvrbals := &L10NFrReportsSendVatReportBankAccountLines{}
	if err := c.Read(L10NFrReportsSendVatReportBankAccountLineModel, ids, nil, lsvrbals); err != nil {
		return nil, err
	}
	return lsvrbals, nil
}

// FindL10NFrReportsSendVatReportBankAccountLine finds l10n_fr_reports.send.vat.report.bank.account.line record by querying it with criteria.
func (c *Client) FindL10NFrReportsSendVatReportBankAccountLine(criteria *Criteria) (*L10NFrReportsSendVatReportBankAccountLine, error) {
	lsvrbals := &L10NFrReportsSendVatReportBankAccountLines{}
	if err := c.SearchRead(L10NFrReportsSendVatReportBankAccountLineModel, criteria, NewOptions().Limit(1), lsvrbals); err != nil {
		return nil, err
	}
	return &((*lsvrbals)[0]), nil
}

// FindL10NFrReportsSendVatReportBankAccountLines finds l10n_fr_reports.send.vat.report.bank.account.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrReportsSendVatReportBankAccountLines(criteria *Criteria, options *Options) (*L10NFrReportsSendVatReportBankAccountLines, error) {
	lsvrbals := &L10NFrReportsSendVatReportBankAccountLines{}
	if err := c.SearchRead(L10NFrReportsSendVatReportBankAccountLineModel, criteria, options, lsvrbals); err != nil {
		return nil, err
	}
	return lsvrbals, nil
}

// FindL10NFrReportsSendVatReportBankAccountLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrReportsSendVatReportBankAccountLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NFrReportsSendVatReportBankAccountLineModel, criteria, options)
}

// FindL10NFrReportsSendVatReportBankAccountLineId finds record id by querying it with criteria.
func (c *Client) FindL10NFrReportsSendVatReportBankAccountLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NFrReportsSendVatReportBankAccountLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
