package odoo

// StockPickingBatch represents stock.picking.batch model.
type StockPickingBatch struct {
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AllowedPickingIds           *Relation   `xmlrpc:"allowed_picking_ids,omitempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	Description                 *String     `xmlrpc:"description,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty"`
	EstimatedShippingVolume     *Float      `xmlrpc:"estimated_shipping_volume,omitempty"`
	EstimatedShippingWeight     *Float      `xmlrpc:"estimated_shipping_weight,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty"`
	IsWave                      *Bool       `xmlrpc:"is_wave,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MoveIds                     *Relation   `xmlrpc:"move_ids,omitempty"`
	MoveLineIds                 *Relation   `xmlrpc:"move_line_ids,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String     `xmlrpc:"name,omitempty"`
	PickingIds                  *Relation   `xmlrpc:"picking_ids,omitempty"`
	PickingTypeCode             *Selection  `xmlrpc:"picking_type_code,omitempty"`
	PickingTypeId               *Many2One   `xmlrpc:"picking_type_id,omitempty"`
	Properties                  interface{} `xmlrpc:"properties,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	ScheduledDate               *Time       `xmlrpc:"scheduled_date,omitempty"`
	ShowAllocation              *Bool       `xmlrpc:"show_allocation,omitempty"`
	ShowCheckAvailability       *Bool       `xmlrpc:"show_check_availability,omitempty"`
	ShowLotsText                *Bool       `xmlrpc:"show_lots_text,omitempty"`
	State                       *Selection  `xmlrpc:"state,omitempty"`
	UserId                      *Many2One   `xmlrpc:"user_id,omitempty"`
	WarehouseId                 *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// StockPickingBatchs represents array of stock.picking.batch model.
type StockPickingBatchs []StockPickingBatch

// StockPickingBatchModel is the odoo model name.
const StockPickingBatchModel = "stock.picking.batch"

// Many2One convert StockPickingBatch to *Many2One.
func (spb *StockPickingBatch) Many2One() *Many2One {
	return NewMany2One(spb.Id.Get(), "")
}

// CreateStockPickingBatch creates a new stock.picking.batch model and returns its id.
func (c *Client) CreateStockPickingBatch(spb *StockPickingBatch) (int64, error) {
	ids, err := c.CreateStockPickingBatchs([]*StockPickingBatch{spb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPickingBatch creates a new stock.picking.batch model and returns its id.
func (c *Client) CreateStockPickingBatchs(spbs []*StockPickingBatch) ([]int64, error) {
	var vv []interface{}
	for _, v := range spbs {
		vv = append(vv, v)
	}
	return c.Create(StockPickingBatchModel, vv, nil)
}

// UpdateStockPickingBatch updates an existing stock.picking.batch record.
func (c *Client) UpdateStockPickingBatch(spb *StockPickingBatch) error {
	return c.UpdateStockPickingBatchs([]int64{spb.Id.Get()}, spb)
}

// UpdateStockPickingBatchs updates existing stock.picking.batch records.
// All records (represented by ids) will be updated by spb values.
func (c *Client) UpdateStockPickingBatchs(ids []int64, spb *StockPickingBatch) error {
	return c.Update(StockPickingBatchModel, ids, spb, nil)
}

// DeleteStockPickingBatch deletes an existing stock.picking.batch record.
func (c *Client) DeleteStockPickingBatch(id int64) error {
	return c.DeleteStockPickingBatchs([]int64{id})
}

// DeleteStockPickingBatchs deletes existing stock.picking.batch records.
func (c *Client) DeleteStockPickingBatchs(ids []int64) error {
	return c.Delete(StockPickingBatchModel, ids)
}

// GetStockPickingBatch gets stock.picking.batch existing record.
func (c *Client) GetStockPickingBatch(id int64) (*StockPickingBatch, error) {
	spbs, err := c.GetStockPickingBatchs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*spbs)[0]), nil
}

// GetStockPickingBatchs gets stock.picking.batch existing records.
func (c *Client) GetStockPickingBatchs(ids []int64) (*StockPickingBatchs, error) {
	spbs := &StockPickingBatchs{}
	if err := c.Read(StockPickingBatchModel, ids, nil, spbs); err != nil {
		return nil, err
	}
	return spbs, nil
}

// FindStockPickingBatch finds stock.picking.batch record by querying it with criteria.
func (c *Client) FindStockPickingBatch(criteria *Criteria) (*StockPickingBatch, error) {
	spbs := &StockPickingBatchs{}
	if err := c.SearchRead(StockPickingBatchModel, criteria, NewOptions().Limit(1), spbs); err != nil {
		return nil, err
	}
	return &((*spbs)[0]), nil
}

// FindStockPickingBatchs finds stock.picking.batch records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingBatchs(criteria *Criteria, options *Options) (*StockPickingBatchs, error) {
	spbs := &StockPickingBatchs{}
	if err := c.SearchRead(StockPickingBatchModel, criteria, options, spbs); err != nil {
		return nil, err
	}
	return spbs, nil
}

// FindStockPickingBatchIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingBatchIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPickingBatchModel, criteria, options)
}

// FindStockPickingBatchId finds record id by querying it with criteria.
func (c *Client) FindStockPickingBatchId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingBatchModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
