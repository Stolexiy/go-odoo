package odoo

import (
	"fmt"
)

// XCollection represents x_collection model.
type XCollection struct {
	LastUpdate                      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                     *String   `xmlrpc:"display_name,omptempty"`
	Id                              *Int      `xmlrpc:"id,omptempty"`
	WriteDate                       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                        *Many2One `xmlrpc:"write_uid,omptempty"`
	XName                           *String   `xmlrpc:"x_name,omptempty"`
	XStudioDiscogsFolderId          *Int      `xmlrpc:"x_studio_discogs_folder_id,omptempty"`
	XStudioNombreDeVinyles          *String   `xmlrpc:"x_studio_nombre_de_vinyles,omptempty"`
	XStudioPrixDachatDeLaCollection *String   `xmlrpc:"x_studio_prix_dachat_de_la_collection,omptempty"`
}

// XCollections represents array of x_collection model.
type XCollections []XCollection

// XCollectionModel is the odoo model name.
const XCollectionModel = "x_collection"

// Many2One convert XCollection to *Many2One.
func (x *XCollection) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXCollection creates a new x_collection model and returns its id.
func (c *Client) CreateXCollection(x *XCollection) (int64, error) {
	ids, err := c.CreateXCollections([]*XCollection{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXCollection creates a new x_collection model and returns its id.
func (c *Client) CreateXCollections(xs []*XCollection) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XCollectionModel, vv)
}

// UpdateXCollection updates an existing x_collection record.
func (c *Client) UpdateXCollection(x *XCollection) error {
	return c.UpdateXCollections([]int64{x.Id.Get()}, x)
}

// UpdateXCollections updates existing x_collection records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXCollections(ids []int64, x *XCollection) error {
	return c.Update(XCollectionModel, ids, x)
}

// DeleteXCollection deletes an existing x_collection record.
func (c *Client) DeleteXCollection(id int64) error {
	return c.DeleteXCollections([]int64{id})
}

// DeleteXCollections deletes existing x_collection records.
func (c *Client) DeleteXCollections(ids []int64) error {
	return c.Delete(XCollectionModel, ids)
}

// GetXCollection gets x_collection existing record.
func (c *Client) GetXCollection(id int64) (*XCollection, error) {
	xs, err := c.GetXCollections([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_collection not found", id)
}

// GetXCollections gets x_collection existing records.
func (c *Client) GetXCollections(ids []int64) (*XCollections, error) {
	xs := &XCollections{}
	if err := c.Read(XCollectionModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXCollection finds x_collection record by querying it with criteria.
func (c *Client) FindXCollection(criteria *Criteria) (*XCollection, error) {
	xs := &XCollections{}
	if err := c.SearchRead(XCollectionModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_collection was not found with criteria %v", criteria)
}

// FindXCollections finds x_collection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXCollections(criteria *Criteria, options *Options) (*XCollections, error) {
	xs := &XCollections{}
	if err := c.SearchRead(XCollectionModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXCollectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXCollectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XCollectionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXCollectionId finds record id by querying it with criteria.
func (c *Client) FindXCollectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XCollectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_collection was not found with criteria %v and options %v", criteria, options)
}
