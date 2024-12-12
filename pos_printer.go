package odoo

// PosPrinter represents pos.printer model.
type PosPrinter struct {
	CompanyId            *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	EpsonPrinterIp       *String    `xmlrpc:"epson_printer_ip,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty"`
	PrinterType          *Selection `xmlrpc:"printer_type,omitempty"`
	ProductCategoriesIds *Relation  `xmlrpc:"product_categories_ids,omitempty"`
	ProxyIp              *String    `xmlrpc:"proxy_ip,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PosPrinters represents array of pos.printer model.
type PosPrinters []PosPrinter

// PosPrinterModel is the odoo model name.
const PosPrinterModel = "pos.printer"

// Many2One convert PosPrinter to *Many2One.
func (pp *PosPrinter) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePosPrinter creates a new pos.printer model and returns its id.
func (c *Client) CreatePosPrinter(pp *PosPrinter) (int64, error) {
	ids, err := c.CreatePosPrinters([]*PosPrinter{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPrinter creates a new pos.printer model and returns its id.
func (c *Client) CreatePosPrinters(pps []*PosPrinter) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(PosPrinterModel, vv, nil)
}

// UpdatePosPrinter updates an existing pos.printer record.
func (c *Client) UpdatePosPrinter(pp *PosPrinter) error {
	return c.UpdatePosPrinters([]int64{pp.Id.Get()}, pp)
}

// UpdatePosPrinters updates existing pos.printer records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePosPrinters(ids []int64, pp *PosPrinter) error {
	return c.Update(PosPrinterModel, ids, pp, nil)
}

// DeletePosPrinter deletes an existing pos.printer record.
func (c *Client) DeletePosPrinter(id int64) error {
	return c.DeletePosPrinters([]int64{id})
}

// DeletePosPrinters deletes existing pos.printer records.
func (c *Client) DeletePosPrinters(ids []int64) error {
	return c.Delete(PosPrinterModel, ids)
}

// GetPosPrinter gets pos.printer existing record.
func (c *Client) GetPosPrinter(id int64) (*PosPrinter, error) {
	pps, err := c.GetPosPrinters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetPosPrinters gets pos.printer existing records.
func (c *Client) GetPosPrinters(ids []int64) (*PosPrinters, error) {
	pps := &PosPrinters{}
	if err := c.Read(PosPrinterModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPrinter finds pos.printer record by querying it with criteria.
func (c *Client) FindPosPrinter(criteria *Criteria) (*PosPrinter, error) {
	pps := &PosPrinters{}
	if err := c.SearchRead(PosPrinterModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindPosPrinters finds pos.printer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPrinters(criteria *Criteria, options *Options) (*PosPrinters, error) {
	pps := &PosPrinters{}
	if err := c.SearchRead(PosPrinterModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPrinterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPrinterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPrinterModel, criteria, options)
}

// FindPosPrinterId finds record id by querying it with criteria.
func (c *Client) FindPosPrinterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPrinterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
