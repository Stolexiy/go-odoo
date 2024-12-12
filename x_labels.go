package odoo

// XLabels represents x_labels model.
type XLabels struct {
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
	XName            *String   `xmlrpc:"x_name,omitempty"`
	XStudioDiscogsId *Int      `xmlrpc:"x_studio_discogs_id,omitempty"`
	XStudioReleases  *Relation `xmlrpc:"x_studio_releases,omitempty"`
}

// XLabelss represents array of x_labels model.
type XLabelss []XLabels

// XLabelsModel is the odoo model name.
const XLabelsModel = "x_labels"

// Many2One convert XLabels to *Many2One.
func (x *XLabels) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXLabels creates a new x_labels model and returns its id.
func (c *Client) CreateXLabels(x *XLabels) (int64, error) {
	ids, err := c.CreateXLabelss([]*XLabels{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXLabels creates a new x_labels model and returns its id.
func (c *Client) CreateXLabelss(xs []*XLabels) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XLabelsModel, vv, nil)
}

// UpdateXLabels updates an existing x_labels record.
func (c *Client) UpdateXLabels(x *XLabels) error {
	return c.UpdateXLabelss([]int64{x.Id.Get()}, x)
}

// UpdateXLabelss updates existing x_labels records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXLabelss(ids []int64, x *XLabels) error {
	return c.Update(XLabelsModel, ids, x, nil)
}

// DeleteXLabels deletes an existing x_labels record.
func (c *Client) DeleteXLabels(id int64) error {
	return c.DeleteXLabelss([]int64{id})
}

// DeleteXLabelss deletes existing x_labels records.
func (c *Client) DeleteXLabelss(ids []int64) error {
	return c.Delete(XLabelsModel, ids)
}

// GetXLabels gets x_labels existing record.
func (c *Client) GetXLabels(id int64) (*XLabels, error) {
	xs, err := c.GetXLabelss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXLabelss gets x_labels existing records.
func (c *Client) GetXLabelss(ids []int64) (*XLabelss, error) {
	xs := &XLabelss{}
	if err := c.Read(XLabelsModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXLabels finds x_labels record by querying it with criteria.
func (c *Client) FindXLabels(criteria *Criteria) (*XLabels, error) {
	xs := &XLabelss{}
	if err := c.SearchRead(XLabelsModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXLabelss finds x_labels records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXLabelss(criteria *Criteria, options *Options) (*XLabelss, error) {
	xs := &XLabelss{}
	if err := c.SearchRead(XLabelsModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXLabelsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXLabelsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XLabelsModel, criteria, options)
}

// FindXLabelsId finds record id by querying it with criteria.
func (c *Client) FindXLabelsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XLabelsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}