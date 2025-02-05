package odoo

import (
	"fmt"
)

// XRelease represents x_release model.
type XRelease struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
	XName           *String   `xmlrpc:"x_name,omptempty"`
	XStudioLabel    *Many2One `xmlrpc:"x_studio_label,omptempty"`
	XStudioSequence *Int      `xmlrpc:"x_studio_sequence,omptempty"`
}

// XReleases represents array of x_release model.
type XReleases []XRelease

// XReleaseModel is the odoo model name.
const XReleaseModel = "x_release"

// Many2One convert XRelease to *Many2One.
func (x *XRelease) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXRelease creates a new x_release model and returns its id.
func (c *Client) CreateXRelease(x *XRelease) (int64, error) {
	ids, err := c.CreateXReleases([]*XRelease{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXRelease creates a new x_release model and returns its id.
func (c *Client) CreateXReleases(xs []*XRelease) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XReleaseModel, vv)
}

// UpdateXRelease updates an existing x_release record.
func (c *Client) UpdateXRelease(x *XRelease) error {
	return c.UpdateXReleases([]int64{x.Id.Get()}, x)
}

// UpdateXReleases updates existing x_release records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXReleases(ids []int64, x *XRelease) error {
	return c.Update(XReleaseModel, ids, x)
}

// DeleteXRelease deletes an existing x_release record.
func (c *Client) DeleteXRelease(id int64) error {
	return c.DeleteXReleases([]int64{id})
}

// DeleteXReleases deletes existing x_release records.
func (c *Client) DeleteXReleases(ids []int64) error {
	return c.Delete(XReleaseModel, ids)
}

// GetXRelease gets x_release existing record.
func (c *Client) GetXRelease(id int64) (*XRelease, error) {
	xs, err := c.GetXReleases([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_release not found", id)
}

// GetXReleases gets x_release existing records.
func (c *Client) GetXReleases(ids []int64) (*XReleases, error) {
	xs := &XReleases{}
	if err := c.Read(XReleaseModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXRelease finds x_release record by querying it with criteria.
func (c *Client) FindXRelease(criteria *Criteria) (*XRelease, error) {
	xs := &XReleases{}
	if err := c.SearchRead(XReleaseModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_release was not found with criteria %v", criteria)
}

// FindXReleases finds x_release records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXReleases(criteria *Criteria, options *Options) (*XReleases, error) {
	xs := &XReleases{}
	if err := c.SearchRead(XReleaseModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXReleaseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXReleaseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XReleaseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXReleaseId finds record id by querying it with criteria.
func (c *Client) FindXReleaseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XReleaseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_release was not found with criteria %v and options %v", criteria, options)
}
