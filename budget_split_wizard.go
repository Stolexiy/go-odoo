package odoo

// BudgetSplitWizard represents budget.split.wizard model.
type BudgetSplitWizard struct {
	AnalyticalPlanIds *Relation  `xmlrpc:"analytical_plan_ids,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom          *Time      `xmlrpc:"date_from,omitempty"`
	DateTo            *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	Period            *Selection `xmlrpc:"period,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BudgetSplitWizards represents array of budget.split.wizard model.
type BudgetSplitWizards []BudgetSplitWizard

// BudgetSplitWizardModel is the odoo model name.
const BudgetSplitWizardModel = "budget.split.wizard"

// Many2One convert BudgetSplitWizard to *Many2One.
func (bsw *BudgetSplitWizard) Many2One() *Many2One {
	return NewMany2One(bsw.Id.Get(), "")
}

// CreateBudgetSplitWizard creates a new budget.split.wizard model and returns its id.
func (c *Client) CreateBudgetSplitWizard(bsw *BudgetSplitWizard) (int64, error) {
	ids, err := c.CreateBudgetSplitWizards([]*BudgetSplitWizard{bsw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBudgetSplitWizard creates a new budget.split.wizard model and returns its id.
func (c *Client) CreateBudgetSplitWizards(bsws []*BudgetSplitWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range bsws {
		vv = append(vv, v)
	}
	return c.Create(BudgetSplitWizardModel, vv, nil)
}

// UpdateBudgetSplitWizard updates an existing budget.split.wizard record.
func (c *Client) UpdateBudgetSplitWizard(bsw *BudgetSplitWizard) error {
	return c.UpdateBudgetSplitWizards([]int64{bsw.Id.Get()}, bsw)
}

// UpdateBudgetSplitWizards updates existing budget.split.wizard records.
// All records (represented by ids) will be updated by bsw values.
func (c *Client) UpdateBudgetSplitWizards(ids []int64, bsw *BudgetSplitWizard) error {
	return c.Update(BudgetSplitWizardModel, ids, bsw, nil)
}

// DeleteBudgetSplitWizard deletes an existing budget.split.wizard record.
func (c *Client) DeleteBudgetSplitWizard(id int64) error {
	return c.DeleteBudgetSplitWizards([]int64{id})
}

// DeleteBudgetSplitWizards deletes existing budget.split.wizard records.
func (c *Client) DeleteBudgetSplitWizards(ids []int64) error {
	return c.Delete(BudgetSplitWizardModel, ids)
}

// GetBudgetSplitWizard gets budget.split.wizard existing record.
func (c *Client) GetBudgetSplitWizard(id int64) (*BudgetSplitWizard, error) {
	bsws, err := c.GetBudgetSplitWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bsws)[0]), nil
}

// GetBudgetSplitWizards gets budget.split.wizard existing records.
func (c *Client) GetBudgetSplitWizards(ids []int64) (*BudgetSplitWizards, error) {
	bsws := &BudgetSplitWizards{}
	if err := c.Read(BudgetSplitWizardModel, ids, nil, bsws); err != nil {
		return nil, err
	}
	return bsws, nil
}

// FindBudgetSplitWizard finds budget.split.wizard record by querying it with criteria.
func (c *Client) FindBudgetSplitWizard(criteria *Criteria) (*BudgetSplitWizard, error) {
	bsws := &BudgetSplitWizards{}
	if err := c.SearchRead(BudgetSplitWizardModel, criteria, NewOptions().Limit(1), bsws); err != nil {
		return nil, err
	}
	return &((*bsws)[0]), nil
}

// FindBudgetSplitWizards finds budget.split.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetSplitWizards(criteria *Criteria, options *Options) (*BudgetSplitWizards, error) {
	bsws := &BudgetSplitWizards{}
	if err := c.SearchRead(BudgetSplitWizardModel, criteria, options, bsws); err != nil {
		return nil, err
	}
	return bsws, nil
}

// FindBudgetSplitWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetSplitWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BudgetSplitWizardModel, criteria, options)
}

// FindBudgetSplitWizardId finds record id by querying it with criteria.
func (c *Client) FindBudgetSplitWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BudgetSplitWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
