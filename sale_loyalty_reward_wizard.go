package odoo

// SaleLoyaltyRewardWizard represents sale.loyalty.reward.wizard model.
type SaleLoyaltyRewardWizard struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	MultiProductReward *Bool     `xmlrpc:"multi_product_reward,omitempty"`
	OrderId            *Many2One `xmlrpc:"order_id,omitempty"`
	RewardIds          *Relation `xmlrpc:"reward_ids,omitempty"`
	RewardProductIds   *Relation `xmlrpc:"reward_product_ids,omitempty"`
	SelectedProductId  *Many2One `xmlrpc:"selected_product_id,omitempty"`
	SelectedRewardId   *Many2One `xmlrpc:"selected_reward_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleLoyaltyRewardWizards represents array of sale.loyalty.reward.wizard model.
type SaleLoyaltyRewardWizards []SaleLoyaltyRewardWizard

// SaleLoyaltyRewardWizardModel is the odoo model name.
const SaleLoyaltyRewardWizardModel = "sale.loyalty.reward.wizard"

// Many2One convert SaleLoyaltyRewardWizard to *Many2One.
func (slrw *SaleLoyaltyRewardWizard) Many2One() *Many2One {
	return NewMany2One(slrw.Id.Get(), "")
}

// CreateSaleLoyaltyRewardWizard creates a new sale.loyalty.reward.wizard model and returns its id.
func (c *Client) CreateSaleLoyaltyRewardWizard(slrw *SaleLoyaltyRewardWizard) (int64, error) {
	ids, err := c.CreateSaleLoyaltyRewardWizards([]*SaleLoyaltyRewardWizard{slrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleLoyaltyRewardWizard creates a new sale.loyalty.reward.wizard model and returns its id.
func (c *Client) CreateSaleLoyaltyRewardWizards(slrws []*SaleLoyaltyRewardWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range slrws {
		vv = append(vv, v)
	}
	return c.Create(SaleLoyaltyRewardWizardModel, vv, nil)
}

// UpdateSaleLoyaltyRewardWizard updates an existing sale.loyalty.reward.wizard record.
func (c *Client) UpdateSaleLoyaltyRewardWizard(slrw *SaleLoyaltyRewardWizard) error {
	return c.UpdateSaleLoyaltyRewardWizards([]int64{slrw.Id.Get()}, slrw)
}

// UpdateSaleLoyaltyRewardWizards updates existing sale.loyalty.reward.wizard records.
// All records (represented by ids) will be updated by slrw values.
func (c *Client) UpdateSaleLoyaltyRewardWizards(ids []int64, slrw *SaleLoyaltyRewardWizard) error {
	return c.Update(SaleLoyaltyRewardWizardModel, ids, slrw, nil)
}

// DeleteSaleLoyaltyRewardWizard deletes an existing sale.loyalty.reward.wizard record.
func (c *Client) DeleteSaleLoyaltyRewardWizard(id int64) error {
	return c.DeleteSaleLoyaltyRewardWizards([]int64{id})
}

// DeleteSaleLoyaltyRewardWizards deletes existing sale.loyalty.reward.wizard records.
func (c *Client) DeleteSaleLoyaltyRewardWizards(ids []int64) error {
	return c.Delete(SaleLoyaltyRewardWizardModel, ids)
}

// GetSaleLoyaltyRewardWizard gets sale.loyalty.reward.wizard existing record.
func (c *Client) GetSaleLoyaltyRewardWizard(id int64) (*SaleLoyaltyRewardWizard, error) {
	slrws, err := c.GetSaleLoyaltyRewardWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slrws)[0]), nil
}

// GetSaleLoyaltyRewardWizards gets sale.loyalty.reward.wizard existing records.
func (c *Client) GetSaleLoyaltyRewardWizards(ids []int64) (*SaleLoyaltyRewardWizards, error) {
	slrws := &SaleLoyaltyRewardWizards{}
	if err := c.Read(SaleLoyaltyRewardWizardModel, ids, nil, slrws); err != nil {
		return nil, err
	}
	return slrws, nil
}

// FindSaleLoyaltyRewardWizard finds sale.loyalty.reward.wizard record by querying it with criteria.
func (c *Client) FindSaleLoyaltyRewardWizard(criteria *Criteria) (*SaleLoyaltyRewardWizard, error) {
	slrws := &SaleLoyaltyRewardWizards{}
	if err := c.SearchRead(SaleLoyaltyRewardWizardModel, criteria, NewOptions().Limit(1), slrws); err != nil {
		return nil, err
	}
	return &((*slrws)[0]), nil
}

// FindSaleLoyaltyRewardWizards finds sale.loyalty.reward.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLoyaltyRewardWizards(criteria *Criteria, options *Options) (*SaleLoyaltyRewardWizards, error) {
	slrws := &SaleLoyaltyRewardWizards{}
	if err := c.SearchRead(SaleLoyaltyRewardWizardModel, criteria, options, slrws); err != nil {
		return nil, err
	}
	return slrws, nil
}

// FindSaleLoyaltyRewardWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleLoyaltyRewardWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleLoyaltyRewardWizardModel, criteria, options)
}

// FindSaleLoyaltyRewardWizardId finds record id by querying it with criteria.
func (c *Client) FindSaleLoyaltyRewardWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleLoyaltyRewardWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
