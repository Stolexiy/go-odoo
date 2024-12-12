package odoo

// LoyaltyHistory represents loyalty.history model.
type LoyaltyHistory struct {
	CardId      *Many2One `xmlrpc:"card_id,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Issued      *Float    `xmlrpc:"issued,omitempty"`
	OrderId     *Many2One `xmlrpc:"order_id,omitempty"`
	OrderModel  *String   `xmlrpc:"order_model,omitempty"`
	Used        *Float    `xmlrpc:"used,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyHistorys represents array of loyalty.history model.
type LoyaltyHistorys []LoyaltyHistory

// LoyaltyHistoryModel is the odoo model name.
const LoyaltyHistoryModel = "loyalty.history"

// Many2One convert LoyaltyHistory to *Many2One.
func (lh *LoyaltyHistory) Many2One() *Many2One {
	return NewMany2One(lh.Id.Get(), "")
}

// CreateLoyaltyHistory creates a new loyalty.history model and returns its id.
func (c *Client) CreateLoyaltyHistory(lh *LoyaltyHistory) (int64, error) {
	ids, err := c.CreateLoyaltyHistorys([]*LoyaltyHistory{lh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyHistory creates a new loyalty.history model and returns its id.
func (c *Client) CreateLoyaltyHistorys(lhs []*LoyaltyHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range lhs {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyHistoryModel, vv, nil)
}

// UpdateLoyaltyHistory updates an existing loyalty.history record.
func (c *Client) UpdateLoyaltyHistory(lh *LoyaltyHistory) error {
	return c.UpdateLoyaltyHistorys([]int64{lh.Id.Get()}, lh)
}

// UpdateLoyaltyHistorys updates existing loyalty.history records.
// All records (represented by ids) will be updated by lh values.
func (c *Client) UpdateLoyaltyHistorys(ids []int64, lh *LoyaltyHistory) error {
	return c.Update(LoyaltyHistoryModel, ids, lh, nil)
}

// DeleteLoyaltyHistory deletes an existing loyalty.history record.
func (c *Client) DeleteLoyaltyHistory(id int64) error {
	return c.DeleteLoyaltyHistorys([]int64{id})
}

// DeleteLoyaltyHistorys deletes existing loyalty.history records.
func (c *Client) DeleteLoyaltyHistorys(ids []int64) error {
	return c.Delete(LoyaltyHistoryModel, ids)
}

// GetLoyaltyHistory gets loyalty.history existing record.
func (c *Client) GetLoyaltyHistory(id int64) (*LoyaltyHistory, error) {
	lhs, err := c.GetLoyaltyHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lhs)[0]), nil
}

// GetLoyaltyHistorys gets loyalty.history existing records.
func (c *Client) GetLoyaltyHistorys(ids []int64) (*LoyaltyHistorys, error) {
	lhs := &LoyaltyHistorys{}
	if err := c.Read(LoyaltyHistoryModel, ids, nil, lhs); err != nil {
		return nil, err
	}
	return lhs, nil
}

// FindLoyaltyHistory finds loyalty.history record by querying it with criteria.
func (c *Client) FindLoyaltyHistory(criteria *Criteria) (*LoyaltyHistory, error) {
	lhs := &LoyaltyHistorys{}
	if err := c.SearchRead(LoyaltyHistoryModel, criteria, NewOptions().Limit(1), lhs); err != nil {
		return nil, err
	}
	return &((*lhs)[0]), nil
}

// FindLoyaltyHistorys finds loyalty.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyHistorys(criteria *Criteria, options *Options) (*LoyaltyHistorys, error) {
	lhs := &LoyaltyHistorys{}
	if err := c.SearchRead(LoyaltyHistoryModel, criteria, options, lhs); err != nil {
		return nil, err
	}
	return lhs, nil
}

// FindLoyaltyHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyHistoryModel, criteria, options)
}

// FindLoyaltyHistoryId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
