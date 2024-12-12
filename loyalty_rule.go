package odoo

// LoyaltyRule represents loyalty.rule model.
type LoyaltyRule struct {
	Active               *Bool      `xmlrpc:"active,omitempty"`
	AnyProduct           *Bool      `xmlrpc:"any_product,omitempty"`
	Code                 *String    `xmlrpc:"code,omitempty"`
	CompanyId            *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId           *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	MinimumAmount        *Float     `xmlrpc:"minimum_amount,omitempty"`
	MinimumAmountTaxMode *Selection `xmlrpc:"minimum_amount_tax_mode,omitempty"`
	MinimumQty           *Int       `xmlrpc:"minimum_qty,omitempty"`
	Mode                 *Selection `xmlrpc:"mode,omitempty"`
	ProductCategoryId    *Many2One  `xmlrpc:"product_category_id,omitempty"`
	ProductDomain        *String    `xmlrpc:"product_domain,omitempty"`
	ProductIds           *Relation  `xmlrpc:"product_ids,omitempty"`
	ProductTagId         *Many2One  `xmlrpc:"product_tag_id,omitempty"`
	ProgramId            *Many2One  `xmlrpc:"program_id,omitempty"`
	ProgramType          *Selection `xmlrpc:"program_type,omitempty"`
	PromoBarcode         *String    `xmlrpc:"promo_barcode,omitempty"`
	RewardPointAmount    *Float     `xmlrpc:"reward_point_amount,omitempty"`
	RewardPointMode      *Selection `xmlrpc:"reward_point_mode,omitempty"`
	RewardPointName      *String    `xmlrpc:"reward_point_name,omitempty"`
	RewardPointSplit     *Bool      `xmlrpc:"reward_point_split,omitempty"`
	UserHasDebug         *Bool      `xmlrpc:"user_has_debug,omitempty"`
	ValidProductIds      *Relation  `xmlrpc:"valid_product_ids,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyRules represents array of loyalty.rule model.
type LoyaltyRules []LoyaltyRule

// LoyaltyRuleModel is the odoo model name.
const LoyaltyRuleModel = "loyalty.rule"

// Many2One convert LoyaltyRule to *Many2One.
func (lr *LoyaltyRule) Many2One() *Many2One {
	return NewMany2One(lr.Id.Get(), "")
}

// CreateLoyaltyRule creates a new loyalty.rule model and returns its id.
func (c *Client) CreateLoyaltyRule(lr *LoyaltyRule) (int64, error) {
	ids, err := c.CreateLoyaltyRules([]*LoyaltyRule{lr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyRule creates a new loyalty.rule model and returns its id.
func (c *Client) CreateLoyaltyRules(lrs []*LoyaltyRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range lrs {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyRuleModel, vv, nil)
}

// UpdateLoyaltyRule updates an existing loyalty.rule record.
func (c *Client) UpdateLoyaltyRule(lr *LoyaltyRule) error {
	return c.UpdateLoyaltyRules([]int64{lr.Id.Get()}, lr)
}

// UpdateLoyaltyRules updates existing loyalty.rule records.
// All records (represented by ids) will be updated by lr values.
func (c *Client) UpdateLoyaltyRules(ids []int64, lr *LoyaltyRule) error {
	return c.Update(LoyaltyRuleModel, ids, lr, nil)
}

// DeleteLoyaltyRule deletes an existing loyalty.rule record.
func (c *Client) DeleteLoyaltyRule(id int64) error {
	return c.DeleteLoyaltyRules([]int64{id})
}

// DeleteLoyaltyRules deletes existing loyalty.rule records.
func (c *Client) DeleteLoyaltyRules(ids []int64) error {
	return c.Delete(LoyaltyRuleModel, ids)
}

// GetLoyaltyRule gets loyalty.rule existing record.
func (c *Client) GetLoyaltyRule(id int64) (*LoyaltyRule, error) {
	lrs, err := c.GetLoyaltyRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lrs)[0]), nil
}

// GetLoyaltyRules gets loyalty.rule existing records.
func (c *Client) GetLoyaltyRules(ids []int64) (*LoyaltyRules, error) {
	lrs := &LoyaltyRules{}
	if err := c.Read(LoyaltyRuleModel, ids, nil, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyRule finds loyalty.rule record by querying it with criteria.
func (c *Client) FindLoyaltyRule(criteria *Criteria) (*LoyaltyRule, error) {
	lrs := &LoyaltyRules{}
	if err := c.SearchRead(LoyaltyRuleModel, criteria, NewOptions().Limit(1), lrs); err != nil {
		return nil, err
	}
	return &((*lrs)[0]), nil
}

// FindLoyaltyRules finds loyalty.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRules(criteria *Criteria, options *Options) (*LoyaltyRules, error) {
	lrs := &LoyaltyRules{}
	if err := c.SearchRead(LoyaltyRuleModel, criteria, options, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyRuleModel, criteria, options)
}

// FindLoyaltyRuleId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
