package odoo

// LoyaltyCardUpdateBalance represents loyalty.card.update.balance model.
type LoyaltyCardUpdateBalance struct {
	CardId      *Many2One `xmlrpc:"card_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	NewBalance  *Float    `xmlrpc:"new_balance,omitempty"`
	OldBalance  *Float    `xmlrpc:"old_balance,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyCardUpdateBalances represents array of loyalty.card.update.balance model.
type LoyaltyCardUpdateBalances []LoyaltyCardUpdateBalance

// LoyaltyCardUpdateBalanceModel is the odoo model name.
const LoyaltyCardUpdateBalanceModel = "loyalty.card.update.balance"

// Many2One convert LoyaltyCardUpdateBalance to *Many2One.
func (lcub *LoyaltyCardUpdateBalance) Many2One() *Many2One {
	return NewMany2One(lcub.Id.Get(), "")
}

// CreateLoyaltyCardUpdateBalance creates a new loyalty.card.update.balance model and returns its id.
func (c *Client) CreateLoyaltyCardUpdateBalance(lcub *LoyaltyCardUpdateBalance) (int64, error) {
	ids, err := c.CreateLoyaltyCardUpdateBalances([]*LoyaltyCardUpdateBalance{lcub})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyCardUpdateBalance creates a new loyalty.card.update.balance model and returns its id.
func (c *Client) CreateLoyaltyCardUpdateBalances(lcubs []*LoyaltyCardUpdateBalance) ([]int64, error) {
	var vv []interface{}
	for _, v := range lcubs {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyCardUpdateBalanceModel, vv, nil)
}

// UpdateLoyaltyCardUpdateBalance updates an existing loyalty.card.update.balance record.
func (c *Client) UpdateLoyaltyCardUpdateBalance(lcub *LoyaltyCardUpdateBalance) error {
	return c.UpdateLoyaltyCardUpdateBalances([]int64{lcub.Id.Get()}, lcub)
}

// UpdateLoyaltyCardUpdateBalances updates existing loyalty.card.update.balance records.
// All records (represented by ids) will be updated by lcub values.
func (c *Client) UpdateLoyaltyCardUpdateBalances(ids []int64, lcub *LoyaltyCardUpdateBalance) error {
	return c.Update(LoyaltyCardUpdateBalanceModel, ids, lcub, nil)
}

// DeleteLoyaltyCardUpdateBalance deletes an existing loyalty.card.update.balance record.
func (c *Client) DeleteLoyaltyCardUpdateBalance(id int64) error {
	return c.DeleteLoyaltyCardUpdateBalances([]int64{id})
}

// DeleteLoyaltyCardUpdateBalances deletes existing loyalty.card.update.balance records.
func (c *Client) DeleteLoyaltyCardUpdateBalances(ids []int64) error {
	return c.Delete(LoyaltyCardUpdateBalanceModel, ids)
}

// GetLoyaltyCardUpdateBalance gets loyalty.card.update.balance existing record.
func (c *Client) GetLoyaltyCardUpdateBalance(id int64) (*LoyaltyCardUpdateBalance, error) {
	lcubs, err := c.GetLoyaltyCardUpdateBalances([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lcubs)[0]), nil
}

// GetLoyaltyCardUpdateBalances gets loyalty.card.update.balance existing records.
func (c *Client) GetLoyaltyCardUpdateBalances(ids []int64) (*LoyaltyCardUpdateBalances, error) {
	lcubs := &LoyaltyCardUpdateBalances{}
	if err := c.Read(LoyaltyCardUpdateBalanceModel, ids, nil, lcubs); err != nil {
		return nil, err
	}
	return lcubs, nil
}

// FindLoyaltyCardUpdateBalance finds loyalty.card.update.balance record by querying it with criteria.
func (c *Client) FindLoyaltyCardUpdateBalance(criteria *Criteria) (*LoyaltyCardUpdateBalance, error) {
	lcubs := &LoyaltyCardUpdateBalances{}
	if err := c.SearchRead(LoyaltyCardUpdateBalanceModel, criteria, NewOptions().Limit(1), lcubs); err != nil {
		return nil, err
	}
	return &((*lcubs)[0]), nil
}

// FindLoyaltyCardUpdateBalances finds loyalty.card.update.balance records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyCardUpdateBalances(criteria *Criteria, options *Options) (*LoyaltyCardUpdateBalances, error) {
	lcubs := &LoyaltyCardUpdateBalances{}
	if err := c.SearchRead(LoyaltyCardUpdateBalanceModel, criteria, options, lcubs); err != nil {
		return nil, err
	}
	return lcubs, nil
}

// FindLoyaltyCardUpdateBalanceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyCardUpdateBalanceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyCardUpdateBalanceModel, criteria, options)
}

// FindLoyaltyCardUpdateBalanceId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyCardUpdateBalanceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyCardUpdateBalanceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
