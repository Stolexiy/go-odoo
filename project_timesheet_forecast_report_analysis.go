package odoo

// ProjectTimesheetForecastReportAnalysis represents project.timesheet.forecast.report.analysis model.
type ProjectTimesheetForecastReportAnalysis struct {
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	Difference                *Float     `xmlrpc:"difference,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	EffectiveBillableHours    *Float     `xmlrpc:"effective_billable_hours,omitempty"`
	EffectiveCosts            *Float     `xmlrpc:"effective_costs,omitempty"`
	EffectiveHours            *Float     `xmlrpc:"effective_hours,omitempty"`
	EffectiveMargin           *Float     `xmlrpc:"effective_margin,omitempty"`
	EffectiveNonBillableHours *Float     `xmlrpc:"effective_non_billable_hours,omitempty"`
	EffectiveRevenues         *Float     `xmlrpc:"effective_revenues,omitempty"`
	EmployeeId                *Many2One  `xmlrpc:"employee_id,omitempty"`
	EntryDate                 *Time      `xmlrpc:"entry_date,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omitempty"`
	LineType                  *Selection `xmlrpc:"line_type,omitempty"`
	PlannedBillableHours      *Float     `xmlrpc:"planned_billable_hours,omitempty"`
	PlannedCosts              *Float     `xmlrpc:"planned_costs,omitempty"`
	PlannedHours              *Float     `xmlrpc:"planned_hours,omitempty"`
	PlannedMargin             *Float     `xmlrpc:"planned_margin,omitempty"`
	PlannedNonBillableHours   *Float     `xmlrpc:"planned_non_billable_hours,omitempty"`
	PlannedRevenues           *Float     `xmlrpc:"planned_revenues,omitempty"`
	ProjectId                 *Many2One  `xmlrpc:"project_id,omitempty"`
	UserId                    *Many2One  `xmlrpc:"user_id,omitempty"`
}

// ProjectTimesheetForecastReportAnalysiss represents array of project.timesheet.forecast.report.analysis model.
type ProjectTimesheetForecastReportAnalysiss []ProjectTimesheetForecastReportAnalysis

// ProjectTimesheetForecastReportAnalysisModel is the odoo model name.
const ProjectTimesheetForecastReportAnalysisModel = "project.timesheet.forecast.report.analysis"

// Many2One convert ProjectTimesheetForecastReportAnalysis to *Many2One.
func (ptfra *ProjectTimesheetForecastReportAnalysis) Many2One() *Many2One {
	return NewMany2One(ptfra.Id.Get(), "")
}

// CreateProjectTimesheetForecastReportAnalysis creates a new project.timesheet.forecast.report.analysis model and returns its id.
func (c *Client) CreateProjectTimesheetForecastReportAnalysis(ptfra *ProjectTimesheetForecastReportAnalysis) (int64, error) {
	ids, err := c.CreateProjectTimesheetForecastReportAnalysiss([]*ProjectTimesheetForecastReportAnalysis{ptfra})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTimesheetForecastReportAnalysis creates a new project.timesheet.forecast.report.analysis model and returns its id.
func (c *Client) CreateProjectTimesheetForecastReportAnalysiss(ptfras []*ProjectTimesheetForecastReportAnalysis) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptfras {
		vv = append(vv, v)
	}
	return c.Create(ProjectTimesheetForecastReportAnalysisModel, vv, nil)
}

// UpdateProjectTimesheetForecastReportAnalysis updates an existing project.timesheet.forecast.report.analysis record.
func (c *Client) UpdateProjectTimesheetForecastReportAnalysis(ptfra *ProjectTimesheetForecastReportAnalysis) error {
	return c.UpdateProjectTimesheetForecastReportAnalysiss([]int64{ptfra.Id.Get()}, ptfra)
}

// UpdateProjectTimesheetForecastReportAnalysiss updates existing project.timesheet.forecast.report.analysis records.
// All records (represented by ids) will be updated by ptfra values.
func (c *Client) UpdateProjectTimesheetForecastReportAnalysiss(ids []int64, ptfra *ProjectTimesheetForecastReportAnalysis) error {
	return c.Update(ProjectTimesheetForecastReportAnalysisModel, ids, ptfra, nil)
}

// DeleteProjectTimesheetForecastReportAnalysis deletes an existing project.timesheet.forecast.report.analysis record.
func (c *Client) DeleteProjectTimesheetForecastReportAnalysis(id int64) error {
	return c.DeleteProjectTimesheetForecastReportAnalysiss([]int64{id})
}

// DeleteProjectTimesheetForecastReportAnalysiss deletes existing project.timesheet.forecast.report.analysis records.
func (c *Client) DeleteProjectTimesheetForecastReportAnalysiss(ids []int64) error {
	return c.Delete(ProjectTimesheetForecastReportAnalysisModel, ids)
}

// GetProjectTimesheetForecastReportAnalysis gets project.timesheet.forecast.report.analysis existing record.
func (c *Client) GetProjectTimesheetForecastReportAnalysis(id int64) (*ProjectTimesheetForecastReportAnalysis, error) {
	ptfras, err := c.GetProjectTimesheetForecastReportAnalysiss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptfras)[0]), nil
}

// GetProjectTimesheetForecastReportAnalysiss gets project.timesheet.forecast.report.analysis existing records.
func (c *Client) GetProjectTimesheetForecastReportAnalysiss(ids []int64) (*ProjectTimesheetForecastReportAnalysiss, error) {
	ptfras := &ProjectTimesheetForecastReportAnalysiss{}
	if err := c.Read(ProjectTimesheetForecastReportAnalysisModel, ids, nil, ptfras); err != nil {
		return nil, err
	}
	return ptfras, nil
}

// FindProjectTimesheetForecastReportAnalysis finds project.timesheet.forecast.report.analysis record by querying it with criteria.
func (c *Client) FindProjectTimesheetForecastReportAnalysis(criteria *Criteria) (*ProjectTimesheetForecastReportAnalysis, error) {
	ptfras := &ProjectTimesheetForecastReportAnalysiss{}
	if err := c.SearchRead(ProjectTimesheetForecastReportAnalysisModel, criteria, NewOptions().Limit(1), ptfras); err != nil {
		return nil, err
	}
	return &((*ptfras)[0]), nil
}

// FindProjectTimesheetForecastReportAnalysiss finds project.timesheet.forecast.report.analysis records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTimesheetForecastReportAnalysiss(criteria *Criteria, options *Options) (*ProjectTimesheetForecastReportAnalysiss, error) {
	ptfras := &ProjectTimesheetForecastReportAnalysiss{}
	if err := c.SearchRead(ProjectTimesheetForecastReportAnalysisModel, criteria, options, ptfras); err != nil {
		return nil, err
	}
	return ptfras, nil
}

// FindProjectTimesheetForecastReportAnalysisIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTimesheetForecastReportAnalysisIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTimesheetForecastReportAnalysisModel, criteria, options)
}

// FindProjectTimesheetForecastReportAnalysisId finds record id by querying it with criteria.
func (c *Client) FindProjectTimesheetForecastReportAnalysisId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTimesheetForecastReportAnalysisModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
