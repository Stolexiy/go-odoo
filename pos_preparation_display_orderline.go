package odoo

// PosPreparationDisplayOrderline represents pos_preparation_display.orderline model.
type PosPreparationDisplayOrderline struct {
	AttributeValueIds         *Relation `xmlrpc:"attribute_value_ids,omitempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	InternalNote              *String   `xmlrpc:"internal_note,omitempty"`
	PosOrderLineUuid          *String   `xmlrpc:"pos_order_line_uuid,omitempty"`
	PreparationDisplayOrderId *Many2One `xmlrpc:"preparation_display_order_id,omitempty"`
	ProductCancelled          *Float    `xmlrpc:"product_cancelled,omitempty"`
	ProductId                 *Many2One `xmlrpc:"product_id,omitempty"`
	ProductQuantity           *Float    `xmlrpc:"product_quantity,omitempty"`
	Todo                      *Bool     `xmlrpc:"todo,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPreparationDisplayOrderlines represents array of pos_preparation_display.orderline model.
type PosPreparationDisplayOrderlines []PosPreparationDisplayOrderline

// PosPreparationDisplayOrderlineModel is the odoo model name.
const PosPreparationDisplayOrderlineModel = "pos_preparation_display.orderline"

// Many2One convert PosPreparationDisplayOrderline to *Many2One.
func (po *PosPreparationDisplayOrderline) Many2One() *Many2One {
	return NewMany2One(po.Id.Get(), "")
}

// CreatePosPreparationDisplayOrderline creates a new pos_preparation_display.orderline model and returns its id.
func (c *Client) CreatePosPreparationDisplayOrderline(po *PosPreparationDisplayOrderline) (int64, error) {
	ids, err := c.CreatePosPreparationDisplayOrderlines([]*PosPreparationDisplayOrderline{po})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPreparationDisplayOrderline creates a new pos_preparation_display.orderline model and returns its id.
func (c *Client) CreatePosPreparationDisplayOrderlines(pos []*PosPreparationDisplayOrderline) ([]int64, error) {
	var vv []interface{}
	for _, v := range pos {
		vv = append(vv, v)
	}
	return c.Create(PosPreparationDisplayOrderlineModel, vv, nil)
}

// UpdatePosPreparationDisplayOrderline updates an existing pos_preparation_display.orderline record.
func (c *Client) UpdatePosPreparationDisplayOrderline(po *PosPreparationDisplayOrderline) error {
	return c.UpdatePosPreparationDisplayOrderlines([]int64{po.Id.Get()}, po)
}

// UpdatePosPreparationDisplayOrderlines updates existing pos_preparation_display.orderline records.
// All records (represented by ids) will be updated by po values.
func (c *Client) UpdatePosPreparationDisplayOrderlines(ids []int64, po *PosPreparationDisplayOrderline) error {
	return c.Update(PosPreparationDisplayOrderlineModel, ids, po, nil)
}

// DeletePosPreparationDisplayOrderline deletes an existing pos_preparation_display.orderline record.
func (c *Client) DeletePosPreparationDisplayOrderline(id int64) error {
	return c.DeletePosPreparationDisplayOrderlines([]int64{id})
}

// DeletePosPreparationDisplayOrderlines deletes existing pos_preparation_display.orderline records.
func (c *Client) DeletePosPreparationDisplayOrderlines(ids []int64) error {
	return c.Delete(PosPreparationDisplayOrderlineModel, ids)
}

// GetPosPreparationDisplayOrderline gets pos_preparation_display.orderline existing record.
func (c *Client) GetPosPreparationDisplayOrderline(id int64) (*PosPreparationDisplayOrderline, error) {
	pos, err := c.GetPosPreparationDisplayOrderlines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// GetPosPreparationDisplayOrderlines gets pos_preparation_display.orderline existing records.
func (c *Client) GetPosPreparationDisplayOrderlines(ids []int64) (*PosPreparationDisplayOrderlines, error) {
	pos := &PosPreparationDisplayOrderlines{}
	if err := c.Read(PosPreparationDisplayOrderlineModel, ids, nil, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosPreparationDisplayOrderline finds pos_preparation_display.orderline record by querying it with criteria.
func (c *Client) FindPosPreparationDisplayOrderline(criteria *Criteria) (*PosPreparationDisplayOrderline, error) {
	pos := &PosPreparationDisplayOrderlines{}
	if err := c.SearchRead(PosPreparationDisplayOrderlineModel, criteria, NewOptions().Limit(1), pos); err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// FindPosPreparationDisplayOrderlines finds pos_preparation_display.orderline records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayOrderlines(criteria *Criteria, options *Options) (*PosPreparationDisplayOrderlines, error) {
	pos := &PosPreparationDisplayOrderlines{}
	if err := c.SearchRead(PosPreparationDisplayOrderlineModel, criteria, options, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosPreparationDisplayOrderlineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayOrderlineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPreparationDisplayOrderlineModel, criteria, options)
}

// FindPosPreparationDisplayOrderlineId finds record id by querying it with criteria.
func (c *Client) FindPosPreparationDisplayOrderlineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPreparationDisplayOrderlineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
