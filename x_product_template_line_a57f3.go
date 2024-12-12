package odoo

// XProductTemplateLineA57F3 represents x_product_template_line_a57f3 model.
type XProductTemplateLineA57F3 struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
	XName              *String   `xmlrpc:"x_name,omitempty"`
	XProductTemplateId *Many2One `xmlrpc:"x_product_template_id,omitempty"`
	XStudioSequence    *Int      `xmlrpc:"x_studio_sequence,omitempty"`
}

// XProductTemplateLineA57F3s represents array of x_product_template_line_a57f3 model.
type XProductTemplateLineA57F3s []XProductTemplateLineA57F3

// XProductTemplateLineA57F3Model is the odoo model name.
const XProductTemplateLineA57F3Model = "x_product_template_line_a57f3"

// Many2One convert XProductTemplateLineA57F3 to *Many2One.
func (x *XProductTemplateLineA57F3) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXProductTemplateLineA57F3 creates a new x_product_template_line_a57f3 model and returns its id.
func (c *Client) CreateXProductTemplateLineA57F3(x *XProductTemplateLineA57F3) (int64, error) {
	ids, err := c.CreateXProductTemplateLineA57F3s([]*XProductTemplateLineA57F3{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXProductTemplateLineA57F3 creates a new x_product_template_line_a57f3 model and returns its id.
func (c *Client) CreateXProductTemplateLineA57F3s(xs []*XProductTemplateLineA57F3) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XProductTemplateLineA57F3Model, vv, nil)
}

// UpdateXProductTemplateLineA57F3 updates an existing x_product_template_line_a57f3 record.
func (c *Client) UpdateXProductTemplateLineA57F3(x *XProductTemplateLineA57F3) error {
	return c.UpdateXProductTemplateLineA57F3s([]int64{x.Id.Get()}, x)
}

// UpdateXProductTemplateLineA57F3s updates existing x_product_template_line_a57f3 records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXProductTemplateLineA57F3s(ids []int64, x *XProductTemplateLineA57F3) error {
	return c.Update(XProductTemplateLineA57F3Model, ids, x, nil)
}

// DeleteXProductTemplateLineA57F3 deletes an existing x_product_template_line_a57f3 record.
func (c *Client) DeleteXProductTemplateLineA57F3(id int64) error {
	return c.DeleteXProductTemplateLineA57F3s([]int64{id})
}

// DeleteXProductTemplateLineA57F3s deletes existing x_product_template_line_a57f3 records.
func (c *Client) DeleteXProductTemplateLineA57F3s(ids []int64) error {
	return c.Delete(XProductTemplateLineA57F3Model, ids)
}

// GetXProductTemplateLineA57F3 gets x_product_template_line_a57f3 existing record.
func (c *Client) GetXProductTemplateLineA57F3(id int64) (*XProductTemplateLineA57F3, error) {
	xs, err := c.GetXProductTemplateLineA57F3s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXProductTemplateLineA57F3s gets x_product_template_line_a57f3 existing records.
func (c *Client) GetXProductTemplateLineA57F3s(ids []int64) (*XProductTemplateLineA57F3s, error) {
	xs := &XProductTemplateLineA57F3s{}
	if err := c.Read(XProductTemplateLineA57F3Model, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXProductTemplateLineA57F3 finds x_product_template_line_a57f3 record by querying it with criteria.
func (c *Client) FindXProductTemplateLineA57F3(criteria *Criteria) (*XProductTemplateLineA57F3, error) {
	xs := &XProductTemplateLineA57F3s{}
	if err := c.SearchRead(XProductTemplateLineA57F3Model, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXProductTemplateLineA57F3s finds x_product_template_line_a57f3 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXProductTemplateLineA57F3s(criteria *Criteria, options *Options) (*XProductTemplateLineA57F3s, error) {
	xs := &XProductTemplateLineA57F3s{}
	if err := c.SearchRead(XProductTemplateLineA57F3Model, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXProductTemplateLineA57F3Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXProductTemplateLineA57F3Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XProductTemplateLineA57F3Model, criteria, options)
}

// FindXProductTemplateLineA57F3Id finds record id by querying it with criteria.
func (c *Client) FindXProductTemplateLineA57F3Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XProductTemplateLineA57F3Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
