package odoo

// AccountSaleClosing represents account.sale.closing model.
type AccountSaleClosing struct {
	CompanyId        *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	CumulativeTotal  *Float     `xmlrpc:"cumulative_total,omitempty"`
	CurrencyId       *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateClosingStart *Time      `xmlrpc:"date_closing_start,omitempty"`
	DateClosingStop  *Time      `xmlrpc:"date_closing_stop,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Frequency        *Selection `xmlrpc:"frequency,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	LastOrderHash    *String    `xmlrpc:"last_order_hash,omitempty"`
	LastOrderId      *Many2One  `xmlrpc:"last_order_id,omitempty"`
	Name             *String    `xmlrpc:"name,omitempty"`
	SequenceNumber   *Int       `xmlrpc:"sequence_number,omitempty"`
	TotalInterval    *Float     `xmlrpc:"total_interval,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountSaleClosings represents array of account.sale.closing model.
type AccountSaleClosings []AccountSaleClosing

// AccountSaleClosingModel is the odoo model name.
const AccountSaleClosingModel = "account.sale.closing"

// Many2One convert AccountSaleClosing to *Many2One.
func (asc *AccountSaleClosing) Many2One() *Many2One {
	return NewMany2One(asc.Id.Get(), "")
}

// CreateAccountSaleClosing creates a new account.sale.closing model and returns its id.
func (c *Client) CreateAccountSaleClosing(asc *AccountSaleClosing) (int64, error) {
	ids, err := c.CreateAccountSaleClosings([]*AccountSaleClosing{asc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountSaleClosing creates a new account.sale.closing model and returns its id.
func (c *Client) CreateAccountSaleClosings(ascs []*AccountSaleClosing) ([]int64, error) {
	var vv []interface{}
	for _, v := range ascs {
		vv = append(vv, v)
	}
	return c.Create(AccountSaleClosingModel, vv, nil)
}

// UpdateAccountSaleClosing updates an existing account.sale.closing record.
func (c *Client) UpdateAccountSaleClosing(asc *AccountSaleClosing) error {
	return c.UpdateAccountSaleClosings([]int64{asc.Id.Get()}, asc)
}

// UpdateAccountSaleClosings updates existing account.sale.closing records.
// All records (represented by ids) will be updated by asc values.
func (c *Client) UpdateAccountSaleClosings(ids []int64, asc *AccountSaleClosing) error {
	return c.Update(AccountSaleClosingModel, ids, asc, nil)
}

// DeleteAccountSaleClosing deletes an existing account.sale.closing record.
func (c *Client) DeleteAccountSaleClosing(id int64) error {
	return c.DeleteAccountSaleClosings([]int64{id})
}

// DeleteAccountSaleClosings deletes existing account.sale.closing records.
func (c *Client) DeleteAccountSaleClosings(ids []int64) error {
	return c.Delete(AccountSaleClosingModel, ids)
}

// GetAccountSaleClosing gets account.sale.closing existing record.
func (c *Client) GetAccountSaleClosing(id int64) (*AccountSaleClosing, error) {
	ascs, err := c.GetAccountSaleClosings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ascs)[0]), nil
}

// GetAccountSaleClosings gets account.sale.closing existing records.
func (c *Client) GetAccountSaleClosings(ids []int64) (*AccountSaleClosings, error) {
	ascs := &AccountSaleClosings{}
	if err := c.Read(AccountSaleClosingModel, ids, nil, ascs); err != nil {
		return nil, err
	}
	return ascs, nil
}

// FindAccountSaleClosing finds account.sale.closing record by querying it with criteria.
func (c *Client) FindAccountSaleClosing(criteria *Criteria) (*AccountSaleClosing, error) {
	ascs := &AccountSaleClosings{}
	if err := c.SearchRead(AccountSaleClosingModel, criteria, NewOptions().Limit(1), ascs); err != nil {
		return nil, err
	}
	return &((*ascs)[0]), nil
}

// FindAccountSaleClosings finds account.sale.closing records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSaleClosings(criteria *Criteria, options *Options) (*AccountSaleClosings, error) {
	ascs := &AccountSaleClosings{}
	if err := c.SearchRead(AccountSaleClosingModel, criteria, options, ascs); err != nil {
		return nil, err
	}
	return ascs, nil
}

// FindAccountSaleClosingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSaleClosingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountSaleClosingModel, criteria, options)
}

// FindAccountSaleClosingId finds record id by querying it with criteria.
func (c *Client) FindAccountSaleClosingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountSaleClosingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
