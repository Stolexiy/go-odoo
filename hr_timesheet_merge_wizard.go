package odoo

// HrTimesheetMergeWizard represents hr_timesheet.merge.wizard model.
type HrTimesheetMergeWizard struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	Date          *Time     `xmlrpc:"date,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omitempty"`
	EncodingUomId *Many2One `xmlrpc:"encoding_uom_id,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	ProjectId     *Many2One `xmlrpc:"project_id,omitempty"`
	TaskId        *Many2One `xmlrpc:"task_id,omitempty"`
	TimesheetIds  *Relation `xmlrpc:"timesheet_ids,omitempty"`
	UnitAmount    *Float    `xmlrpc:"unit_amount,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrTimesheetMergeWizards represents array of hr_timesheet.merge.wizard model.
type HrTimesheetMergeWizards []HrTimesheetMergeWizard

// HrTimesheetMergeWizardModel is the odoo model name.
const HrTimesheetMergeWizardModel = "hr_timesheet.merge.wizard"

// Many2One convert HrTimesheetMergeWizard to *Many2One.
func (hmw *HrTimesheetMergeWizard) Many2One() *Many2One {
	return NewMany2One(hmw.Id.Get(), "")
}

// CreateHrTimesheetMergeWizard creates a new hr_timesheet.merge.wizard model and returns its id.
func (c *Client) CreateHrTimesheetMergeWizard(hmw *HrTimesheetMergeWizard) (int64, error) {
	ids, err := c.CreateHrTimesheetMergeWizards([]*HrTimesheetMergeWizard{hmw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrTimesheetMergeWizard creates a new hr_timesheet.merge.wizard model and returns its id.
func (c *Client) CreateHrTimesheetMergeWizards(hmws []*HrTimesheetMergeWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hmws {
		vv = append(vv, v)
	}
	return c.Create(HrTimesheetMergeWizardModel, vv, nil)
}

// UpdateHrTimesheetMergeWizard updates an existing hr_timesheet.merge.wizard record.
func (c *Client) UpdateHrTimesheetMergeWizard(hmw *HrTimesheetMergeWizard) error {
	return c.UpdateHrTimesheetMergeWizards([]int64{hmw.Id.Get()}, hmw)
}

// UpdateHrTimesheetMergeWizards updates existing hr_timesheet.merge.wizard records.
// All records (represented by ids) will be updated by hmw values.
func (c *Client) UpdateHrTimesheetMergeWizards(ids []int64, hmw *HrTimesheetMergeWizard) error {
	return c.Update(HrTimesheetMergeWizardModel, ids, hmw, nil)
}

// DeleteHrTimesheetMergeWizard deletes an existing hr_timesheet.merge.wizard record.
func (c *Client) DeleteHrTimesheetMergeWizard(id int64) error {
	return c.DeleteHrTimesheetMergeWizards([]int64{id})
}

// DeleteHrTimesheetMergeWizards deletes existing hr_timesheet.merge.wizard records.
func (c *Client) DeleteHrTimesheetMergeWizards(ids []int64) error {
	return c.Delete(HrTimesheetMergeWizardModel, ids)
}

// GetHrTimesheetMergeWizard gets hr_timesheet.merge.wizard existing record.
func (c *Client) GetHrTimesheetMergeWizard(id int64) (*HrTimesheetMergeWizard, error) {
	hmws, err := c.GetHrTimesheetMergeWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hmws)[0]), nil
}

// GetHrTimesheetMergeWizards gets hr_timesheet.merge.wizard existing records.
func (c *Client) GetHrTimesheetMergeWizards(ids []int64) (*HrTimesheetMergeWizards, error) {
	hmws := &HrTimesheetMergeWizards{}
	if err := c.Read(HrTimesheetMergeWizardModel, ids, nil, hmws); err != nil {
		return nil, err
	}
	return hmws, nil
}

// FindHrTimesheetMergeWizard finds hr_timesheet.merge.wizard record by querying it with criteria.
func (c *Client) FindHrTimesheetMergeWizard(criteria *Criteria) (*HrTimesheetMergeWizard, error) {
	hmws := &HrTimesheetMergeWizards{}
	if err := c.SearchRead(HrTimesheetMergeWizardModel, criteria, NewOptions().Limit(1), hmws); err != nil {
		return nil, err
	}
	return &((*hmws)[0]), nil
}

// FindHrTimesheetMergeWizards finds hr_timesheet.merge.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetMergeWizards(criteria *Criteria, options *Options) (*HrTimesheetMergeWizards, error) {
	hmws := &HrTimesheetMergeWizards{}
	if err := c.SearchRead(HrTimesheetMergeWizardModel, criteria, options, hmws); err != nil {
		return nil, err
	}
	return hmws, nil
}

// FindHrTimesheetMergeWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetMergeWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrTimesheetMergeWizardModel, criteria, options)
}

// FindHrTimesheetMergeWizardId finds record id by querying it with criteria.
func (c *Client) FindHrTimesheetMergeWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrTimesheetMergeWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
