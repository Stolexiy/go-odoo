package odoo

// XTechnician represents x_technician model.
type XTechnician struct {
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
	XActive         *Bool     `xmlrpc:"x_active,omitempty"`
	XColor          *Int      `xmlrpc:"x_color,omitempty"`
	XName           *String   `xmlrpc:"x_name,omitempty"`
	XStudioSequence *Int      `xmlrpc:"x_studio_sequence,omitempty"`
}

// XTechnicians represents array of x_technician model.
type XTechnicians []XTechnician

// XTechnicianModel is the odoo model name.
const XTechnicianModel = "x_technician"

// Many2One convert XTechnician to *Many2One.
func (x *XTechnician) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXTechnician creates a new x_technician model and returns its id.
func (c *Client) CreateXTechnician(x *XTechnician) (int64, error) {
	ids, err := c.CreateXTechnicians([]*XTechnician{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXTechnician creates a new x_technician model and returns its id.
func (c *Client) CreateXTechnicians(xs []*XTechnician) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XTechnicianModel, vv, nil)
}

// UpdateXTechnician updates an existing x_technician record.
func (c *Client) UpdateXTechnician(x *XTechnician) error {
	return c.UpdateXTechnicians([]int64{x.Id.Get()}, x)
}

// UpdateXTechnicians updates existing x_technician records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXTechnicians(ids []int64, x *XTechnician) error {
	return c.Update(XTechnicianModel, ids, x, nil)
}

// DeleteXTechnician deletes an existing x_technician record.
func (c *Client) DeleteXTechnician(id int64) error {
	return c.DeleteXTechnicians([]int64{id})
}

// DeleteXTechnicians deletes existing x_technician records.
func (c *Client) DeleteXTechnicians(ids []int64) error {
	return c.Delete(XTechnicianModel, ids)
}

// GetXTechnician gets x_technician existing record.
func (c *Client) GetXTechnician(id int64) (*XTechnician, error) {
	xs, err := c.GetXTechnicians([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXTechnicians gets x_technician existing records.
func (c *Client) GetXTechnicians(ids []int64) (*XTechnicians, error) {
	xs := &XTechnicians{}
	if err := c.Read(XTechnicianModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXTechnician finds x_technician record by querying it with criteria.
func (c *Client) FindXTechnician(criteria *Criteria) (*XTechnician, error) {
	xs := &XTechnicians{}
	if err := c.SearchRead(XTechnicianModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXTechnicians finds x_technician records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXTechnicians(criteria *Criteria, options *Options) (*XTechnicians, error) {
	xs := &XTechnicians{}
	if err := c.SearchRead(XTechnicianModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXTechnicianIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXTechnicianIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XTechnicianModel, criteria, options)
}

// FindXTechnicianId finds record id by querying it with criteria.
func (c *Client) FindXTechnicianId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XTechnicianModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
