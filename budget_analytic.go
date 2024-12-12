package odoo

// BudgetAnalytic represents budget.analytic model.
type BudgetAnalytic struct {
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	BudgetLineIds               *Relation  `xmlrpc:"budget_line_ids,omitempty"`
	BudgetType                  *Selection `xmlrpc:"budget_type,omitempty"`
	ChildrenIds                 *Relation  `xmlrpc:"children_ids,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omitempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BudgetAnalytics represents array of budget.analytic model.
type BudgetAnalytics []BudgetAnalytic

// BudgetAnalyticModel is the odoo model name.
const BudgetAnalyticModel = "budget.analytic"

// Many2One convert BudgetAnalytic to *Many2One.
func (ba *BudgetAnalytic) Many2One() *Many2One {
	return NewMany2One(ba.Id.Get(), "")
}

// CreateBudgetAnalytic creates a new budget.analytic model and returns its id.
func (c *Client) CreateBudgetAnalytic(ba *BudgetAnalytic) (int64, error) {
	ids, err := c.CreateBudgetAnalytics([]*BudgetAnalytic{ba})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBudgetAnalytic creates a new budget.analytic model and returns its id.
func (c *Client) CreateBudgetAnalytics(bas []*BudgetAnalytic) ([]int64, error) {
	var vv []interface{}
	for _, v := range bas {
		vv = append(vv, v)
	}
	return c.Create(BudgetAnalyticModel, vv, nil)
}

// UpdateBudgetAnalytic updates an existing budget.analytic record.
func (c *Client) UpdateBudgetAnalytic(ba *BudgetAnalytic) error {
	return c.UpdateBudgetAnalytics([]int64{ba.Id.Get()}, ba)
}

// UpdateBudgetAnalytics updates existing budget.analytic records.
// All records (represented by ids) will be updated by ba values.
func (c *Client) UpdateBudgetAnalytics(ids []int64, ba *BudgetAnalytic) error {
	return c.Update(BudgetAnalyticModel, ids, ba, nil)
}

// DeleteBudgetAnalytic deletes an existing budget.analytic record.
func (c *Client) DeleteBudgetAnalytic(id int64) error {
	return c.DeleteBudgetAnalytics([]int64{id})
}

// DeleteBudgetAnalytics deletes existing budget.analytic records.
func (c *Client) DeleteBudgetAnalytics(ids []int64) error {
	return c.Delete(BudgetAnalyticModel, ids)
}

// GetBudgetAnalytic gets budget.analytic existing record.
func (c *Client) GetBudgetAnalytic(id int64) (*BudgetAnalytic, error) {
	bas, err := c.GetBudgetAnalytics([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bas)[0]), nil
}

// GetBudgetAnalytics gets budget.analytic existing records.
func (c *Client) GetBudgetAnalytics(ids []int64) (*BudgetAnalytics, error) {
	bas := &BudgetAnalytics{}
	if err := c.Read(BudgetAnalyticModel, ids, nil, bas); err != nil {
		return nil, err
	}
	return bas, nil
}

// FindBudgetAnalytic finds budget.analytic record by querying it with criteria.
func (c *Client) FindBudgetAnalytic(criteria *Criteria) (*BudgetAnalytic, error) {
	bas := &BudgetAnalytics{}
	if err := c.SearchRead(BudgetAnalyticModel, criteria, NewOptions().Limit(1), bas); err != nil {
		return nil, err
	}
	return &((*bas)[0]), nil
}

// FindBudgetAnalytics finds budget.analytic records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetAnalytics(criteria *Criteria, options *Options) (*BudgetAnalytics, error) {
	bas := &BudgetAnalytics{}
	if err := c.SearchRead(BudgetAnalyticModel, criteria, options, bas); err != nil {
		return nil, err
	}
	return bas, nil
}

// FindBudgetAnalyticIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBudgetAnalyticIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BudgetAnalyticModel, criteria, options)
}

// FindBudgetAnalyticId finds record id by querying it with criteria.
func (c *Client) FindBudgetAnalyticId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BudgetAnalyticModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
