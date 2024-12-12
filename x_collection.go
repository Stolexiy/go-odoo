package odoo

// XCollection represents x_collection model.
type XCollection struct {
	CreateDate                      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                     *String   `xmlrpc:"display_name,omitempty"`
	Id                              *Int      `xmlrpc:"id,omitempty"`
	WriteDate                       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                        *Many2One `xmlrpc:"write_uid,omitempty"`
	XName                           *String   `xmlrpc:"x_name,omitempty"`
	XStudioDiscogsFolderId          *Int      `xmlrpc:"x_studio_discogs_folder_id,omitempty"`
	XStudioNombreDeVinyles          *String   `xmlrpc:"x_studio_nombre_de_vinyles,omitempty"`
	XStudioPrixDachatDeLaCollection *String   `xmlrpc:"x_studio_prix_dachat_de_la_collection,omitempty"`
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
	return c.Create(XCollectionModel, vv, nil)
}

// UpdateXCollection updates an existing x_collection record.
func (c *Client) UpdateXCollection(x *XCollection) error {
	return c.UpdateXCollections([]int64{x.Id.Get()}, x)
}

// UpdateXCollections updates existing x_collection records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXCollections(ids []int64, x *XCollection) error {
	return c.Update(XCollectionModel, ids, x, nil)
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
	return &((*xs)[0]), nil
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
	return &((*xs)[0]), nil
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
	return c.Search(XCollectionModel, criteria, options)
}

// FindXCollectionId finds record id by querying it with criteria.
func (c *Client) FindXCollectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XCollectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
