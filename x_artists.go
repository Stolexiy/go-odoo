package odoo

import (
	"fmt"
)

// XArtists represents x_artists model.
type XArtists struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
	XActive         *Bool     `xmlrpc:"x_active,omptempty"`
	XName           *String   `xmlrpc:"x_name,omptempty"`
	XStudioSequence *Int      `xmlrpc:"x_studio_sequence,omptempty"`
	XStudioTagIds   *Relation `xmlrpc:"x_studio_tag_ids,omptempty"`
}

// XArtistss represents array of x_artists model.
type XArtistss []XArtists

// XArtistsModel is the odoo model name.
const XArtistsModel = "x_artists"

// Many2One convert XArtists to *Many2One.
func (x *XArtists) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXArtists creates a new x_artists model and returns its id.
func (c *Client) CreateXArtists(x *XArtists) (int64, error) {
	ids, err := c.CreateXArtistss([]*XArtists{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXArtists creates a new x_artists model and returns its id.
func (c *Client) CreateXArtistss(xs []*XArtists) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XArtistsModel, vv)
}

// UpdateXArtists updates an existing x_artists record.
func (c *Client) UpdateXArtists(x *XArtists) error {
	return c.UpdateXArtistss([]int64{x.Id.Get()}, x)
}

// UpdateXArtistss updates existing x_artists records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXArtistss(ids []int64, x *XArtists) error {
	return c.Update(XArtistsModel, ids, x)
}

// DeleteXArtists deletes an existing x_artists record.
func (c *Client) DeleteXArtists(id int64) error {
	return c.DeleteXArtistss([]int64{id})
}

// DeleteXArtistss deletes existing x_artists records.
func (c *Client) DeleteXArtistss(ids []int64) error {
	return c.Delete(XArtistsModel, ids)
}

// GetXArtists gets x_artists existing record.
func (c *Client) GetXArtists(id int64) (*XArtists, error) {
	xs, err := c.GetXArtistss([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_artists not found", id)
}

// GetXArtistss gets x_artists existing records.
func (c *Client) GetXArtistss(ids []int64) (*XArtistss, error) {
	xs := &XArtistss{}
	if err := c.Read(XArtistsModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXArtists finds x_artists record by querying it with criteria.
func (c *Client) FindXArtists(criteria *Criteria) (*XArtists, error) {
	xs := &XArtistss{}
	if err := c.SearchRead(XArtistsModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_artists was not found with criteria %v", criteria)
}

// FindXArtistss finds x_artists records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXArtistss(criteria *Criteria, options *Options) (*XArtistss, error) {
	xs := &XArtistss{}
	if err := c.SearchRead(XArtistsModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXArtistsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXArtistsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XArtistsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXArtistsId finds record id by querying it with criteria.
func (c *Client) FindXArtistsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XArtistsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_artists was not found with criteria %v and options %v", criteria, options)
}
