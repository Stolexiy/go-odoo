package odoo

// StockPickingToBatch represents stock.picking.to.batch model.
type StockPickingToBatch struct {
	BatchId       *Many2One  `xmlrpc:"batch_id,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description   *String    `xmlrpc:"description,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	IsCreateDraft *Bool      `xmlrpc:"is_create_draft,omitempty"`
	Mode          *Selection `xmlrpc:"mode,omitempty"`
	UserId        *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockPickingToBatchs represents array of stock.picking.to.batch model.
type StockPickingToBatchs []StockPickingToBatch

// StockPickingToBatchModel is the odoo model name.
const StockPickingToBatchModel = "stock.picking.to.batch"

// Many2One convert StockPickingToBatch to *Many2One.
func (sptb *StockPickingToBatch) Many2One() *Many2One {
	return NewMany2One(sptb.Id.Get(), "")
}

// CreateStockPickingToBatch creates a new stock.picking.to.batch model and returns its id.
func (c *Client) CreateStockPickingToBatch(sptb *StockPickingToBatch) (int64, error) {
	ids, err := c.CreateStockPickingToBatchs([]*StockPickingToBatch{sptb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPickingToBatch creates a new stock.picking.to.batch model and returns its id.
func (c *Client) CreateStockPickingToBatchs(sptbs []*StockPickingToBatch) ([]int64, error) {
	var vv []interface{}
	for _, v := range sptbs {
		vv = append(vv, v)
	}
	return c.Create(StockPickingToBatchModel, vv, nil)
}

// UpdateStockPickingToBatch updates an existing stock.picking.to.batch record.
func (c *Client) UpdateStockPickingToBatch(sptb *StockPickingToBatch) error {
	return c.UpdateStockPickingToBatchs([]int64{sptb.Id.Get()}, sptb)
}

// UpdateStockPickingToBatchs updates existing stock.picking.to.batch records.
// All records (represented by ids) will be updated by sptb values.
func (c *Client) UpdateStockPickingToBatchs(ids []int64, sptb *StockPickingToBatch) error {
	return c.Update(StockPickingToBatchModel, ids, sptb, nil)
}

// DeleteStockPickingToBatch deletes an existing stock.picking.to.batch record.
func (c *Client) DeleteStockPickingToBatch(id int64) error {
	return c.DeleteStockPickingToBatchs([]int64{id})
}

// DeleteStockPickingToBatchs deletes existing stock.picking.to.batch records.
func (c *Client) DeleteStockPickingToBatchs(ids []int64) error {
	return c.Delete(StockPickingToBatchModel, ids)
}

// GetStockPickingToBatch gets stock.picking.to.batch existing record.
func (c *Client) GetStockPickingToBatch(id int64) (*StockPickingToBatch, error) {
	sptbs, err := c.GetStockPickingToBatchs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sptbs)[0]), nil
}

// GetStockPickingToBatchs gets stock.picking.to.batch existing records.
func (c *Client) GetStockPickingToBatchs(ids []int64) (*StockPickingToBatchs, error) {
	sptbs := &StockPickingToBatchs{}
	if err := c.Read(StockPickingToBatchModel, ids, nil, sptbs); err != nil {
		return nil, err
	}
	return sptbs, nil
}

// FindStockPickingToBatch finds stock.picking.to.batch record by querying it with criteria.
func (c *Client) FindStockPickingToBatch(criteria *Criteria) (*StockPickingToBatch, error) {
	sptbs := &StockPickingToBatchs{}
	if err := c.SearchRead(StockPickingToBatchModel, criteria, NewOptions().Limit(1), sptbs); err != nil {
		return nil, err
	}
	return &((*sptbs)[0]), nil
}

// FindStockPickingToBatchs finds stock.picking.to.batch records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingToBatchs(criteria *Criteria, options *Options) (*StockPickingToBatchs, error) {
	sptbs := &StockPickingToBatchs{}
	if err := c.SearchRead(StockPickingToBatchModel, criteria, options, sptbs); err != nil {
		return nil, err
	}
	return sptbs, nil
}

// FindStockPickingToBatchIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingToBatchIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPickingToBatchModel, criteria, options)
}

// FindStockPickingToBatchId finds record id by querying it with criteria.
func (c *Client) FindStockPickingToBatchId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingToBatchModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
