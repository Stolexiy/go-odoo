package odoo

// PosBill represents pos.bill model.
type PosBill struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	ForAllConfig *Bool     `xmlrpc:"for_all_config,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	PosConfigIds *Relation `xmlrpc:"pos_config_ids,omitempty"`
	Value        *Float    `xmlrpc:"value,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosBills represents array of pos.bill model.
type PosBills []PosBill

// PosBillModel is the odoo model name.
const PosBillModel = "pos.bill"

// Many2One convert PosBill to *Many2One.
func (pb *PosBill) Many2One() *Many2One {
	return NewMany2One(pb.Id.Get(), "")
}

// CreatePosBill creates a new pos.bill model and returns its id.
func (c *Client) CreatePosBill(pb *PosBill) (int64, error) {
	ids, err := c.CreatePosBills([]*PosBill{pb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosBill creates a new pos.bill model and returns its id.
func (c *Client) CreatePosBills(pbs []*PosBill) ([]int64, error) {
	var vv []interface{}
	for _, v := range pbs {
		vv = append(vv, v)
	}
	return c.Create(PosBillModel, vv, nil)
}

// UpdatePosBill updates an existing pos.bill record.
func (c *Client) UpdatePosBill(pb *PosBill) error {
	return c.UpdatePosBills([]int64{pb.Id.Get()}, pb)
}

// UpdatePosBills updates existing pos.bill records.
// All records (represented by ids) will be updated by pb values.
func (c *Client) UpdatePosBills(ids []int64, pb *PosBill) error {
	return c.Update(PosBillModel, ids, pb, nil)
}

// DeletePosBill deletes an existing pos.bill record.
func (c *Client) DeletePosBill(id int64) error {
	return c.DeletePosBills([]int64{id})
}

// DeletePosBills deletes existing pos.bill records.
func (c *Client) DeletePosBills(ids []int64) error {
	return c.Delete(PosBillModel, ids)
}

// GetPosBill gets pos.bill existing record.
func (c *Client) GetPosBill(id int64) (*PosBill, error) {
	pbs, err := c.GetPosBills([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pbs)[0]), nil
}

// GetPosBills gets pos.bill existing records.
func (c *Client) GetPosBills(ids []int64) (*PosBills, error) {
	pbs := &PosBills{}
	if err := c.Read(PosBillModel, ids, nil, pbs); err != nil {
		return nil, err
	}
	return pbs, nil
}

// FindPosBill finds pos.bill record by querying it with criteria.
func (c *Client) FindPosBill(criteria *Criteria) (*PosBill, error) {
	pbs := &PosBills{}
	if err := c.SearchRead(PosBillModel, criteria, NewOptions().Limit(1), pbs); err != nil {
		return nil, err
	}
	return &((*pbs)[0]), nil
}

// FindPosBills finds pos.bill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosBills(criteria *Criteria, options *Options) (*PosBills, error) {
	pbs := &PosBills{}
	if err := c.SearchRead(PosBillModel, criteria, options, pbs); err != nil {
		return nil, err
	}
	return pbs, nil
}

// FindPosBillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosBillIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosBillModel, criteria, options)
}

// FindPosBillId finds record id by querying it with criteria.
func (c *Client) FindPosBillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosBillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
