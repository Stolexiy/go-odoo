package odoo

import (
	"fmt"
)

// XGenre represents x_genre model.
type XGenre struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
	XColor      *Int      `xmlrpc:"x_color,omptempty"`
	XName       *String   `xmlrpc:"x_name,omptempty"`
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
	return c.Create(XGenreModel, vv)
}

// UpdateXGenre updates an existing x_genre record.
func (c *Client) UpdateXGenre(x *XGenre) error {
	return c.UpdateXGenres([]int64{x.Id.Get()}, x)
}

// UpdateXGenres updates existing x_genre records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXGenres(ids []int64, x *XGenre) error {
	return c.Update(XGenreModel, ids, x)
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
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_genre not found", id)
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
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_genre was not found with criteria %v", criteria)
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
	ids, err := c.Search(XGenreModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXGenreId finds record id by querying it with criteria.
func (c *Client) FindXGenreId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XGenreModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_genre was not found with criteria %v and options %v", criteria, options)
}
