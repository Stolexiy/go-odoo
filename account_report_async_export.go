package odoo

// AccountReportAsyncExport represents account.report.async.export model.
type AccountReportAsyncExport struct {
	AttachmentIds  *Relation  `xmlrpc:"attachment_ids,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom       *Time      `xmlrpc:"date_from,omitempty"`
	DateTo         *Time      `xmlrpc:"date_to,omitempty"`
	DeclarationUid *String    `xmlrpc:"declaration_uid,omitempty"`
	DepositUid     *String    `xmlrpc:"deposit_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Message        *String    `xmlrpc:"message,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	Recipient      *Selection `xmlrpc:"recipient,omitempty"`
	ReportId       *Many2One  `xmlrpc:"report_id,omitempty"`
	State          *Selection `xmlrpc:"state,omitempty"`
	Step1Logs      *String    `xmlrpc:"step_1_logs,omitempty"`
	Step2Logs      *String    `xmlrpc:"step_2_logs,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReportAsyncExports represents array of account.report.async.export model.
type AccountReportAsyncExports []AccountReportAsyncExport

// AccountReportAsyncExportModel is the odoo model name.
const AccountReportAsyncExportModel = "account.report.async.export"

// Many2One convert AccountReportAsyncExport to *Many2One.
func (arae *AccountReportAsyncExport) Many2One() *Many2One {
	return NewMany2One(arae.Id.Get(), "")
}

// CreateAccountReportAsyncExport creates a new account.report.async.export model and returns its id.
func (c *Client) CreateAccountReportAsyncExport(arae *AccountReportAsyncExport) (int64, error) {
	ids, err := c.CreateAccountReportAsyncExports([]*AccountReportAsyncExport{arae})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportAsyncExport creates a new account.report.async.export model and returns its id.
func (c *Client) CreateAccountReportAsyncExports(araes []*AccountReportAsyncExport) ([]int64, error) {
	var vv []interface{}
	for _, v := range araes {
		vv = append(vv, v)
	}
	return c.Create(AccountReportAsyncExportModel, vv, nil)
}

// UpdateAccountReportAsyncExport updates an existing account.report.async.export record.
func (c *Client) UpdateAccountReportAsyncExport(arae *AccountReportAsyncExport) error {
	return c.UpdateAccountReportAsyncExports([]int64{arae.Id.Get()}, arae)
}

// UpdateAccountReportAsyncExports updates existing account.report.async.export records.
// All records (represented by ids) will be updated by arae values.
func (c *Client) UpdateAccountReportAsyncExports(ids []int64, arae *AccountReportAsyncExport) error {
	return c.Update(AccountReportAsyncExportModel, ids, arae, nil)
}

// DeleteAccountReportAsyncExport deletes an existing account.report.async.export record.
func (c *Client) DeleteAccountReportAsyncExport(id int64) error {
	return c.DeleteAccountReportAsyncExports([]int64{id})
}

// DeleteAccountReportAsyncExports deletes existing account.report.async.export records.
func (c *Client) DeleteAccountReportAsyncExports(ids []int64) error {
	return c.Delete(AccountReportAsyncExportModel, ids)
}

// GetAccountReportAsyncExport gets account.report.async.export existing record.
func (c *Client) GetAccountReportAsyncExport(id int64) (*AccountReportAsyncExport, error) {
	araes, err := c.GetAccountReportAsyncExports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*araes)[0]), nil
}

// GetAccountReportAsyncExports gets account.report.async.export existing records.
func (c *Client) GetAccountReportAsyncExports(ids []int64) (*AccountReportAsyncExports, error) {
	araes := &AccountReportAsyncExports{}
	if err := c.Read(AccountReportAsyncExportModel, ids, nil, araes); err != nil {
		return nil, err
	}
	return araes, nil
}

// FindAccountReportAsyncExport finds account.report.async.export record by querying it with criteria.
func (c *Client) FindAccountReportAsyncExport(criteria *Criteria) (*AccountReportAsyncExport, error) {
	araes := &AccountReportAsyncExports{}
	if err := c.SearchRead(AccountReportAsyncExportModel, criteria, NewOptions().Limit(1), araes); err != nil {
		return nil, err
	}
	return &((*araes)[0]), nil
}

// FindAccountReportAsyncExports finds account.report.async.export records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportAsyncExports(criteria *Criteria, options *Options) (*AccountReportAsyncExports, error) {
	araes := &AccountReportAsyncExports{}
	if err := c.SearchRead(AccountReportAsyncExportModel, criteria, options, araes); err != nil {
		return nil, err
	}
	return araes, nil
}

// FindAccountReportAsyncExportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportAsyncExportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportAsyncExportModel, criteria, options)
}

// FindAccountReportAsyncExportId finds record id by querying it with criteria.
func (c *Client) FindAccountReportAsyncExportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportAsyncExportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
