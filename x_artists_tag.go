package odoo

// XArtistsTag represents x_artists_tag model.
type XArtistsTag struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XColor      *Int      `xmlrpc:"x_color,omitempty"`
	XName       *String   `xmlrpc:"x_name,omitempty"`
}

// XArtistsTags represents array of x_artists_tag model.
type XArtistsTags []XArtistsTag

// XArtistsTagModel is the odoo model name.
const XArtistsTagModel = "x_artists_tag"

// Many2One convert XArtistsTag to *Many2One.
func (x *XArtistsTag) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXArtistsTag creates a new x_artists_tag model and returns its id.
func (c *Client) CreateXArtistsTag(x *XArtistsTag) (int64, error) {
	ids, err := c.CreateXArtistsTags([]*XArtistsTag{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXArtistsTag creates a new x_artists_tag model and returns its id.
func (c *Client) CreateXArtistsTags(xs []*XArtistsTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XArtistsTagModel, vv, nil)
}

// UpdateXArtistsTag updates an existing x_artists_tag record.
func (c *Client) UpdateXArtistsTag(x *XArtistsTag) error {
	return c.UpdateXArtistsTags([]int64{x.Id.Get()}, x)
}

// UpdateXArtistsTags updates existing x_artists_tag records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXArtistsTags(ids []int64, x *XArtistsTag) error {
	return c.Update(XArtistsTagModel, ids, x, nil)
}

// DeleteXArtistsTag deletes an existing x_artists_tag record.
func (c *Client) DeleteXArtistsTag(id int64) error {
	return c.DeleteXArtistsTags([]int64{id})
}

// DeleteXArtistsTags deletes existing x_artists_tag records.
func (c *Client) DeleteXArtistsTags(ids []int64) error {
	return c.Delete(XArtistsTagModel, ids)
}

// GetXArtistsTag gets x_artists_tag existing record.
func (c *Client) GetXArtistsTag(id int64) (*XArtistsTag, error) {
	xs, err := c.GetXArtistsTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// GetXArtistsTags gets x_artists_tag existing records.
func (c *Client) GetXArtistsTags(ids []int64) (*XArtistsTags, error) {
	xs := &XArtistsTags{}
	if err := c.Read(XArtistsTagModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXArtistsTag finds x_artists_tag record by querying it with criteria.
func (c *Client) FindXArtistsTag(criteria *Criteria) (*XArtistsTag, error) {
	xs := &XArtistsTags{}
	if err := c.SearchRead(XArtistsTagModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	return &((*xs)[0]), nil
}

// FindXArtistsTags finds x_artists_tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXArtistsTags(criteria *Criteria, options *Options) (*XArtistsTags, error) {
	xs := &XArtistsTags{}
	if err := c.SearchRead(XArtistsTagModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXArtistsTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXArtistsTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(XArtistsTagModel, criteria, options)
}

// FindXArtistsTagId finds record id by querying it with criteria.
func (c *Client) FindXArtistsTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XArtistsTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
