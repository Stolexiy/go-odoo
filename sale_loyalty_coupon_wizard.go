package odoo

// SaleLoyaltyCouponWizard represents sale.loyalty.coupon.wizard model.
type SaleLoyaltyCouponWizard struct {
	CouponCode  *String   `xmlrpc:"coupon_code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	OrderId     *Many2One `xmlrpc:"order_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleLoyaltyCouponWizards represents array of sale.loyalty.coupon.wizard model.
type SaleLoyaltyCouponWizards []SaleLoyaltyCouponWizard

// SaleLoyaltyCouponWizardModel is the odoo model name.
const SaleLoyaltyCouponWizardModel = "sale.loyalty.coupon.wizard"

// Many2One convert SaleLoyaltyCouponWizard to *Many2One.
func (slcw *SaleLoyaltyCouponWizard) Many2One() *Many2One {
	return NewMany2One(slcw.Id.Get(), "")
}

// CreateSaleLoyaltyCouponWizard creates a new sale.loyalty.coupon.wizard model and returns its id.
func (c *Client) CreateSaleLoyaltyCouponWizard(slcw *SaleLoyaltyCouponWizard) (int64, error) {
	ids, err := c.CreateSaleLoyaltyCouponWizards([]*SaleLoyaltyCouponWizard{slcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleLoyaltyCouponWizard creates a new sale.loyalty.coupon.wizard model and returns its id.
func (c *Client) CreateSaleLoyaltyCouponWizards(slcws []*SaleLoyaltyCouponWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range slcws {
		vv = append(vv, v)
	}
	return c.Create(SaleLoyaltyCouponWizardModel, vv, nil)
}

// UpdateSaleLoyaltyCouponWizard updates an existing sale.loyalty.coupon.wizard record.
func (c *Client) UpdateSaleLoyaltyCouponWizard(slcw *SaleLoyaltyCouponWizard) error {
	return c.UpdateSaleLoyaltyCouponWizards([]int64{slcw.Id.Get()}, slcw)
}

// UpdateSaleLoyaltyCouponWizards updates existing sale.loyalty.coupon.wizard records.
// All records (represented by ids) will be updated by slcw values.
func (c *Client) UpdateSaleLoyaltyCouponWizards(ids []int64, slcw *SaleLoyaltyCouponWizard) error {
	return c.Update(SaleLoyaltyCouponWizardModel, ids, slcw, nil)
}

// DeleteSaleLoyaltyCouponWizard deletes an existing sale.loyalty.coupon.wizard record.
func (c *Client) DeleteSaleLoyaltyCouponWizard(id int64) error {
	return c.DeleteSaleLoyaltyCouponWizards([]int64{id})
}

// DeleteSaleLoyaltyCouponWizards deletes existing sale.loyalty.coupon.wizard records.
func (c *Client) DeleteSaleLoyaltyCouponWizards(ids []int64) error {
	return c.Delete(SaleLoyaltyCouponWizardModel, ids)
}

// GetSaleLoyaltyCouponWizard gets sale.loyalty.coupon.wizard existing record.
func (c *Client) GetSaleLoyaltyCouponWizard(id int64) (*SaleLoyaltyCouponWizard, error) {
	slcws, err := c.GetSaleLoyaltyCouponWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slcws)[0]), nil
}

// GetSaleLoyaltyCouponWizards gets sale.loyalty.coupon.wizard existing records.
func (c *Client) GetSaleLoyaltyCouponWizards(ids []int64) (*SaleLoyaltyCouponWizards, error) {
	slcws := &SaleLoyaltyCouponWizards{}
	if err := c.Read(SaleLoyaltyCouponWizardModel, ids, nil, slcws); err != nil {
		return nil, err
	}
	return slcws, nil
}

// FindSaleLoyaltyCouponWizard finds sale.loyalty.coupon.wizard record by querying it with criteria.
func (c *Client) FindSaleLoyaltyCouponWizard(criteria *Criteria) (*SaleLoyaltyCouponWizard, error) {
	slcws := &SaleLoyaltyCouponWizards{}
	if err := c.SearchRead(SaleLoyaltyCouponWizardModel, criteria, NewOptions().Limit(1), slcws); err != nil {
		return nil, err
	}
	return &((*slcws)[0]), nil
}

// FindSaleLoyaltyCouponWizards finds sale.loyalty.coupon.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLoyaltyCouponWizards(criteria *Criteria, options *Options) (*SaleLoyaltyCouponWizards, error) {
	slcws := &SaleLoyaltyCouponWizards{}
	if err := c.SearchRead(SaleLoyaltyCouponWizardModel, criteria, options, slcws); err != nil {
		return nil, err
	}
	return slcws, nil
}

// FindSaleLoyaltyCouponWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLoyaltyCouponWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleLoyaltyCouponWizardModel, criteria, options)
}

// FindSaleLoyaltyCouponWizardId finds record id by querying it with criteria.
func (c *Client) FindSaleLoyaltyCouponWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleLoyaltyCouponWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
