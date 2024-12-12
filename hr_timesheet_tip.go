package odoo

// HrTimesheetTip represents hr.timesheet.tip model.
type HrTimesheetTip struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrTimesheetTips represents array of hr.timesheet.tip model.
type HrTimesheetTips []HrTimesheetTip

// HrTimesheetTipModel is the odoo model name.
const HrTimesheetTipModel = "hr.timesheet.tip"

// Many2One convert HrTimesheetTip to *Many2One.
func (htt *HrTimesheetTip) Many2One() *Many2One {
	return NewMany2One(htt.Id.Get(), "")
}

// CreateHrTimesheetTip creates a new hr.timesheet.tip model and returns its id.
func (c *Client) CreateHrTimesheetTip(htt *HrTimesheetTip) (int64, error) {
	ids, err := c.CreateHrTimesheetTips([]*HrTimesheetTip{htt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrTimesheetTip creates a new hr.timesheet.tip model and returns its id.
func (c *Client) CreateHrTimesheetTips(htts []*HrTimesheetTip) ([]int64, error) {
	var vv []interface{}
	for _, v := range htts {
		vv = append(vv, v)
	}
	return c.Create(HrTimesheetTipModel, vv, nil)
}

// UpdateHrTimesheetTip updates an existing hr.timesheet.tip record.
func (c *Client) UpdateHrTimesheetTip(htt *HrTimesheetTip) error {
	return c.UpdateHrTimesheetTips([]int64{htt.Id.Get()}, htt)
}

// UpdateHrTimesheetTips updates existing hr.timesheet.tip records.
// All records (represented by ids) will be updated by htt values.
func (c *Client) UpdateHrTimesheetTips(ids []int64, htt *HrTimesheetTip) error {
	return c.Update(HrTimesheetTipModel, ids, htt, nil)
}

// DeleteHrTimesheetTip deletes an existing hr.timesheet.tip record.
func (c *Client) DeleteHrTimesheetTip(id int64) error {
	return c.DeleteHrTimesheetTips([]int64{id})
}

// DeleteHrTimesheetTips deletes existing hr.timesheet.tip records.
func (c *Client) DeleteHrTimesheetTips(ids []int64) error {
	return c.Delete(HrTimesheetTipModel, ids)
}

// GetHrTimesheetTip gets hr.timesheet.tip existing record.
func (c *Client) GetHrTimesheetTip(id int64) (*HrTimesheetTip, error) {
	htts, err := c.GetHrTimesheetTips([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*htts)[0]), nil
}

// GetHrTimesheetTips gets hr.timesheet.tip existing records.
func (c *Client) GetHrTimesheetTips(ids []int64) (*HrTimesheetTips, error) {
	htts := &HrTimesheetTips{}
	if err := c.Read(HrTimesheetTipModel, ids, nil, htts); err != nil {
		return nil, err
	}
	return htts, nil
}

// FindHrTimesheetTip finds hr.timesheet.tip record by querying it with criteria.
func (c *Client) FindHrTimesheetTip(criteria *Criteria) (*HrTimesheetTip, error) {
	htts := &HrTimesheetTips{}
	if err := c.SearchRead(HrTimesheetTipModel, criteria, NewOptions().Limit(1), htts); err != nil {
		return nil, err
	}
	return &((*htts)[0]), nil
}

// FindHrTimesheetTips finds hr.timesheet.tip records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetTips(criteria *Criteria, options *Options) (*HrTimesheetTips, error) {
	htts := &HrTimesheetTips{}
	if err := c.SearchRead(HrTimesheetTipModel, criteria, options, htts); err != nil {
		return nil, err
	}
	return htts, nil
}

// FindHrTimesheetTipIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetTipIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrTimesheetTipModel, criteria, options)
}

// FindHrTimesheetTipId finds record id by querying it with criteria.
func (c *Client) FindHrTimesheetTipId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrTimesheetTipModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
