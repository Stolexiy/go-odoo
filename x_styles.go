package odoo

import (
	"fmt"
)

// XStyles represents x_styles model.
type XStyles struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
	XName       *String   `xmlrpc:"x_name,omptempty"`
}

// XStyless represents array of x_styles model.
type XStyless []XStyles

// XStylesModel is the odoo model name.
const XStylesModel = "x_styles"

// Many2One convert XStyles to *Many2One.
func (x *XStyles) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXStyles creates a new x_styles model and returns its id.
func (c *Client) CreateXStyles(x *XStyles) (int64, error) {
	ids, err := c.CreateXStyless([]*XStyles{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXStyles creates a new x_styles model and returns its id.
func (c *Client) CreateXStyless(xs []*XStyles) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XStylesModel, vv)
}

// UpdateXStyles updates an existing x_styles record.
func (c *Client) UpdateXStyles(x *XStyles) error {
	return c.UpdateXStyless([]int64{x.Id.Get()}, x)
}

// UpdateXStyless updates existing x_styles records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXStyless(ids []int64, x *XStyles) error {
	return c.Update(XStylesModel, ids, x)
}

// DeleteXStyles deletes an existing x_styles record.
func (c *Client) DeleteXStyles(id int64) error {
	return c.DeleteXStyless([]int64{id})
}

// DeleteXStyless deletes existing x_styles records.
func (c *Client) DeleteXStyless(ids []int64) error {
	return c.Delete(XStylesModel, ids)
}

// GetXStyles gets x_styles existing record.
func (c *Client) GetXStyles(id int64) (*XStyles, error) {
	xs, err := c.GetXStyless([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_styles not found", id)
}

// GetXStyless gets x_styles existing records.
func (c *Client) GetXStyless(ids []int64) (*XStyless, error) {
	xs := &XStyless{}
	if err := c.Read(XStylesModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXStyles finds x_styles record by querying it with criteria.
func (c *Client) FindXStyles(criteria *Criteria) (*XStyles, error) {
	xs := &XStyless{}
	if err := c.SearchRead(XStylesModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_styles was not found with criteria %v", criteria)
}

// FindXStyless finds x_styles records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXStyless(criteria *Criteria, options *Options) (*XStyless, error) {
	xs := &XStyless{}
	if err := c.SearchRead(XStylesModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXStylesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXStylesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XStylesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXStylesId finds record id by querying it with criteria.
func (c *Client) FindXStylesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XStylesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_styles was not found with criteria %v and options %v", criteria, options)
}
