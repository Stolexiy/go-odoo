package odoo

// ProductMargin represents product.margin model.
type ProductMargin struct {
	CreateDate   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	FromDate     *Time      `xmlrpc:"from_date,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	InvoiceState *Selection `xmlrpc:"invoice_state,omitempty"`
	ToDate       *Time      `xmlrpc:"to_date,omitempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductMargins represents array of product.margin model.
type ProductMargins []ProductMargin

// ProductMarginModel is the odoo model name.
const ProductMarginModel = "product.margin"

// Many2One convert ProductMargin to *Many2One.
func (pm *ProductMargin) Many2One() *Many2One {
	return NewMany2One(pm.Id.Get(), "")
}

// CreateProductMargin creates a new product.margin model and returns its id.
func (c *Client) CreateProductMargin(pm *ProductMargin) (int64, error) {
	ids, err := c.CreateProductMargins([]*ProductMargin{pm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductMargin creates a new product.margin model and returns its id.
func (c *Client) CreateProductMargins(pms []*ProductMargin) ([]int64, error) {
	var vv []interface{}
	for _, v := range pms {
		vv = append(vv, v)
	}
	return c.Create(ProductMarginModel, vv, nil)
}

// UpdateProductMargin updates an existing product.margin record.
func (c *Client) UpdateProductMargin(pm *ProductMargin) error {
	return c.UpdateProductMargins([]int64{pm.Id.Get()}, pm)
}

// UpdateProductMargins updates existing product.margin records.
// All records (represented by ids) will be updated by pm values.
func (c *Client) UpdateProductMargins(ids []int64, pm *ProductMargin) error {
	return c.Update(ProductMarginModel, ids, pm, nil)
}

// DeleteProductMargin deletes an existing product.margin record.
func (c *Client) DeleteProductMargin(id int64) error {
	return c.DeleteProductMargins([]int64{id})
}

// DeleteProductMargins deletes existing product.margin records.
func (c *Client) DeleteProductMargins(ids []int64) error {
	return c.Delete(ProductMarginModel, ids)
}

// GetProductMargin gets product.margin existing record.
func (c *Client) GetProductMargin(id int64) (*ProductMargin, error) {
	pms, err := c.GetProductMargins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pms)[0]), nil
}

// GetProductMargins gets product.margin existing records.
func (c *Client) GetProductMargins(ids []int64) (*ProductMargins, error) {
	pms := &ProductMargins{}
	if err := c.Read(ProductMarginModel, ids, nil, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindProductMargin finds product.margin record by querying it with criteria.
func (c *Client) FindProductMargin(criteria *Criteria) (*ProductMargin, error) {
	pms := &ProductMargins{}
	if err := c.SearchRead(ProductMarginModel, criteria, NewOptions().Limit(1), pms); err != nil {
		return nil, err
	}
	return &((*pms)[0]), nil
}

// FindProductMargins finds product.margin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductMargins(criteria *Criteria, options *Options) (*ProductMargins, error) {
	pms := &ProductMargins{}
	if err := c.SearchRead(ProductMarginModel, criteria, options, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindProductMarginIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductMarginIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductMarginModel, criteria, options)
}

// FindProductMarginId finds record id by querying it with criteria.
func (c *Client) FindProductMarginId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductMarginModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
