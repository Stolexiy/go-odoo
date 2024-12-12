package odoo

// PosDailySalesReportsWizard represents pos.daily.sales.reports.wizard model.
type PosDailySalesReportsWizard struct {
	AddReportPerEmployee *Bool     `xmlrpc:"add_report_per_employee,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	EmployeeIds          *Relation `xmlrpc:"employee_ids,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	PosSessionId         *Many2One `xmlrpc:"pos_session_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosDailySalesReportsWizards represents array of pos.daily.sales.reports.wizard model.
type PosDailySalesReportsWizards []PosDailySalesReportsWizard

// PosDailySalesReportsWizardModel is the odoo model name.
const PosDailySalesReportsWizardModel = "pos.daily.sales.reports.wizard"

// Many2One convert PosDailySalesReportsWizard to *Many2One.
func (pdsrw *PosDailySalesReportsWizard) Many2One() *Many2One {
	return NewMany2One(pdsrw.Id.Get(), "")
}

// CreatePosDailySalesReportsWizard creates a new pos.daily.sales.reports.wizard model and returns its id.
func (c *Client) CreatePosDailySalesReportsWizard(pdsrw *PosDailySalesReportsWizard) (int64, error) {
	ids, err := c.CreatePosDailySalesReportsWizards([]*PosDailySalesReportsWizard{pdsrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosDailySalesReportsWizard creates a new pos.daily.sales.reports.wizard model and returns its id.
func (c *Client) CreatePosDailySalesReportsWizards(pdsrws []*PosDailySalesReportsWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pdsrws {
		vv = append(vv, v)
	}
	return c.Create(PosDailySalesReportsWizardModel, vv, nil)
}

// UpdatePosDailySalesReportsWizard updates an existing pos.daily.sales.reports.wizard record.
func (c *Client) UpdatePosDailySalesReportsWizard(pdsrw *PosDailySalesReportsWizard) error {
	return c.UpdatePosDailySalesReportsWizards([]int64{pdsrw.Id.Get()}, pdsrw)
}

// UpdatePosDailySalesReportsWizards updates existing pos.daily.sales.reports.wizard records.
// All records (represented by ids) will be updated by pdsrw values.
func (c *Client) UpdatePosDailySalesReportsWizards(ids []int64, pdsrw *PosDailySalesReportsWizard) error {
	return c.Update(PosDailySalesReportsWizardModel, ids, pdsrw, nil)
}

// DeletePosDailySalesReportsWizard deletes an existing pos.daily.sales.reports.wizard record.
func (c *Client) DeletePosDailySalesReportsWizard(id int64) error {
	return c.DeletePosDailySalesReportsWizards([]int64{id})
}

// DeletePosDailySalesReportsWizards deletes existing pos.daily.sales.reports.wizard records.
func (c *Client) DeletePosDailySalesReportsWizards(ids []int64) error {
	return c.Delete(PosDailySalesReportsWizardModel, ids)
}

// GetPosDailySalesReportsWizard gets pos.daily.sales.reports.wizard existing record.
func (c *Client) GetPosDailySalesReportsWizard(id int64) (*PosDailySalesReportsWizard, error) {
	pdsrws, err := c.GetPosDailySalesReportsWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pdsrws)[0]), nil
}

// GetPosDailySalesReportsWizards gets pos.daily.sales.reports.wizard existing records.
func (c *Client) GetPosDailySalesReportsWizards(ids []int64) (*PosDailySalesReportsWizards, error) {
	pdsrws := &PosDailySalesReportsWizards{}
	if err := c.Read(PosDailySalesReportsWizardModel, ids, nil, pdsrws); err != nil {
		return nil, err
	}
	return pdsrws, nil
}

// FindPosDailySalesReportsWizard finds pos.daily.sales.reports.wizard record by querying it with criteria.
func (c *Client) FindPosDailySalesReportsWizard(criteria *Criteria) (*PosDailySalesReportsWizard, error) {
	pdsrws := &PosDailySalesReportsWizards{}
	if err := c.SearchRead(PosDailySalesReportsWizardModel, criteria, NewOptions().Limit(1), pdsrws); err != nil {
		return nil, err
	}
	return &((*pdsrws)[0]), nil
}

// FindPosDailySalesReportsWizards finds pos.daily.sales.reports.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDailySalesReportsWizards(criteria *Criteria, options *Options) (*PosDailySalesReportsWizards, error) {
	pdsrws := &PosDailySalesReportsWizards{}
	if err := c.SearchRead(PosDailySalesReportsWizardModel, criteria, options, pdsrws); err != nil {
		return nil, err
	}
	return pdsrws, nil
}

// FindPosDailySalesReportsWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDailySalesReportsWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosDailySalesReportsWizardModel, criteria, options)
}

// FindPosDailySalesReportsWizardId finds record id by querying it with criteria.
func (c *Client) FindPosDailySalesReportsWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosDailySalesReportsWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
