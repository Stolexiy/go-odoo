package odoo

// LoyaltyReward represents loyalty.reward model.
type LoyaltyReward struct {
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	AllDiscountProductIds     *Relation  `xmlrpc:"all_discount_product_ids,omitempty"`
	ClearWallet               *Bool      `xmlrpc:"clear_wallet,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty"`
	Description               *String    `xmlrpc:"description,omitempty"`
	Discount                  *Float     `xmlrpc:"discount,omitempty"`
	DiscountApplicability     *Selection `xmlrpc:"discount_applicability,omitempty"`
	DiscountLineProductId     *Many2One  `xmlrpc:"discount_line_product_id,omitempty"`
	DiscountMaxAmount         *Float     `xmlrpc:"discount_max_amount,omitempty"`
	DiscountMode              *Selection `xmlrpc:"discount_mode,omitempty"`
	DiscountProductCategoryId *Many2One  `xmlrpc:"discount_product_category_id,omitempty"`
	DiscountProductDomain     *String    `xmlrpc:"discount_product_domain,omitempty"`
	DiscountProductIds        *Relation  `xmlrpc:"discount_product_ids,omitempty"`
	DiscountProductTagId      *Many2One  `xmlrpc:"discount_product_tag_id,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	IsGlobalDiscount          *Bool      `xmlrpc:"is_global_discount,omitempty"`
	MultiProduct              *Bool      `xmlrpc:"multi_product,omitempty"`
	PointName                 *String    `xmlrpc:"point_name,omitempty"`
	ProgramId                 *Many2One  `xmlrpc:"program_id,omitempty"`
	ProgramType               *Selection `xmlrpc:"program_type,omitempty"`
	RequiredPoints            *Float     `xmlrpc:"required_points,omitempty"`
	RewardProductDomain       *String    `xmlrpc:"reward_product_domain,omitempty"`
	RewardProductId           *Many2One  `xmlrpc:"reward_product_id,omitempty"`
	RewardProductIds          *Relation  `xmlrpc:"reward_product_ids,omitempty"`
	RewardProductQty          *Int       `xmlrpc:"reward_product_qty,omitempty"`
	RewardProductTagId        *Many2One  `xmlrpc:"reward_product_tag_id,omitempty"`
	RewardProductUomId        *Many2One  `xmlrpc:"reward_product_uom_id,omitempty"`
	RewardType                *Selection `xmlrpc:"reward_type,omitempty"`
	TaxIds                    *Relation  `xmlrpc:"tax_ids,omitempty"`
	UserHasDebug              *Bool      `xmlrpc:"user_has_debug,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyRewards represents array of loyalty.reward model.
type LoyaltyRewards []LoyaltyReward

// LoyaltyRewardModel is the odoo model name.
const LoyaltyRewardModel = "loyalty.reward"

// Many2One convert LoyaltyReward to *Many2One.
func (lr *LoyaltyReward) Many2One() *Many2One {
	return NewMany2One(lr.Id.Get(), "")
}

// CreateLoyaltyReward creates a new loyalty.reward model and returns its id.
func (c *Client) CreateLoyaltyReward(lr *LoyaltyReward) (int64, error) {
	ids, err := c.CreateLoyaltyRewards([]*LoyaltyReward{lr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyReward creates a new loyalty.reward model and returns its id.
func (c *Client) CreateLoyaltyRewards(lrs []*LoyaltyReward) ([]int64, error) {
	var vv []interface{}
	for _, v := range lrs {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyRewardModel, vv, nil)
}

// UpdateLoyaltyReward updates an existing loyalty.reward record.
func (c *Client) UpdateLoyaltyReward(lr *LoyaltyReward) error {
	return c.UpdateLoyaltyRewards([]int64{lr.Id.Get()}, lr)
}

// UpdateLoyaltyRewards updates existing loyalty.reward records.
// All records (represented by ids) will be updated by lr values.
func (c *Client) UpdateLoyaltyRewards(ids []int64, lr *LoyaltyReward) error {
	return c.Update(LoyaltyRewardModel, ids, lr, nil)
}

// DeleteLoyaltyReward deletes an existing loyalty.reward record.
func (c *Client) DeleteLoyaltyReward(id int64) error {
	return c.DeleteLoyaltyRewards([]int64{id})
}

// DeleteLoyaltyRewards deletes existing loyalty.reward records.
func (c *Client) DeleteLoyaltyRewards(ids []int64) error {
	return c.Delete(LoyaltyRewardModel, ids)
}

// GetLoyaltyReward gets loyalty.reward existing record.
func (c *Client) GetLoyaltyReward(id int64) (*LoyaltyReward, error) {
	lrs, err := c.GetLoyaltyRewards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lrs)[0]), nil
}

// GetLoyaltyRewards gets loyalty.reward existing records.
func (c *Client) GetLoyaltyRewards(ids []int64) (*LoyaltyRewards, error) {
	lrs := &LoyaltyRewards{}
	if err := c.Read(LoyaltyRewardModel, ids, nil, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyReward finds loyalty.reward record by querying it with criteria.
func (c *Client) FindLoyaltyReward(criteria *Criteria) (*LoyaltyReward, error) {
	lrs := &LoyaltyRewards{}
	if err := c.SearchRead(LoyaltyRewardModel, criteria, NewOptions().Limit(1), lrs); err != nil {
		return nil, err
	}
	return &((*lrs)[0]), nil
}

// FindLoyaltyRewards finds loyalty.reward records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRewards(criteria *Criteria, options *Options) (*LoyaltyRewards, error) {
	lrs := &LoyaltyRewards{}
	if err := c.SearchRead(LoyaltyRewardModel, criteria, options, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyRewardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRewardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyRewardModel, criteria, options)
}

// FindLoyaltyRewardId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyRewardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyRewardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
