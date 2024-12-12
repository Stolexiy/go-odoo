package odoo

// SaleOrderSpreadsheet represents sale.order.spreadsheet model.
type SaleOrderSpreadsheet struct {
	CompanyId              *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrentRevisionUuid    *String   `xmlrpc:"current_revision_uuid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	Name                   *String   `xmlrpc:"name,omitempty"`
	OrderId                *Many2One `xmlrpc:"order_id,omitempty"`
	SpreadsheetBinaryData  *String   `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData        *String   `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName    *String   `xmlrpc:"spreadsheet_file_name,omitempty"`
	SpreadsheetRevisionIds *Relation `xmlrpc:"spreadsheet_revision_ids,omitempty"`
	SpreadsheetSnapshot    *String   `xmlrpc:"spreadsheet_snapshot,omitempty"`
	Thumbnail              *String   `xmlrpc:"thumbnail,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderSpreadsheets represents array of sale.order.spreadsheet model.
type SaleOrderSpreadsheets []SaleOrderSpreadsheet

// SaleOrderSpreadsheetModel is the odoo model name.
const SaleOrderSpreadsheetModel = "sale.order.spreadsheet"

// Many2One convert SaleOrderSpreadsheet to *Many2One.
func (sos *SaleOrderSpreadsheet) Many2One() *Many2One {
	return NewMany2One(sos.Id.Get(), "")
}

// CreateSaleOrderSpreadsheet creates a new sale.order.spreadsheet model and returns its id.
func (c *Client) CreateSaleOrderSpreadsheet(sos *SaleOrderSpreadsheet) (int64, error) {
	ids, err := c.CreateSaleOrderSpreadsheets([]*SaleOrderSpreadsheet{sos})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderSpreadsheet creates a new sale.order.spreadsheet model and returns its id.
func (c *Client) CreateSaleOrderSpreadsheets(soss []*SaleOrderSpreadsheet) ([]int64, error) {
	var vv []interface{}
	for _, v := range soss {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderSpreadsheetModel, vv, nil)
}

// UpdateSaleOrderSpreadsheet updates an existing sale.order.spreadsheet record.
func (c *Client) UpdateSaleOrderSpreadsheet(sos *SaleOrderSpreadsheet) error {
	return c.UpdateSaleOrderSpreadsheets([]int64{sos.Id.Get()}, sos)
}

// UpdateSaleOrderSpreadsheets updates existing sale.order.spreadsheet records.
// All records (represented by ids) will be updated by sos values.
func (c *Client) UpdateSaleOrderSpreadsheets(ids []int64, sos *SaleOrderSpreadsheet) error {
	return c.Update(SaleOrderSpreadsheetModel, ids, sos, nil)
}

// DeleteSaleOrderSpreadsheet deletes an existing sale.order.spreadsheet record.
func (c *Client) DeleteSaleOrderSpreadsheet(id int64) error {
	return c.DeleteSaleOrderSpreadsheets([]int64{id})
}

// DeleteSaleOrderSpreadsheets deletes existing sale.order.spreadsheet records.
func (c *Client) DeleteSaleOrderSpreadsheets(ids []int64) error {
	return c.Delete(SaleOrderSpreadsheetModel, ids)
}

// GetSaleOrderSpreadsheet gets sale.order.spreadsheet existing record.
func (c *Client) GetSaleOrderSpreadsheet(id int64) (*SaleOrderSpreadsheet, error) {
	soss, err := c.GetSaleOrderSpreadsheets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*soss)[0]), nil
}

// GetSaleOrderSpreadsheets gets sale.order.spreadsheet existing records.
func (c *Client) GetSaleOrderSpreadsheets(ids []int64) (*SaleOrderSpreadsheets, error) {
	soss := &SaleOrderSpreadsheets{}
	if err := c.Read(SaleOrderSpreadsheetModel, ids, nil, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindSaleOrderSpreadsheet finds sale.order.spreadsheet record by querying it with criteria.
func (c *Client) FindSaleOrderSpreadsheet(criteria *Criteria) (*SaleOrderSpreadsheet, error) {
	soss := &SaleOrderSpreadsheets{}
	if err := c.SearchRead(SaleOrderSpreadsheetModel, criteria, NewOptions().Limit(1), soss); err != nil {
		return nil, err
	}
	return &((*soss)[0]), nil
}

// FindSaleOrderSpreadsheets finds sale.order.spreadsheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderSpreadsheets(criteria *Criteria, options *Options) (*SaleOrderSpreadsheets, error) {
	soss := &SaleOrderSpreadsheets{}
	if err := c.SearchRead(SaleOrderSpreadsheetModel, criteria, options, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindSaleOrderSpreadsheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderSpreadsheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderSpreadsheetModel, criteria, options)
}

// FindSaleOrderSpreadsheetId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderSpreadsheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderSpreadsheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
