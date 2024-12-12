package odoo

// LoyaltyGenerateWizard represents loyalty.generate.wizard model.
type LoyaltyGenerateWizard struct {
	ConfirmationMessage *String    `xmlrpc:"confirmation_message,omitempty"`
	CouponQty           *Int       `xmlrpc:"coupon_qty,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomerIds         *Relation  `xmlrpc:"customer_ids,omitempty"`
	CustomerTagIds      *Relation  `xmlrpc:"customer_tag_ids,omitempty"`
	Description         *String    `xmlrpc:"description,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	Mode                *Selection `xmlrpc:"mode,omitempty"`
	PointsGranted       *Float     `xmlrpc:"points_granted,omitempty"`
	PointsName          *String    `xmlrpc:"points_name,omitempty"`
	ProgramId           *Many2One  `xmlrpc:"program_id,omitempty"`
	ProgramType         *Selection `xmlrpc:"program_type,omitempty"`
	ValidUntil          *Time      `xmlrpc:"valid_until,omitempty"`
	WillSendMail        *Bool      `xmlrpc:"will_send_mail,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyGenerateWizards represents array of loyalty.generate.wizard model.
type LoyaltyGenerateWizards []LoyaltyGenerateWizard

// LoyaltyGenerateWizardModel is the odoo model name.
const LoyaltyGenerateWizardModel = "loyalty.generate.wizard"

// Many2One convert LoyaltyGenerateWizard to *Many2One.
func (lgw *LoyaltyGenerateWizard) Many2One() *Many2One {
	return NewMany2One(lgw.Id.Get(), "")
}

// CreateLoyaltyGenerateWizard creates a new loyalty.generate.wizard model and returns its id.
func (c *Client) CreateLoyaltyGenerateWizard(lgw *LoyaltyGenerateWizard) (int64, error) {
	ids, err := c.CreateLoyaltyGenerateWizards([]*LoyaltyGenerateWizard{lgw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyGenerateWizard creates a new loyalty.generate.wizard model and returns its id.
func (c *Client) CreateLoyaltyGenerateWizards(lgws []*LoyaltyGenerateWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lgws {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyGenerateWizardModel, vv, nil)
}

// UpdateLoyaltyGenerateWizard updates an existing loyalty.generate.wizard record.
func (c *Client) UpdateLoyaltyGenerateWizard(lgw *LoyaltyGenerateWizard) error {
	return c.UpdateLoyaltyGenerateWizards([]int64{lgw.Id.Get()}, lgw)
}

// UpdateLoyaltyGenerateWizards updates existing loyalty.generate.wizard records.
// All records (represented by ids) will be updated by lgw values.
func (c *Client) UpdateLoyaltyGenerateWizards(ids []int64, lgw *LoyaltyGenerateWizard) error {
	return c.Update(LoyaltyGenerateWizardModel, ids, lgw, nil)
}

// DeleteLoyaltyGenerateWizard deletes an existing loyalty.generate.wizard record.
func (c *Client) DeleteLoyaltyGenerateWizard(id int64) error {
	return c.DeleteLoyaltyGenerateWizards([]int64{id})
}

// DeleteLoyaltyGenerateWizards deletes existing loyalty.generate.wizard records.
func (c *Client) DeleteLoyaltyGenerateWizards(ids []int64) error {
	return c.Delete(LoyaltyGenerateWizardModel, ids)
}

// GetLoyaltyGenerateWizard gets loyalty.generate.wizard existing record.
func (c *Client) GetLoyaltyGenerateWizard(id int64) (*LoyaltyGenerateWizard, error) {
	lgws, err := c.GetLoyaltyGenerateWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lgws)[0]), nil
}

// GetLoyaltyGenerateWizards gets loyalty.generate.wizard existing records.
func (c *Client) GetLoyaltyGenerateWizards(ids []int64) (*LoyaltyGenerateWizards, error) {
	lgws := &LoyaltyGenerateWizards{}
	if err := c.Read(LoyaltyGenerateWizardModel, ids, nil, lgws); err != nil {
		return nil, err
	}
	return lgws, nil
}

// FindLoyaltyGenerateWizard finds loyalty.generate.wizard record by querying it with criteria.
func (c *Client) FindLoyaltyGenerateWizard(criteria *Criteria) (*LoyaltyGenerateWizard, error) {
	lgws := &LoyaltyGenerateWizards{}
	if err := c.SearchRead(LoyaltyGenerateWizardModel, criteria, NewOptions().Limit(1), lgws); err != nil {
		return nil, err
	}
	return &((*lgws)[0]), nil
}

// FindLoyaltyGenerateWizards finds loyalty.generate.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyGenerateWizards(criteria *Criteria, options *Options) (*LoyaltyGenerateWizards, error) {
	lgws := &LoyaltyGenerateWizards{}
	if err := c.SearchRead(LoyaltyGenerateWizardModel, criteria, options, lgws); err != nil {
		return nil, err
	}
	return lgws, nil
}

// FindLoyaltyGenerateWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyGenerateWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyGenerateWizardModel, criteria, options)
}

// FindLoyaltyGenerateWizardId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyGenerateWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyGenerateWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
