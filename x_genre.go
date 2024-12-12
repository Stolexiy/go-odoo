package odoo

// XGenre represents x_genre model.
type XGenre struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XColor      *Int      `xmlrpc:"x_color,omitempty"`
	XName       *String   `xmlrpc:"x_name,omitempty"`
}

// XGenres represents array of x_genre model.
type XGenres []XGenre

// XGenreModel is the odoo model name.
const XGenreModel = "x_genre"

// Many2One convert XGenre to *Many2One.
func (x *XGenre) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXGenre creates a new x_genre model and returns its id.
func (c *Client) CreateXGenre(x *XGenre) (int64, error) {
	ids, err := c.CreateXGenres([]*XGenre{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXGenre creates a new x_genre model and returns its id.
func (c *Client) CreateXGenres(xs []*XGenre) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XGenreModel, vv, nil)
}

// UpdateXGenre updates an existing x_genre record.
func (c *Client) UpdateXGenre(x *XGenre) error {
	return c.UpdateXGenres([]int64{x.Id.Get()}, x)
}

// UpdateXGenres updates existing x_genre records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXGenres(ids []int64, x *XGenre) error {
	return c.Update(XGenreModel, ids, x, nil)
}

// DeleteXGenre deletes an existing x_genre record.
func (c *Client) DeleteXGenre(id int64) error {
	return c.DeleteXGenres([]int64{id})
}

// DeleteXGenres deletes existing x_genre records.
func (c *Client) DeleteXGenres(ids []int64) error {
	return c.Delete(XGenreModel, ids)
}

// GetXGenre gets x_genre existing record.
func (c *Client) GetXGenre(id int64) (*XGenre, error) {
	xs, err := c.GetXGenres([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXGenres gets x_genre existing records.
func (c *Client) GetXGenres(ids []int64) (*XGenres, error) {
	xs := &XGenres{}
	if err := c.Read(XGenreModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXGenre finds x_genre record by querying it with criteria.
func (c *Client) FindXGenre(criteria *Criteria) (*XGenre, error) {
	xs := &XGenres{}
	if err := c.SearchRead(XGenreModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXGenres finds x_genre records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXGenres(criteria *Criteria, options *Options) (*XGenres, error) {
	xs := &XGenres{}
	if err := c.SearchRead(XGenreModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXGenreIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXGenreIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XGenreModel, criteria, options)
}

// FindXGenreId finds record id by querying it with criteria.
func (c *Client) FindXGenreId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XGenreModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
