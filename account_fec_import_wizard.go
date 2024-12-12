package odoo

// AccountFecImportWizard represents account.fec.import.wizard model.
type AccountFecImportWizard struct {
	AttachmentId               *String    `xmlrpc:"attachment_id,omitempty"`
	AttachmentName             *String    `xmlrpc:"attachment_name,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	DocumentPrefix             *String    `xmlrpc:"document_prefix,omitempty"`
	DuplicateDocumentsHandling *Selection `xmlrpc:"duplicate_documents_handling,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountFecImportWizards represents array of account.fec.import.wizard model.
type AccountFecImportWizards []AccountFecImportWizard

// AccountFecImportWizardModel is the odoo model name.
const AccountFecImportWizardModel = "account.fec.import.wizard"

// Many2One convert AccountFecImportWizard to *Many2One.
func (afiw *AccountFecImportWizard) Many2One() *Many2One {
	return NewMany2One(afiw.Id.Get(), "")
}

// CreateAccountFecImportWizard creates a new account.fec.import.wizard model and returns its id.
func (c *Client) CreateAccountFecImportWizard(afiw *AccountFecImportWizard) (int64, error) {
	ids, err := c.CreateAccountFecImportWizards([]*AccountFecImportWizard{afiw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFecImportWizard creates a new account.fec.import.wizard model and returns its id.
func (c *Client) CreateAccountFecImportWizards(afiws []*AccountFecImportWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range afiws {
		vv = append(vv, v)
	}
	return c.Create(AccountFecImportWizardModel, vv, nil)
}

// UpdateAccountFecImportWizard updates an existing account.fec.import.wizard record.
func (c *Client) UpdateAccountFecImportWizard(afiw *AccountFecImportWizard) error {
	return c.UpdateAccountFecImportWizards([]int64{afiw.Id.Get()}, afiw)
}

// UpdateAccountFecImportWizards updates existing account.fec.import.wizard records.
// All records (represented by ids) will be updated by afiw values.
func (c *Client) UpdateAccountFecImportWizards(ids []int64, afiw *AccountFecImportWizard) error {
	return c.Update(AccountFecImportWizardModel, ids, afiw, nil)
}

// DeleteAccountFecImportWizard deletes an existing account.fec.import.wizard record.
func (c *Client) DeleteAccountFecImportWizard(id int64) error {
	return c.DeleteAccountFecImportWizards([]int64{id})
}

// DeleteAccountFecImportWizards deletes existing account.fec.import.wizard records.
func (c *Client) DeleteAccountFecImportWizards(ids []int64) error {
	return c.Delete(AccountFecImportWizardModel, ids)
}

// GetAccountFecImportWizard gets account.fec.import.wizard existing record.
func (c *Client) GetAccountFecImportWizard(id int64) (*AccountFecImportWizard, error) {
	afiws, err := c.GetAccountFecImportWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afiws)[0]), nil
}

// GetAccountFecImportWizards gets account.fec.import.wizard existing records.
func (c *Client) GetAccountFecImportWizards(ids []int64) (*AccountFecImportWizards, error) {
	afiws := &AccountFecImportWizards{}
	if err := c.Read(AccountFecImportWizardModel, ids, nil, afiws); err != nil {
		return nil, err
	}
	return afiws, nil
}

// FindAccountFecImportWizard finds account.fec.import.wizard record by querying it with criteria.
func (c *Client) FindAccountFecImportWizard(criteria *Criteria) (*AccountFecImportWizard, error) {
	afiws := &AccountFecImportWizards{}
	if err := c.SearchRead(AccountFecImportWizardModel, criteria, NewOptions().Limit(1), afiws); err != nil {
		return nil, err
	}
	return &((*afiws)[0]), nil
}

// FindAccountFecImportWizards finds account.fec.import.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFecImportWizards(criteria *Criteria, options *Options) (*AccountFecImportWizards, error) {
	afiws := &AccountFecImportWizards{}
	if err := c.SearchRead(AccountFecImportWizardModel, criteria, options, afiws); err != nil {
		return nil, err
	}
	return afiws, nil
}

// FindAccountFecImportWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFecImportWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFecImportWizardModel, criteria, options)
}

// FindAccountFecImportWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountFecImportWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFecImportWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
