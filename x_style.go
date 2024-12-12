package odoo

// XStyle represents x_style model.
type XStyle struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XName       *String   `xmlrpc:"x_name,omitempty"`
}

// XStyleS represents array of x_style model.
type XStyleS []XStyle

// XStyleModel is the odoo model name.
const XStyleModel = "x_style"

// Many2One convert XStyle to *Many2One.
func (x *XStyle) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXStyle creates a new x_style model and returns its id.
func (c *Client) CreateXStyle(x *XStyle) (int64, error) {
	ids, err := c.CreateXStyleS([]*XStyle{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXStyleS creates a new x_style model and returns its id.
func (c *Client) CreateXStyleS(xs []*XStyle) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XStyleModel, vv, nil)
}

// UpdateXStyle updates an existing x_style record.
func (c *Client) UpdateXStyle(x *XStyle) error {
	return c.UpdateXStyleS([]int64{x.Id.Get()}, x)
}

// UpdateXStyleS updates existing x_style records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXStyleS(ids []int64, x *XStyle) error {
	return c.Update(XStyleModel, ids, x, nil)
}

// DeleteXStyle deletes an existing x_style record.
func (c *Client) DeleteXStyle(id int64) error {
	return c.DeleteXStyleS([]int64{id})
}

// DeleteXStyleS deletes existing x_style records.
func (c *Client) DeleteXStyleS(ids []int64) error {
	return c.Delete(XStyleModel, ids)
}

// GetXStyle gets x_style existing record.
func (c *Client) GetXStyle(id int64) (*XStyle, error) {
	xs, err := c.GetXStyleS([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXStyleS gets x_style existing records.
func (c *Client) GetXStyleS(ids []int64) (*XStyleS, error) {
	xs := &XStyleS{}
	if err := c.Read(XStyleModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXStyle finds x_style record by querying it with criteria.
func (c *Client) FindXStyle(criteria *Criteria) (*XStyle, error) {
	xs := &XStyleS{}
	if err := c.SearchRead(XStyleModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXStyleS finds x_style records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXStyleS(criteria *Criteria, options *Options) (*XStyleS, error) {
	xs := &XStyleS{}
	if err := c.SearchRead(XStyleModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXStyleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXStyleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XStyleModel, criteria, options)
}

// FindXStyleId finds record id by querying it with criteria.
func (c *Client) FindXStyleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XStyleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
