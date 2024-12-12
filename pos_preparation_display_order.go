package odoo

// PosPreparationDisplayOrder represents pos_preparation_display.order model.
type PosPreparationDisplayOrder struct {
	CreateDate                     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                    *String   `xmlrpc:"display_name,omitempty"`
	Displayed                      *Bool     `xmlrpc:"displayed,omitempty"`
	EmployeeId                     *Many2One `xmlrpc:"employee_id,omitempty"`
	Id                             *Int      `xmlrpc:"id,omitempty"`
	OrderStageIds                  *Relation `xmlrpc:"order_stage_ids,omitempty"`
	PdisGeneralNote                *String   `xmlrpc:"pdis_general_note,omitempty"`
	PosConfigId                    *Many2One `xmlrpc:"pos_config_id,omitempty"`
	PosOrderId                     *Many2One `xmlrpc:"pos_order_id,omitempty"`
	PreparationDisplayOrderLineIds *Relation `xmlrpc:"preparation_display_order_line_ids,omitempty"`
	WriteDate                      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPreparationDisplayOrders represents array of pos_preparation_display.order model.
type PosPreparationDisplayOrders []PosPreparationDisplayOrder

// PosPreparationDisplayOrderModel is the odoo model name.
const PosPreparationDisplayOrderModel = "pos_preparation_display.order"

// Many2One convert PosPreparationDisplayOrder to *Many2One.
func (po *PosPreparationDisplayOrder) Many2One() *Many2One {
	return NewMany2One(po.Id.Get(), "")
}

// CreatePosPreparationDisplayOrder creates a new pos_preparation_display.order model and returns its id.
func (c *Client) CreatePosPreparationDisplayOrder(po *PosPreparationDisplayOrder) (int64, error) {
	ids, err := c.CreatePosPreparationDisplayOrders([]*PosPreparationDisplayOrder{po})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPreparationDisplayOrder creates a new pos_preparation_display.order model and returns its id.
func (c *Client) CreatePosPreparationDisplayOrders(pos []*PosPreparationDisplayOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range pos {
		vv = append(vv, v)
	}
	return c.Create(PosPreparationDisplayOrderModel, vv, nil)
}

// UpdatePosPreparationDisplayOrder updates an existing pos_preparation_display.order record.
func (c *Client) UpdatePosPreparationDisplayOrder(po *PosPreparationDisplayOrder) error {
	return c.UpdatePosPreparationDisplayOrders([]int64{po.Id.Get()}, po)
}

// UpdatePosPreparationDisplayOrders updates existing pos_preparation_display.order records.
// All records (represented by ids) will be updated by po values.
func (c *Client) UpdatePosPreparationDisplayOrders(ids []int64, po *PosPreparationDisplayOrder) error {
	return c.Update(PosPreparationDisplayOrderModel, ids, po, nil)
}

// DeletePosPreparationDisplayOrder deletes an existing pos_preparation_display.order record.
func (c *Client) DeletePosPreparationDisplayOrder(id int64) error {
	return c.DeletePosPreparationDisplayOrders([]int64{id})
}

// DeletePosPreparationDisplayOrders deletes existing pos_preparation_display.order records.
func (c *Client) DeletePosPreparationDisplayOrders(ids []int64) error {
	return c.Delete(PosPreparationDisplayOrderModel, ids)
}

// GetPosPreparationDisplayOrder gets pos_preparation_display.order existing record.
func (c *Client) GetPosPreparationDisplayOrder(id int64) (*PosPreparationDisplayOrder, error) {
	pos, err := c.GetPosPreparationDisplayOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// GetPosPreparationDisplayOrders gets pos_preparation_display.order existing records.
func (c *Client) GetPosPreparationDisplayOrders(ids []int64) (*PosPreparationDisplayOrders, error) {
	pos := &PosPreparationDisplayOrders{}
	if err := c.Read(PosPreparationDisplayOrderModel, ids, nil, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosPreparationDisplayOrder finds pos_preparation_display.order record by querying it with criteria.
func (c *Client) FindPosPreparationDisplayOrder(criteria *Criteria) (*PosPreparationDisplayOrder, error) {
	pos := &PosPreparationDisplayOrders{}
	if err := c.SearchRead(PosPreparationDisplayOrderModel, criteria, NewOptions().Limit(1), pos); err != nil {
		return nil, err
	}
	return &((*pos)[0]), nil
}

// FindPosPreparationDisplayOrders finds pos_preparation_display.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayOrders(criteria *Criteria, options *Options) (*PosPreparationDisplayOrders, error) {
	pos := &PosPreparationDisplayOrders{}
	if err := c.SearchRead(PosPreparationDisplayOrderModel, criteria, options, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosPreparationDisplayOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPreparationDisplayOrderModel, criteria, options)
}

// FindPosPreparationDisplayOrderId finds record id by querying it with criteria.
func (c *Client) FindPosPreparationDisplayOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPreparationDisplayOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
