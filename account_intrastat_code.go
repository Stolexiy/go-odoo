package odoo

// AccountIntrastatCode represents account.intrastat.code model.
type AccountIntrastatCode struct {
	Code              *String    `xmlrpc:"code,omitempty"`
	CountryId         *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description       *String    `xmlrpc:"description,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	ExpiryDate        *Time      `xmlrpc:"expiry_date,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	Name              *String    `xmlrpc:"name,omitempty"`
	StartDate         *Time      `xmlrpc:"start_date,omitempty"`
	SupplementaryUnit *Selection `xmlrpc:"supplementary_unit,omitempty"`
	Type              *Selection `xmlrpc:"type,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountIntrastatCodes represents array of account.intrastat.code model.
type AccountIntrastatCodes []AccountIntrastatCode

// AccountIntrastatCodeModel is the odoo model name.
const AccountIntrastatCodeModel = "account.intrastat.code"

// Many2One convert AccountIntrastatCode to *Many2One.
func (aic *AccountIntrastatCode) Many2One() *Many2One {
	return NewMany2One(aic.Id.Get(), "")
}

// CreateAccountIntrastatCode creates a new account.intrastat.code model and returns its id.
func (c *Client) CreateAccountIntrastatCode(aic *AccountIntrastatCode) (int64, error) {
	ids, err := c.CreateAccountIntrastatCodes([]*AccountIntrastatCode{aic})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountIntrastatCode creates a new account.intrastat.code model and returns its id.
func (c *Client) CreateAccountIntrastatCodes(aics []*AccountIntrastatCode) ([]int64, error) {
	var vv []interface{}
	for _, v := range aics {
		vv = append(vv, v)
	}
	return c.Create(AccountIntrastatCodeModel, vv, nil)
}

// UpdateAccountIntrastatCode updates an existing account.intrastat.code record.
func (c *Client) UpdateAccountIntrastatCode(aic *AccountIntrastatCode) error {
	return c.UpdateAccountIntrastatCodes([]int64{aic.Id.Get()}, aic)
}

// UpdateAccountIntrastatCodes updates existing account.intrastat.code records.
// All records (represented by ids) will be updated by aic values.
func (c *Client) UpdateAccountIntrastatCodes(ids []int64, aic *AccountIntrastatCode) error {
	return c.Update(AccountIntrastatCodeModel, ids, aic, nil)
}

// DeleteAccountIntrastatCode deletes an existing account.intrastat.code record.
func (c *Client) DeleteAccountIntrastatCode(id int64) error {
	return c.DeleteAccountIntrastatCodes([]int64{id})
}

// DeleteAccountIntrastatCodes deletes existing account.intrastat.code records.
func (c *Client) DeleteAccountIntrastatCodes(ids []int64) error {
	return c.Delete(AccountIntrastatCodeModel, ids)
}

// GetAccountIntrastatCode gets account.intrastat.code existing record.
func (c *Client) GetAccountIntrastatCode(id int64) (*AccountIntrastatCode, error) {
	aics, err := c.GetAccountIntrastatCodes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aics)[0]), nil
}

// GetAccountIntrastatCodes gets account.intrastat.code existing records.
func (c *Client) GetAccountIntrastatCodes(ids []int64) (*AccountIntrastatCodes, error) {
	aics := &AccountIntrastatCodes{}
	if err := c.Read(AccountIntrastatCodeModel, ids, nil, aics); err != nil {
		return nil, err
	}
	return aics, nil
}

// FindAccountIntrastatCode finds account.intrastat.code record by querying it with criteria.
func (c *Client) FindAccountIntrastatCode(criteria *Criteria) (*AccountIntrastatCode, error) {
	aics := &AccountIntrastatCodes{}
	if err := c.SearchRead(AccountIntrastatCodeModel, criteria, NewOptions().Limit(1), aics); err != nil {
		return nil, err
	}
	return &((*aics)[0]), nil
}

// FindAccountIntrastatCodes finds account.intrastat.code records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountIntrastatCodes(criteria *Criteria, options *Options) (*AccountIntrastatCodes, error) {
	aics := &AccountIntrastatCodes{}
	if err := c.SearchRead(AccountIntrastatCodeModel, criteria, options, aics); err != nil {
		return nil, err
	}
	return aics, nil
}

// FindAccountIntrastatCodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountIntrastatCodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountIntrastatCodeModel, criteria, options)
}

// FindAccountIntrastatCodeId finds record id by querying it with criteria.
func (c *Client) FindAccountIntrastatCodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountIntrastatCodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
