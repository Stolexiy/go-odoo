package odoo

// StockAddToWave represents stock.add.to.wave model.
type StockAddToWave struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	LineIds     *Relation  `xmlrpc:"line_ids,omitempty"`
	Mode        *Selection `xmlrpc:"mode,omitempty"`
	PickingIds  *Relation  `xmlrpc:"picking_ids,omitempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omitempty"`
	WaveId      *Many2One  `xmlrpc:"wave_id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockAddToWaves represents array of stock.add.to.wave model.
type StockAddToWaves []StockAddToWave

// StockAddToWaveModel is the odoo model name.
const StockAddToWaveModel = "stock.add.to.wave"

// Many2One convert StockAddToWave to *Many2One.
func (satw *StockAddToWave) Many2One() *Many2One {
	return NewMany2One(satw.Id.Get(), "")
}

// CreateStockAddToWave creates a new stock.add.to.wave model and returns its id.
func (c *Client) CreateStockAddToWave(satw *StockAddToWave) (int64, error) {
	ids, err := c.CreateStockAddToWaves([]*StockAddToWave{satw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockAddToWave creates a new stock.add.to.wave model and returns its id.
func (c *Client) CreateStockAddToWaves(satws []*StockAddToWave) ([]int64, error) {
	var vv []interface{}
	for _, v := range satws {
		vv = append(vv, v)
	}
	return c.Create(StockAddToWaveModel, vv, nil)
}

// UpdateStockAddToWave updates an existing stock.add.to.wave record.
func (c *Client) UpdateStockAddToWave(satw *StockAddToWave) error {
	return c.UpdateStockAddToWaves([]int64{satw.Id.Get()}, satw)
}

// UpdateStockAddToWaves updates existing stock.add.to.wave records.
// All records (represented by ids) will be updated by satw values.
func (c *Client) UpdateStockAddToWaves(ids []int64, satw *StockAddToWave) error {
	return c.Update(StockAddToWaveModel, ids, satw, nil)
}

// DeleteStockAddToWave deletes an existing stock.add.to.wave record.
func (c *Client) DeleteStockAddToWave(id int64) error {
	return c.DeleteStockAddToWaves([]int64{id})
}

// DeleteStockAddToWaves deletes existing stock.add.to.wave records.
func (c *Client) DeleteStockAddToWaves(ids []int64) error {
	return c.Delete(StockAddToWaveModel, ids)
}

// GetStockAddToWave gets stock.add.to.wave existing record.
func (c *Client) GetStockAddToWave(id int64) (*StockAddToWave, error) {
	satws, err := c.GetStockAddToWaves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*satws)[0]), nil
}

// GetStockAddToWaves gets stock.add.to.wave existing records.
func (c *Client) GetStockAddToWaves(ids []int64) (*StockAddToWaves, error) {
	satws := &StockAddToWaves{}
	if err := c.Read(StockAddToWaveModel, ids, nil, satws); err != nil {
		return nil, err
	}
	return satws, nil
}

// FindStockAddToWave finds stock.add.to.wave record by querying it with criteria.
func (c *Client) FindStockAddToWave(criteria *Criteria) (*StockAddToWave, error) {
	satws := &StockAddToWaves{}
	if err := c.SearchRead(StockAddToWaveModel, criteria, NewOptions().Limit(1), satws); err != nil {
		return nil, err
	}
	return &((*satws)[0]), nil
}

// FindStockAddToWaves finds stock.add.to.wave records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAddToWaves(criteria *Criteria, options *Options) (*StockAddToWaves, error) {
	satws := &StockAddToWaves{}
	if err := c.SearchRead(StockAddToWaveModel, criteria, options, satws); err != nil {
		return nil, err
	}
	return satws, nil
}

// FindStockAddToWaveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAddToWaveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockAddToWaveModel, criteria, options)
}

// FindStockAddToWaveId finds record id by querying it with criteria.
func (c *Client) FindStockAddToWaveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockAddToWaveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
