package odoo

// BudgetLine represents budget.line model.
type BudgetLine struct {
	AccountId             *Many2One  `xmlrpc:"account_id,omitempty"`
	AchievedAmount        *Float     `xmlrpc:"achieved_amount,omitempty"`
	AchievedPercentage    *Float     `xmlrpc:"achieved_percentage,omitempty"`
	AutoAccountId         *Many2One  `xmlrpc:"auto_account_id,omitempty"`
	BudgetAmount          *Float     `xmlrpc:"budget_amount,omitempty"`
	BudgetAnalyticId      *Many2One  `xmlrpc:"budget_analytic_id,omitempty"`
	BudgetAnalyticState   *Selection `xmlrpc:"budget_analytic_state,omitempty"`
	CommittedAmount       *Float     `xmlrpc:"committed_amount,omitempty"`
	CommittedPercentage   *Float     `xmlrpc:"committed_percentage,omitempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId            *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateFrom              *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	IsAboveBudget         *Bool      `xmlrpc:"is_above_budget,omitempty"`
	Name                  *String    `xmlrpc:"name,omitempty"`
	Sequence              *Int       `xmlrpc:"sequence,omitempty"`
	TheoriticalAmount     *Float     `xmlrpc:"theoritical_amount,omitempty"`
	TheoriticalPercentage *Float     `xmlrpc:"theoritical_percentage,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
	XPlan155Id            *Many2One  `xmlrpc:"x_plan155_id,omitempty"`
	XPlan159Id            *Many2One  `xmlrpc:"x_plan159_id,omitempty"`
	XPlan161Id            *Many2One  `xmlrpc:"x_plan161_id,omitempty"`
	XPlan16Id             *Many2One  `xmlrpc:"x_plan16_id,omitempty"`
	XPlan17Id             *Many2One  `xmlrpc:"x_plan17_id,omitempty"`
	XPlan18Id             *Many2One  `xmlrpc:"x_plan18_id,omitempty"`
	XPlan19Id             *Many2One  `xmlrpc:"x_plan19_id,omitempty"`
	XPlan24Id             *Many2One  `xmlrpc:"x_plan24_id,omitempty"`
	XPlan26Id             *Many2One  `xmlrpc:"x_plan26_id,omitempty"`
}

// BudgetLines represents array of budget.line model.
type BudgetLines []BudgetLine

// BudgetLineModel is the odoo model name.
const BudgetLineModel = "budget.line"

// Many2One convert BudgetLine to *Many2One.
func (bl *BudgetLine) Many2One() *Many2One {
	return NewMany2One(bl.Id.Get(), "")
}

// CreateBudgetLine creates a new budget.line model and returns its id.
func (c *Client) CreateBudgetLine(bl *BudgetLine) (int64, error) {
	ids, err := c.CreateBudgetLines([]*BudgetLine{bl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBudgetLine creates a new budget.line model and returns its id.
func (c *Client) CreateBudgetLines(bls []*BudgetLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range bls {
		vv = append(vv, v)
	}
	return c.Create(BudgetLineModel, vv, nil)
}

// UpdateBudgetLine updates an existing budget.line record.
func (c *Client) UpdateBudgetLine(bl *BudgetLine) error {
	return c.UpdateBudgetLines([]int64{bl.Id.Get()}, bl)
}

// UpdateBudgetLines updates existing budget.line records.
// All records (represented by ids) will be updated by bl values.
func (c *Client) UpdateBudgetLines(ids []int64, bl *BudgetLine) error {
	return c.Update(BudgetLineModel, ids, bl, nil)
}

// DeleteBudgetLine deletes an existing budget.line record.
func (c *Client) DeleteBudgetLine(id int64) error {
	return c.DeleteBudgetLines([]int64{id})
}

// DeleteBudgetLines deletes existing budget.line records.
func (c *Client) DeleteBudgetLines(ids []int64) error {
	return c.Delete(BudgetLineModel, ids)
}

// GetBudgetLine gets budget.line existing record.
func (c *Client) GetBudgetLine(id int64) (*BudgetLine, error) {
	bls, err := c.GetBudgetLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bls)[0]), nil
}

// GetBudgetLines gets budget.line existing records.
func (c *Client) GetBudgetLines(ids []int64) (*BudgetLines, error) {
	bls := &BudgetLines{}
	if err := c.Read(BudgetLineModel, ids, nil, bls); err != nil {
		return nil, err
	}
	return bls, nil
}

// FindBudgetLine finds budget.line record by querying it with criteria.
func (c *Client) FindBudgetLine(criteria *Criteria) (*BudgetLine, error) {
	bls := &BudgetLines{}
	if err := c.SearchRead(BudgetLineModel, criteria, NewOptions().Limit(1), bls); err != nil {
		return nil, err
	}
	return &((*bls)[0]), nil
}

// FindBudgetLines finds budget.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetLines(criteria *Criteria, options *Options) (*BudgetLines, error) {
	bls := &BudgetLines{}
	if err := c.SearchRead(BudgetLineModel, criteria, options, bls); err != nil {
		return nil, err
	}
	return bls, nil
}

// FindBudgetLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BudgetLineModel, criteria, options)
}

// FindBudgetLineId finds record id by querying it with criteria.
func (c *Client) FindBudgetLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BudgetLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
