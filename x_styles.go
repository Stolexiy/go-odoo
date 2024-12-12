package odoo

// XStyles represents x_styles model.
type XStyles struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XName       *String   `xmlrpc:"x_name,omitempty"`
}

// XStyless represents array of x_styles model.
type XStyless []XStyles

// XStylesModel is the odoo model name.
const XStylesModel = "x_styles"

// Many2One convert XStyles to *Many2One.
func (x *XStyles) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXStyles creates a new x_styles model and returns its id.
func (c *Client) CreateXStyles(x *XStyles) (int64, error) {
	ids, err := c.CreateXStyless([]*XStyles{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXStyles creates a new x_styles model and returns its id.
func (c *Client) CreateXStyless(xs []*XStyles) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XStylesModel, vv, nil)
}

// UpdateXStyles updates an existing x_styles record.
func (c *Client) UpdateXStyles(x *XStyles) error {
	return c.UpdateXStyless([]int64{x.Id.Get()}, x)
}

// UpdateXStyless updates existing x_styles records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXStyless(ids []int64, x *XStyles) error {
	return c.Update(XStylesModel, ids, x, nil)
}

// DeleteXStyles deletes an existing x_styles record.
func (c *Client) DeleteXStyles(id int64) error {
	return c.DeleteXStyless([]int64{id})
}

// DeleteXStyless deletes existing x_styles records.
func (c *Client) DeleteXStyless(ids []int64) error {
	return c.Delete(XStylesModel, ids)
}

// GetXStyles gets x_styles existing record.
func (c *Client) GetXStyles(id int64) (*XStyles, error) {
	xs, err := c.GetXStyless([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXStyless gets x_styles existing records.
func (c *Client) GetXStyless(ids []int64) (*XStyless, error) {
	xs := &XStyless{}
	if err := c.Read(XStylesModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXStyles finds x_styles record by querying it with criteria.
func (c *Client) FindXStyles(criteria *Criteria) (*XStyles, error) {
	xs := &XStyless{}
	if err := c.SearchRead(XStylesModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXStyless finds x_styles records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXStyless(criteria *Criteria, options *Options) (*XStyless, error) {
	xs := &XStyless{}
	if err := c.SearchRead(XStylesModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXStylesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXStylesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XStylesModel, criteria, options)
}

// FindXStylesId finds record id by querying it with criteria.
func (c *Client) FindXStylesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XStylesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
