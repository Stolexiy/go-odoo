package odoo

// XCountry represents x_country model.
type XCountry struct {
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
	XName           *String   `xmlrpc:"x_name,omitempty"`
	XStudioSequence *Int      `xmlrpc:"x_studio_sequence,omitempty"`
}

// XCountrys represents array of x_country model.
type XCountrys []XCountry

// XCountryModel is the odoo model name.
const XCountryModel = "x_country"

// Many2One convert XCountry to *Many2One.
func (x *XCountry) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXCountry creates a new x_country model and returns its id.
func (c *Client) CreateXCountry(x *XCountry) (int64, error) {
	ids, err := c.CreateXCountrys([]*XCountry{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXCountry creates a new x_country model and returns its id.
func (c *Client) CreateXCountrys(xs []*XCountry) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XCountryModel, vv, nil)
}

// UpdateXCountry updates an existing x_country record.
func (c *Client) UpdateXCountry(x *XCountry) error {
	return c.UpdateXCountrys([]int64{x.Id.Get()}, x)
}

// UpdateXCountrys updates existing x_country records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXCountrys(ids []int64, x *XCountry) error {
	return c.Update(XCountryModel, ids, x, nil)
}

// DeleteXCountry deletes an existing x_country record.
func (c *Client) DeleteXCountry(id int64) error {
	return c.DeleteXCountrys([]int64{id})
}

// DeleteXCountrys deletes existing x_country records.
func (c *Client) DeleteXCountrys(ids []int64) error {
	return c.Delete(XCountryModel, ids)
}

// GetXCountry gets x_country existing record.
func (c *Client) GetXCountry(id int64) (*XCountry, error) {
	xs, err := c.GetXCountrys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXCountrys gets x_country existing records.
func (c *Client) GetXCountrys(ids []int64) (*XCountrys, error) {
	xs := &XCountrys{}
	if err := c.Read(XCountryModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXCountry finds x_country record by querying it with criteria.
func (c *Client) FindXCountry(criteria *Criteria) (*XCountry, error) {
	xs := &XCountrys{}
	if err := c.SearchRead(XCountryModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXCountrys finds x_country records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXCountrys(criteria *Criteria, options *Options) (*XCountrys, error) {
	xs := &XCountrys{}
	if err := c.SearchRead(XCountryModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXCountryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXCountryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XCountryModel, criteria, options)
}

// FindXCountryId finds record id by querying it with criteria.
func (c *Client) FindXCountryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XCountryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
