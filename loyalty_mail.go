package odoo

// LoyaltyMail represents loyalty.mail model.
type LoyaltyMail struct {
	Active           *Bool      `xmlrpc:"active,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	MailTemplateId   *Many2One  `xmlrpc:"mail_template_id,omitempty"`
	Points           *Float     `xmlrpc:"points,omitempty"`
	PosReportPrintId *Many2One  `xmlrpc:"pos_report_print_id,omitempty"`
	ProgramId        *Many2One  `xmlrpc:"program_id,omitempty"`
	Trigger          *Selection `xmlrpc:"trigger,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyMails represents array of loyalty.mail model.
type LoyaltyMails []LoyaltyMail

// LoyaltyMailModel is the odoo model name.
const LoyaltyMailModel = "loyalty.mail"

// Many2One convert LoyaltyMail to *Many2One.
func (lm *LoyaltyMail) Many2One() *Many2One {
	return NewMany2One(lm.Id.Get(), "")
}

// CreateLoyaltyMail creates a new loyalty.mail model and returns its id.
func (c *Client) CreateLoyaltyMail(lm *LoyaltyMail) (int64, error) {
	ids, err := c.CreateLoyaltyMails([]*LoyaltyMail{lm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyMail creates a new loyalty.mail model and returns its id.
func (c *Client) CreateLoyaltyMails(lms []*LoyaltyMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range lms {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyMailModel, vv, nil)
}

// UpdateLoyaltyMail updates an existing loyalty.mail record.
func (c *Client) UpdateLoyaltyMail(lm *LoyaltyMail) error {
	return c.UpdateLoyaltyMails([]int64{lm.Id.Get()}, lm)
}

// UpdateLoyaltyMails updates existing loyalty.mail records.
// All records (represented by ids) will be updated by lm values.
func (c *Client) UpdateLoyaltyMails(ids []int64, lm *LoyaltyMail) error {
	return c.Update(LoyaltyMailModel, ids, lm, nil)
}

// DeleteLoyaltyMail deletes an existing loyalty.mail record.
func (c *Client) DeleteLoyaltyMail(id int64) error {
	return c.DeleteLoyaltyMails([]int64{id})
}

// DeleteLoyaltyMails deletes existing loyalty.mail records.
func (c *Client) DeleteLoyaltyMails(ids []int64) error {
	return c.Delete(LoyaltyMailModel, ids)
}

// GetLoyaltyMail gets loyalty.mail existing record.
func (c *Client) GetLoyaltyMail(id int64) (*LoyaltyMail, error) {
	lms, err := c.GetLoyaltyMails([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lms)[0]), nil
}

// GetLoyaltyMails gets loyalty.mail existing records.
func (c *Client) GetLoyaltyMails(ids []int64) (*LoyaltyMails, error) {
	lms := &LoyaltyMails{}
	if err := c.Read(LoyaltyMailModel, ids, nil, lms); err != nil {
		return nil, err
	}
	return lms, nil
}

// FindLoyaltyMail finds loyalty.mail record by querying it with criteria.
func (c *Client) FindLoyaltyMail(criteria *Criteria) (*LoyaltyMail, error) {
	lms := &LoyaltyMails{}
	if err := c.SearchRead(LoyaltyMailModel, criteria, NewOptions().Limit(1), lms); err != nil {
		return nil, err
	}
	return &((*lms)[0]), nil
}

// FindLoyaltyMails finds loyalty.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyMails(criteria *Criteria, options *Options) (*LoyaltyMails, error) {
	lms := &LoyaltyMails{}
	if err := c.SearchRead(LoyaltyMailModel, criteria, options, lms); err != nil {
		return nil, err
	}
	return lms, nil
}

// FindLoyaltyMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyMailModel, criteria, options)
}

// FindLoyaltyMailId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
