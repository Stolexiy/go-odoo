package odoo

// BudgetReport represents budget.report model.
type BudgetReport struct {
	AccountId        *Many2One  `xmlrpc:"account_id,omitempty"`
	Achieved         *Float     `xmlrpc:"achieved,omitempty"`
	AutoAccountId    *Many2One  `xmlrpc:"auto_account_id,omitempty"`
	Budget           *Float     `xmlrpc:"budget,omitempty"`
	BudgetAnalyticId *Many2One  `xmlrpc:"budget_analytic_id,omitempty"`
	BudgetLineId     *Many2One  `xmlrpc:"budget_line_id,omitempty"`
	Committed        *Float     `xmlrpc:"committed,omitempty"`
	CompanyId        *Many2One  `xmlrpc:"company_id,omitempty"`
	Date             *Time      `xmlrpc:"date,omitempty"`
	Description      *String    `xmlrpc:"description,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	LineType         *Selection `xmlrpc:"line_type,omitempty"`
	ResId            *Many2One  `xmlrpc:"res_id,omitempty"`
	ResModel         *String    `xmlrpc:"res_model,omitempty"`
	UserId           *Many2One  `xmlrpc:"user_id,omitempty"`
	XPlan155Id       *Many2One  `xmlrpc:"x_plan155_id,omitempty"`
	XPlan159Id       *Many2One  `xmlrpc:"x_plan159_id,omitempty"`
	XPlan161Id       *Many2One  `xmlrpc:"x_plan161_id,omitempty"`
	XPlan16Id        *Many2One  `xmlrpc:"x_plan16_id,omitempty"`
	XPlan17Id        *Many2One  `xmlrpc:"x_plan17_id,omitempty"`
	XPlan18Id        *Many2One  `xmlrpc:"x_plan18_id,omitempty"`
	XPlan19Id        *Many2One  `xmlrpc:"x_plan19_id,omitempty"`
	XPlan24Id        *Many2One  `xmlrpc:"x_plan24_id,omitempty"`
	XPlan26Id        *Many2One  `xmlrpc:"x_plan26_id,omitempty"`
}

// BudgetReports represents array of budget.report model.
type BudgetReports []BudgetReport

// BudgetReportModel is the odoo model name.
const BudgetReportModel = "budget.report"

// Many2One convert BudgetReport to *Many2One.
func (br *BudgetReport) Many2One() *Many2One {
	return NewMany2One(br.Id.Get(), "")
}

// CreateBudgetReport creates a new budget.report model and returns its id.
func (c *Client) CreateBudgetReport(br *BudgetReport) (int64, error) {
	ids, err := c.CreateBudgetReports([]*BudgetReport{br})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBudgetReport creates a new budget.report model and returns its id.
func (c *Client) CreateBudgetReports(brs []*BudgetReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range brs {
		vv = append(vv, v)
	}
	return c.Create(BudgetReportModel, vv, nil)
}

// UpdateBudgetReport updates an existing budget.report record.
func (c *Client) UpdateBudgetReport(br *BudgetReport) error {
	return c.UpdateBudgetReports([]int64{br.Id.Get()}, br)
}

// UpdateBudgetReports updates existing budget.report records.
// All records (represented by ids) will be updated by br values.
func (c *Client) UpdateBudgetReports(ids []int64, br *BudgetReport) error {
	return c.Update(BudgetReportModel, ids, br, nil)
}

// DeleteBudgetReport deletes an existing budget.report record.
func (c *Client) DeleteBudgetReport(id int64) error {
	return c.DeleteBudgetReports([]int64{id})
}

// DeleteBudgetReports deletes existing budget.report records.
func (c *Client) DeleteBudgetReports(ids []int64) error {
	return c.Delete(BudgetReportModel, ids)
}

// GetBudgetReport gets budget.report existing record.
func (c *Client) GetBudgetReport(id int64) (*BudgetReport, error) {
	brs, err := c.GetBudgetReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*brs)[0]), nil
}

// GetBudgetReports gets budget.report existing records.
func (c *Client) GetBudgetReports(ids []int64) (*BudgetReports, error) {
	brs := &BudgetReports{}
	if err := c.Read(BudgetReportModel, ids, nil, brs); err != nil {
		return nil, err
	}
	return brs, nil
}

// FindBudgetReport finds budget.report record by querying it with criteria.
func (c *Client) FindBudgetReport(criteria *Criteria) (*BudgetReport, error) {
	brs := &BudgetReports{}
	if err := c.SearchRead(BudgetReportModel, criteria, NewOptions().Limit(1), brs); err != nil {
		return nil, err
	}
	return &((*brs)[0]), nil
}

// FindBudgetReports finds budget.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetReports(criteria *Criteria, options *Options) (*BudgetReports, error) {
	brs := &BudgetReports{}
	if err := c.SearchRead(BudgetReportModel, criteria, options, brs); err != nil {
		return nil, err
	}
	return brs, nil
}

// FindBudgetReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BudgetReportModel, criteria, options)
}

// FindBudgetReportId finds record id by querying it with criteria.
func (c *Client) FindBudgetReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BudgetReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
