package odoo

// L10NFrReportsSendVatReport represents l10n_fr_reports.send.vat.report model.
type L10NFrReportsSendVatReport struct {
	BankAccountLineCount        *Int       `xmlrpc:"bank_account_line_count,omitempty"`
	BankAccountLineIds          *Relation  `xmlrpc:"bank_account_line_ids,omitempty"`
	ComputedVatAmount           *Float     `xmlrpc:"computed_vat_amount,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	HasWronglyConfiguredAccount *Bool      `xmlrpc:"has_wrongly_configured_account,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsVatDue                    *Bool      `xmlrpc:"is_vat_due,omitempty"`
	Recipient                   *Selection `xmlrpc:"recipient,omitempty"`
	ReportId                    *Many2One  `xmlrpc:"report_id,omitempty"`
	TestInterchange             *Bool      `xmlrpc:"test_interchange,omitempty"`
	VatAmount                   *Float     `xmlrpc:"vat_amount,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// L10NFrReportsSendVatReports represents array of l10n_fr_reports.send.vat.report model.
type L10NFrReportsSendVatReports []L10NFrReportsSendVatReport

// L10NFrReportsSendVatReportModel is the odoo model name.
const L10NFrReportsSendVatReportModel = "l10n_fr_reports.send.vat.report"

// Many2One convert L10NFrReportsSendVatReport to *Many2One.
func (lsvr *L10NFrReportsSendVatReport) Many2One() *Many2One {
	return NewMany2One(lsvr.Id.Get(), "")
}

// CreateL10NFrReportsSendVatReport creates a new l10n_fr_reports.send.vat.report model and returns its id.
func (c *Client) CreateL10NFrReportsSendVatReport(lsvr *L10NFrReportsSendVatReport) (int64, error) {
	ids, err := c.CreateL10NFrReportsSendVatReports([]*L10NFrReportsSendVatReport{lsvr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NFrReportsSendVatReport creates a new l10n_fr_reports.send.vat.report model and returns its id.
func (c *Client) CreateL10NFrReportsSendVatReports(lsvrs []*L10NFrReportsSendVatReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range lsvrs {
		vv = append(vv, v)
	}
	return c.Create(L10NFrReportsSendVatReportModel, vv, nil)
}

// UpdateL10NFrReportsSendVatReport updates an existing l10n_fr_reports.send.vat.report record.
func (c *Client) UpdateL10NFrReportsSendVatReport(lsvr *L10NFrReportsSendVatReport) error {
	return c.UpdateL10NFrReportsSendVatReports([]int64{lsvr.Id.Get()}, lsvr)
}

// UpdateL10NFrReportsSendVatReports updates existing l10n_fr_reports.send.vat.report records.
// All records (represented by ids) will be updated by lsvr values.
func (c *Client) UpdateL10NFrReportsSendVatReports(ids []int64, lsvr *L10NFrReportsSendVatReport) error {
	return c.Update(L10NFrReportsSendVatReportModel, ids, lsvr, nil)
}

// DeleteL10NFrReportsSendVatReport deletes an existing l10n_fr_reports.send.vat.report record.
func (c *Client) DeleteL10NFrReportsSendVatReport(id int64) error {
	return c.DeleteL10NFrReportsSendVatReports([]int64{id})
}

// DeleteL10NFrReportsSendVatReports deletes existing l10n_fr_reports.send.vat.report records.
func (c *Client) DeleteL10NFrReportsSendVatReports(ids []int64) error {
	return c.Delete(L10NFrReportsSendVatReportModel, ids)
}

// GetL10NFrReportsSendVatReport gets l10n_fr_reports.send.vat.report existing record.
func (c *Client) GetL10NFrReportsSendVatReport(id int64) (*L10NFrReportsSendVatReport, error) {
	lsvrs, err := c.GetL10NFrReportsSendVatReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lsvrs)[0]), nil
}

// GetL10NFrReportsSendVatReports gets l10n_fr_reports.send.vat.report existing records.
func (c *Client) GetL10NFrReportsSendVatReports(ids []int64) (*L10NFrReportsSendVatReports, error) {
	lsvrs := &L10NFrReportsSendVatReports{}
	if err := c.Read(L10NFrReportsSendVatReportModel, ids, nil, lsvrs); err != nil {
		return nil, err
	}
	return lsvrs, nil
}

// FindL10NFrReportsSendVatReport finds l10n_fr_reports.send.vat.report record by querying it with criteria.
func (c *Client) FindL10NFrReportsSendVatReport(criteria *Criteria) (*L10NFrReportsSendVatReport, error) {
	lsvrs := &L10NFrReportsSendVatReports{}
	if err := c.SearchRead(L10NFrReportsSendVatReportModel, criteria, NewOptions().Limit(1), lsvrs); err != nil {
		return nil, err
	}
	return &((*lsvrs)[0]), nil
}

// FindL10NFrReportsSendVatReports finds l10n_fr_reports.send.vat.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrReportsSendVatReports(criteria *Criteria, options *Options) (*L10NFrReportsSendVatReports, error) {
	lsvrs := &L10NFrReportsSendVatReports{}
	if err := c.SearchRead(L10NFrReportsSendVatReportModel, criteria, options, lsvrs); err != nil {
		return nil, err
	}
	return lsvrs, nil
}

// FindL10NFrReportsSendVatReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrReportsSendVatReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NFrReportsSendVatReportModel, criteria, options)
}

// FindL10NFrReportsSendVatReportId finds record id by querying it with criteria.
func (c *Client) FindL10NFrReportsSendVatReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NFrReportsSendVatReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
