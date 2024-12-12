package odoo

// PosPackOperationLot represents pos.pack.operation.lot model.
type PosPackOperationLot struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	LotName        *String   `xmlrpc:"lot_name,omitempty"`
	OrderId        *Many2One `xmlrpc:"order_id,omitempty"`
	PosOrderLineId *Many2One `xmlrpc:"pos_order_line_id,omitempty"`
	ProductId      *Many2One `xmlrpc:"product_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPackOperationLots represents array of pos.pack.operation.lot model.
type PosPackOperationLots []PosPackOperationLot

// PosPackOperationLotModel is the odoo model name.
const PosPackOperationLotModel = "pos.pack.operation.lot"

// Many2One convert PosPackOperationLot to *Many2One.
func (ppol *PosPackOperationLot) Many2One() *Many2One {
	return NewMany2One(ppol.Id.Get(), "")
}

// CreatePosPackOperationLot creates a new pos.pack.operation.lot model and returns its id.
func (c *Client) CreatePosPackOperationLot(ppol *PosPackOperationLot) (int64, error) {
	ids, err := c.CreatePosPackOperationLots([]*PosPackOperationLot{ppol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPackOperationLot creates a new pos.pack.operation.lot model and returns its id.
func (c *Client) CreatePosPackOperationLots(ppols []*PosPackOperationLot) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppols {
		vv = append(vv, v)
	}
	return c.Create(PosPackOperationLotModel, vv, nil)
}

// UpdatePosPackOperationLot updates an existing pos.pack.operation.lot record.
func (c *Client) UpdatePosPackOperationLot(ppol *PosPackOperationLot) error {
	return c.UpdatePosPackOperationLots([]int64{ppol.Id.Get()}, ppol)
}

// UpdatePosPackOperationLots updates existing pos.pack.operation.lot records.
// All records (represented by ids) will be updated by ppol values.
func (c *Client) UpdatePosPackOperationLots(ids []int64, ppol *PosPackOperationLot) error {
	return c.Update(PosPackOperationLotModel, ids, ppol, nil)
}

// DeletePosPackOperationLot deletes an existing pos.pack.operation.lot record.
func (c *Client) DeletePosPackOperationLot(id int64) error {
	return c.DeletePosPackOperationLots([]int64{id})
}

// DeletePosPackOperationLots deletes existing pos.pack.operation.lot records.
func (c *Client) DeletePosPackOperationLots(ids []int64) error {
	return c.Delete(PosPackOperationLotModel, ids)
}

// GetPosPackOperationLot gets pos.pack.operation.lot existing record.
func (c *Client) GetPosPackOperationLot(id int64) (*PosPackOperationLot, error) {
	ppols, err := c.GetPosPackOperationLots([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ppols)[0]), nil
}

// GetPosPackOperationLots gets pos.pack.operation.lot existing records.
func (c *Client) GetPosPackOperationLots(ids []int64) (*PosPackOperationLots, error) {
	ppols := &PosPackOperationLots{}
	if err := c.Read(PosPackOperationLotModel, ids, nil, ppols); err != nil {
		return nil, err
	}
	return ppols, nil
}

// FindPosPackOperationLot finds pos.pack.operation.lot record by querying it with criteria.
func (c *Client) FindPosPackOperationLot(criteria *Criteria) (*PosPackOperationLot, error) {
	ppols := &PosPackOperationLots{}
	if err := c.SearchRead(PosPackOperationLotModel, criteria, NewOptions().Limit(1), ppols); err != nil {
		return nil, err
	}
	return &((*ppols)[0]), nil
}

// FindPosPackOperationLots finds pos.pack.operation.lot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPackOperationLots(criteria *Criteria, options *Options) (*PosPackOperationLots, error) {
	ppols := &PosPackOperationLots{}
	if err := c.SearchRead(PosPackOperationLotModel, criteria, options, ppols); err != nil {
		return nil, err
	}
	return ppols, nil
}

// FindPosPackOperationLotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPackOperationLotIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPackOperationLotModel, criteria, options)
}

// FindPosPackOperationLotId finds record id by querying it with criteria.
func (c *Client) FindPosPackOperationLotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPackOperationLotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
