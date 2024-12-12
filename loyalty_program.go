package odoo

// LoyaltyProgram represents loyalty.program model.
type LoyaltyProgram struct {
	Active                          *Bool      `xmlrpc:"active,omitempty"`
	AppliesOn                       *Selection `xmlrpc:"applies_on,omitempty"`
	AvailableOn                     *Bool      `xmlrpc:"available_on,omitempty"`
	CommunicationPlanIds            *Relation  `xmlrpc:"communication_plan_ids,omitempty"`
	CompanyId                       *Many2One  `xmlrpc:"company_id,omitempty"`
	CouponCount                     *Int       `xmlrpc:"coupon_count,omitempty"`
	CouponCountDisplay              *String    `xmlrpc:"coupon_count_display,omitempty"`
	CouponIds                       *Relation  `xmlrpc:"coupon_ids,omitempty"`
	CreateDate                      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                       *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                      *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencySymbol                  *String    `xmlrpc:"currency_symbol,omitempty"`
	DateFrom                        *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                          *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName                     *String    `xmlrpc:"display_name,omitempty"`
	Id                              *Int       `xmlrpc:"id,omitempty"`
	IsNominative                    *Bool      `xmlrpc:"is_nominative,omitempty"`
	IsPaymentProgram                *Bool      `xmlrpc:"is_payment_program,omitempty"`
	LimitUsage                      *Bool      `xmlrpc:"limit_usage,omitempty"`
	MailTemplateId                  *Many2One  `xmlrpc:"mail_template_id,omitempty"`
	MaxUsage                        *Int       `xmlrpc:"max_usage,omitempty"`
	Name                            *String    `xmlrpc:"name,omitempty"`
	OrderCount                      *Int       `xmlrpc:"order_count,omitempty"`
	PaymentProgramDiscountProductId *Many2One  `xmlrpc:"payment_program_discount_product_id,omitempty"`
	PortalPointName                 *String    `xmlrpc:"portal_point_name,omitempty"`
	PortalVisible                   *Bool      `xmlrpc:"portal_visible,omitempty"`
	PosConfigIds                    *Relation  `xmlrpc:"pos_config_ids,omitempty"`
	PosOk                           *Bool      `xmlrpc:"pos_ok,omitempty"`
	PosOrderCount                   *Int       `xmlrpc:"pos_order_count,omitempty"`
	PosReportPrintId                *Many2One  `xmlrpc:"pos_report_print_id,omitempty"`
	PricelistIds                    *Relation  `xmlrpc:"pricelist_ids,omitempty"`
	ProgramType                     *Selection `xmlrpc:"program_type,omitempty"`
	RewardIds                       *Relation  `xmlrpc:"reward_ids,omitempty"`
	RuleIds                         *Relation  `xmlrpc:"rule_ids,omitempty"`
	SaleOk                          *Bool      `xmlrpc:"sale_ok,omitempty"`
	Sequence                        *Int       `xmlrpc:"sequence,omitempty"`
	TotalOrderCount                 *Int       `xmlrpc:"total_order_count,omitempty"`
	Trigger                         *Selection `xmlrpc:"trigger,omitempty"`
	TriggerProductIds               *Relation  `xmlrpc:"trigger_product_ids,omitempty"`
	WriteDate                       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyPrograms represents array of loyalty.program model.
type LoyaltyPrograms []LoyaltyProgram

// LoyaltyProgramModel is the odoo model name.
const LoyaltyProgramModel = "loyalty.program"

// Many2One convert LoyaltyProgram to *Many2One.
func (lp *LoyaltyProgram) Many2One() *Many2One {
	return NewMany2One(lp.Id.Get(), "")
}

// CreateLoyaltyProgram creates a new loyalty.program model and returns its id.
func (c *Client) CreateLoyaltyProgram(lp *LoyaltyProgram) (int64, error) {
	ids, err := c.CreateLoyaltyPrograms([]*LoyaltyProgram{lp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyProgram creates a new loyalty.program model and returns its id.
func (c *Client) CreateLoyaltyPrograms(lps []*LoyaltyProgram) ([]int64, error) {
	var vv []interface{}
	for _, v := range lps {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyProgramModel, vv, nil)
}

// UpdateLoyaltyProgram updates an existing loyalty.program record.
func (c *Client) UpdateLoyaltyProgram(lp *LoyaltyProgram) error {
	return c.UpdateLoyaltyPrograms([]int64{lp.Id.Get()}, lp)
}

// UpdateLoyaltyPrograms updates existing loyalty.program records.
// All records (represented by ids) will be updated by lp values.
func (c *Client) UpdateLoyaltyPrograms(ids []int64, lp *LoyaltyProgram) error {
	return c.Update(LoyaltyProgramModel, ids, lp, nil)
}

// DeleteLoyaltyProgram deletes an existing loyalty.program record.
func (c *Client) DeleteLoyaltyProgram(id int64) error {
	return c.DeleteLoyaltyPrograms([]int64{id})
}

// DeleteLoyaltyPrograms deletes existing loyalty.program records.
func (c *Client) DeleteLoyaltyPrograms(ids []int64) error {
	return c.Delete(LoyaltyProgramModel, ids)
}

// GetLoyaltyProgram gets loyalty.program existing record.
func (c *Client) GetLoyaltyProgram(id int64) (*LoyaltyProgram, error) {
	lps, err := c.GetLoyaltyPrograms([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lps)[0]), nil
}

// GetLoyaltyPrograms gets loyalty.program existing records.
func (c *Client) GetLoyaltyPrograms(ids []int64) (*LoyaltyPrograms, error) {
	lps := &LoyaltyPrograms{}
	if err := c.Read(LoyaltyProgramModel, ids, nil, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLoyaltyProgram finds loyalty.program record by querying it with criteria.
func (c *Client) FindLoyaltyProgram(criteria *Criteria) (*LoyaltyProgram, error) {
	lps := &LoyaltyPrograms{}
	if err := c.SearchRead(LoyaltyProgramModel, criteria, NewOptions().Limit(1), lps); err != nil {
		return nil, err
	}
	return &((*lps)[0]), nil
}

// FindLoyaltyPrograms finds loyalty.program records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyPrograms(criteria *Criteria, options *Options) (*LoyaltyPrograms, error) {
	lps := &LoyaltyPrograms{}
	if err := c.SearchRead(LoyaltyProgramModel, criteria, options, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLoyaltyProgramIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyProgramIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyProgramModel, criteria, options)
}

// FindLoyaltyProgramId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyProgramId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyProgramModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
