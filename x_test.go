package odoo

// XTest represents x_test model.
type XTest struct {
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
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
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
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
	XActive                     *Bool      `xmlrpc:"x_active,omitempty"`
	XName                       *String    `xmlrpc:"x_name,omitempty"`
	XStudioSequence             *Int       `xmlrpc:"x_studio_sequence,omitempty"`
}

// XTests represents array of x_test model.
type XTests []XTest

// XTestModel is the odoo model name.
const XTestModel = "x_test"

// Many2One convert XTest to *Many2One.
func (x *XTest) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXTest creates a new x_test model and returns its id.
func (c *Client) CreateXTest(x *XTest) (int64, error) {
	ids, err := c.CreateXTests([]*XTest{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXTest creates a new x_test model and returns its id.
func (c *Client) CreateXTests(xs []*XTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XTestModel, vv, nil)
}

// UpdateXTest updates an existing x_test record.
func (c *Client) UpdateXTest(x *XTest) error {
	return c.UpdateXTests([]int64{x.Id.Get()}, x)
}

// UpdateXTests updates existing x_test records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXTests(ids []int64, x *XTest) error {
	return c.Update(XTestModel, ids, x, nil)
}

// DeleteXTest deletes an existing x_test record.
func (c *Client) DeleteXTest(id int64) error {
	return c.DeleteXTests([]int64{id})
}

// DeleteXTests deletes existing x_test records.
func (c *Client) DeleteXTests(ids []int64) error {
	return c.Delete(XTestModel, ids)
}

// GetXTest gets x_test existing record.
func (c *Client) GetXTest(id int64) (*XTest, error) {
	xs, err := c.GetXTests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXTests gets x_test existing records.
func (c *Client) GetXTests(ids []int64) (*XTests, error) {
	xs := &XTests{}
	if err := c.Read(XTestModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXTest finds x_test record by querying it with criteria.
func (c *Client) FindXTest(criteria *Criteria) (*XTest, error) {
	xs := &XTests{}
	if err := c.SearchRead(XTestModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXTests finds x_test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXTests(criteria *Criteria, options *Options) (*XTests, error) {
	xs := &XTests{}
	if err := c.SearchRead(XTestModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XTestModel, criteria, options)
}

// FindXTestId finds record id by querying it with criteria.
func (c *Client) FindXTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
