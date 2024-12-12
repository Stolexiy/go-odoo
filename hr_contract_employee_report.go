package odoo

// HrContractEmployeeReport represents hr.contract.employee.report model.
type HrContractEmployeeReport struct {
	AgeSum            *Float    `xmlrpc:"age_sum,omitempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	ContractId        *Many2One `xmlrpc:"contract_id,omitempty"`
	ContractStart     *Time     `xmlrpc:"contract_start,omitempty"`
	CountEmployeeExit *Int      `xmlrpc:"count_employee_exit,omitempty"`
	CountNewEmployee  *Int      `xmlrpc:"count_new_employee,omitempty"`
	Date              *Time     `xmlrpc:"date,omitempty"`
	DateEndContract   *Time     `xmlrpc:"date_end_contract,omitempty"`
	DepartmentId      *Many2One `xmlrpc:"department_id,omitempty"`
	DepartureReasonId *Many2One `xmlrpc:"departure_reason_id,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	EmployeeCount     *Int      `xmlrpc:"employee_count,omitempty"`
	EmployeeId        *Many2One `xmlrpc:"employee_id,omitempty"`
	EndDateMonths     *Int      `xmlrpc:"end_date_months,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	StartDateMonths   *Int      `xmlrpc:"start_date_months,omitempty"`
	Wage              *Float    `xmlrpc:"wage,omitempty"`
}

// HrContractEmployeeReports represents array of hr.contract.employee.report model.
type HrContractEmployeeReports []HrContractEmployeeReport

// HrContractEmployeeReportModel is the odoo model name.
const HrContractEmployeeReportModel = "hr.contract.employee.report"

// Many2One convert HrContractEmployeeReport to *Many2One.
func (hcer *HrContractEmployeeReport) Many2One() *Many2One {
	return NewMany2One(hcer.Id.Get(), "")
}

// CreateHrContractEmployeeReport creates a new hr.contract.employee.report model and returns its id.
func (c *Client) CreateHrContractEmployeeReport(hcer *HrContractEmployeeReport) (int64, error) {
	ids, err := c.CreateHrContractEmployeeReports([]*HrContractEmployeeReport{hcer})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrContractEmployeeReport creates a new hr.contract.employee.report model and returns its id.
func (c *Client) CreateHrContractEmployeeReports(hcers []*HrContractEmployeeReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcers {
		vv = append(vv, v)
	}
	return c.Create(HrContractEmployeeReportModel, vv, nil)
}

// UpdateHrContractEmployeeReport updates an existing hr.contract.employee.report record.
func (c *Client) UpdateHrContractEmployeeReport(hcer *HrContractEmployeeReport) error {
	return c.UpdateHrContractEmployeeReports([]int64{hcer.Id.Get()}, hcer)
}

// UpdateHrContractEmployeeReports updates existing hr.contract.employee.report records.
// All records (represented by ids) will be updated by hcer values.
func (c *Client) UpdateHrContractEmployeeReports(ids []int64, hcer *HrContractEmployeeReport) error {
	return c.Update(HrContractEmployeeReportModel, ids, hcer, nil)
}

// DeleteHrContractEmployeeReport deletes an existing hr.contract.employee.report record.
func (c *Client) DeleteHrContractEmployeeReport(id int64) error {
	return c.DeleteHrContractEmployeeReports([]int64{id})
}

// DeleteHrContractEmployeeReports deletes existing hr.contract.employee.report records.
func (c *Client) DeleteHrContractEmployeeReports(ids []int64) error {
	return c.Delete(HrContractEmployeeReportModel, ids)
}

// GetHrContractEmployeeReport gets hr.contract.employee.report existing record.
func (c *Client) GetHrContractEmployeeReport(id int64) (*HrContractEmployeeReport, error) {
	hcers, err := c.GetHrContractEmployeeReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hcers)[0]), nil
}

// GetHrContractEmployeeReports gets hr.contract.employee.report existing records.
func (c *Client) GetHrContractEmployeeReports(ids []int64) (*HrContractEmployeeReports, error) {
	hcers := &HrContractEmployeeReports{}
	if err := c.Read(HrContractEmployeeReportModel, ids, nil, hcers); err != nil {
		return nil, err
	}
	return hcers, nil
}

// FindHrContractEmployeeReport finds hr.contract.employee.report record by querying it with criteria.
func (c *Client) FindHrContractEmployeeReport(criteria *Criteria) (*HrContractEmployeeReport, error) {
	hcers := &HrContractEmployeeReports{}
	if err := c.SearchRead(HrContractEmployeeReportModel, criteria, NewOptions().Limit(1), hcers); err != nil {
		return nil, err
	}
	return &((*hcers)[0]), nil
}

// FindHrContractEmployeeReports finds hr.contract.employee.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractEmployeeReports(criteria *Criteria, options *Options) (*HrContractEmployeeReports, error) {
	hcers := &HrContractEmployeeReports{}
	if err := c.SearchRead(HrContractEmployeeReportModel, criteria, options, hcers); err != nil {
		return nil, err
	}
	return hcers, nil
}

// FindHrContractEmployeeReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractEmployeeReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrContractEmployeeReportModel, criteria, options)
}

// FindHrContractEmployeeReportId finds record id by querying it with criteria.
func (c *Client) FindHrContractEmployeeReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractEmployeeReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
