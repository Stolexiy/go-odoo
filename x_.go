package odoo

// X represents x_ model.
type X struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XName       *String   `xmlrpc:"x_name,omitempty"`
}

// Xs represents array of x_ model.
type Xs []X

// XModel is the odoo model name.
const XModel = "x_"

// Many2One convert X to *Many2One.
func (x *X) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateX creates a new x_ model and returns its id.
func (c *Client) CreateX(x *X) (int64, error) {
	ids, err := c.CreateXs([]*X{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateX creates a new x_ model and returns its id.
func (c *Client) CreateXs(xs []*X) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XModel, vv, nil)
}

// UpdateX updates an existing x_ record.
func (c *Client) UpdateX(x *X) error {
	return c.UpdateXs([]int64{x.Id.Get()}, x)
}

// UpdateXs updates existing x_ records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXs(ids []int64, x *X) error {
	return c.Update(XModel, ids, x, nil)
}

// DeleteX deletes an existing x_ record.
func (c *Client) DeleteX(id int64) error {
	return c.DeleteXs([]int64{id})
}

// DeleteXs deletes existing x_ records.
func (c *Client) DeleteXs(ids []int64) error {
	return c.Delete(XModel, ids)
}

// GetX gets x_ existing record.
func (c *Client) GetX(id int64) (*X, error) {
	xs, err := c.GetXs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXs gets x_ existing records.
func (c *Client) GetXs(ids []int64) (*Xs, error) {
	xs := &Xs{}
	if err := c.Read(XModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindX finds x_ record by querying it with criteria.
func (c *Client) FindX(criteria *Criteria) (*X, error) {
	xs := &Xs{}
	if err := c.SearchRead(XModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXs finds x_ records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXs(criteria *Criteria, options *Options) (*Xs, error) {
	xs := &Xs{}
	if err := c.SearchRead(XModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XModel, criteria, options)
}

// FindXId finds record id by querying it with criteria.
func (c *Client) FindXId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
