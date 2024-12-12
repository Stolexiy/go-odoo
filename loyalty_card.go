package odoo

// LoyaltyCard represents loyalty.card model.
type LoyaltyCard struct {
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	Code                     *String    `xmlrpc:"code,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	ExpirationDate           *Time      `xmlrpc:"expiration_date,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty"`
	HistoryIds               *Relation  `xmlrpc:"history_ids,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	OrderId                  *Many2One  `xmlrpc:"order_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PointName                *String    `xmlrpc:"point_name,omitempty"`
	Points                   *Float     `xmlrpc:"points,omitempty"`
	PointsDisplay            *String    `xmlrpc:"points_display,omitempty"`
	ProgramId                *Many2One  `xmlrpc:"program_id,omitempty"`
	ProgramType              *Selection `xmlrpc:"program_type,omitempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omitempty"`
	SourcePosOrderId         *Many2One  `xmlrpc:"source_pos_order_id,omitempty"`
	UseCount                 *Int       `xmlrpc:"use_count,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LoyaltyCards represents array of loyalty.card model.
type LoyaltyCards []LoyaltyCard

// LoyaltyCardModel is the odoo model name.
const LoyaltyCardModel = "loyalty.card"

// Many2One convert LoyaltyCard to *Many2One.
func (lc *LoyaltyCard) Many2One() *Many2One {
	return NewMany2One(lc.Id.Get(), "")
}

// CreateLoyaltyCard creates a new loyalty.card model and returns its id.
func (c *Client) CreateLoyaltyCard(lc *LoyaltyCard) (int64, error) {
	ids, err := c.CreateLoyaltyCards([]*LoyaltyCard{lc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLoyaltyCard creates a new loyalty.card model and returns its id.
func (c *Client) CreateLoyaltyCards(lcs []*LoyaltyCard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lcs {
		vv = append(vv, v)
	}
	return c.Create(LoyaltyCardModel, vv, nil)
}

// UpdateLoyaltyCard updates an existing loyalty.card record.
func (c *Client) UpdateLoyaltyCard(lc *LoyaltyCard) error {
	return c.UpdateLoyaltyCards([]int64{lc.Id.Get()}, lc)
}

// UpdateLoyaltyCards updates existing loyalty.card records.
// All records (represented by ids) will be updated by lc values.
func (c *Client) UpdateLoyaltyCards(ids []int64, lc *LoyaltyCard) error {
	return c.Update(LoyaltyCardModel, ids, lc, nil)
}

// DeleteLoyaltyCard deletes an existing loyalty.card record.
func (c *Client) DeleteLoyaltyCard(id int64) error {
	return c.DeleteLoyaltyCards([]int64{id})
}

// DeleteLoyaltyCards deletes existing loyalty.card records.
func (c *Client) DeleteLoyaltyCards(ids []int64) error {
	return c.Delete(LoyaltyCardModel, ids)
}

// GetLoyaltyCard gets loyalty.card existing record.
func (c *Client) GetLoyaltyCard(id int64) (*LoyaltyCard, error) {
	lcs, err := c.GetLoyaltyCards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lcs)[0]), nil
}

// GetLoyaltyCards gets loyalty.card existing records.
func (c *Client) GetLoyaltyCards(ids []int64) (*LoyaltyCards, error) {
	lcs := &LoyaltyCards{}
	if err := c.Read(LoyaltyCardModel, ids, nil, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLoyaltyCard finds loyalty.card record by querying it with criteria.
func (c *Client) FindLoyaltyCard(criteria *Criteria) (*LoyaltyCard, error) {
	lcs := &LoyaltyCards{}
	if err := c.SearchRead(LoyaltyCardModel, criteria, NewOptions().Limit(1), lcs); err != nil {
		return nil, err
	}
	return &((*lcs)[0]), nil
}

// FindLoyaltyCards finds loyalty.card records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyCards(criteria *Criteria, options *Options) (*LoyaltyCards, error) {
	lcs := &LoyaltyCards{}
	if err := c.SearchRead(LoyaltyCardModel, criteria, options, lcs); err != nil {
		return nil, err
	}
	return lcs, nil
}

// FindLoyaltyCardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyCardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LoyaltyCardModel, criteria, options)
}

// FindLoyaltyCardId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyCardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyCardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
