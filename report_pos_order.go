package odoo

// ReportPosOrder represents report.pos.order model.
type ReportPosOrder struct {
	AveragePrice      *Float     `xmlrpc:"average_price,omitempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omitempty"`
	ConfigId          *Many2One  `xmlrpc:"config_id,omitempty"`
	Date              *Time      `xmlrpc:"date,omitempty"`
	DelayValidation   *Int       `xmlrpc:"delay_validation,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId        *Many2One  `xmlrpc:"employee_id,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	Invoiced          *Bool      `xmlrpc:"invoiced,omitempty"`
	JournalId         *Many2One  `xmlrpc:"journal_id,omitempty"`
	Margin            *Float     `xmlrpc:"margin,omitempty"`
	NbrLines          *Int       `xmlrpc:"nbr_lines,omitempty"`
	OrderId           *Many2One  `xmlrpc:"order_id,omitempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omitempty"`
	PaymentMethodId   *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	PosCategId        *Many2One  `xmlrpc:"pos_categ_id,omitempty"`
	PriceSubTotal     *Float     `xmlrpc:"price_sub_total,omitempty"`
	PriceSubtotalExcl *Float     `xmlrpc:"price_subtotal_excl,omitempty"`
	PriceTotal        *Float     `xmlrpc:"price_total,omitempty"`
	PricelistId       *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProductCategId    *Many2One  `xmlrpc:"product_categ_id,omitempty"`
	ProductId         *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductQty        *Int       `xmlrpc:"product_qty,omitempty"`
	ProductTmplId     *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	SessionId         *Many2One  `xmlrpc:"session_id,omitempty"`
	State             *Selection `xmlrpc:"state,omitempty"`
	TotalDiscount     *Float     `xmlrpc:"total_discount,omitempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omitempty"`
}

// ReportPosOrders represents array of report.pos.order model.
type ReportPosOrders []ReportPosOrder

// ReportPosOrderModel is the odoo model name.
const ReportPosOrderModel = "report.pos.order"

// Many2One convert ReportPosOrder to *Many2One.
func (rpo *ReportPosOrder) Many2One() *Many2One {
	return NewMany2One(rpo.Id.Get(), "")
}

// CreateReportPosOrder creates a new report.pos.order model and returns its id.
func (c *Client) CreateReportPosOrder(rpo *ReportPosOrder) (int64, error) {
	ids, err := c.CreateReportPosOrders([]*ReportPosOrder{rpo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportPosOrder creates a new report.pos.order model and returns its id.
func (c *Client) CreateReportPosOrders(rpos []*ReportPosOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range rpos {
		vv = append(vv, v)
	}
	return c.Create(ReportPosOrderModel, vv, nil)
}

// UpdateReportPosOrder updates an existing report.pos.order record.
func (c *Client) UpdateReportPosOrder(rpo *ReportPosOrder) error {
	return c.UpdateReportPosOrders([]int64{rpo.Id.Get()}, rpo)
}

// UpdateReportPosOrders updates existing report.pos.order records.
// All records (represented by ids) will be updated by rpo values.
func (c *Client) UpdateReportPosOrders(ids []int64, rpo *ReportPosOrder) error {
	return c.Update(ReportPosOrderModel, ids, rpo, nil)
}

// DeleteReportPosOrder deletes an existing report.pos.order record.
func (c *Client) DeleteReportPosOrder(id int64) error {
	return c.DeleteReportPosOrders([]int64{id})
}

// DeleteReportPosOrders deletes existing report.pos.order records.
func (c *Client) DeleteReportPosOrders(ids []int64) error {
	return c.Delete(ReportPosOrderModel, ids)
}

// GetReportPosOrder gets report.pos.order existing record.
func (c *Client) GetReportPosOrder(id int64) (*ReportPosOrder, error) {
	rpos, err := c.GetReportPosOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rpos)[0]), nil
}

// GetReportPosOrders gets report.pos.order existing records.
func (c *Client) GetReportPosOrders(ids []int64) (*ReportPosOrders, error) {
	rpos := &ReportPosOrders{}
	if err := c.Read(ReportPosOrderModel, ids, nil, rpos); err != nil {
		return nil, err
	}
	return rpos, nil
}

// FindReportPosOrder finds report.pos.order record by querying it with criteria.
func (c *Client) FindReportPosOrder(criteria *Criteria) (*ReportPosOrder, error) {
	rpos := &ReportPosOrders{}
	if err := c.SearchRead(ReportPosOrderModel, criteria, NewOptions().Limit(1), rpos); err != nil {
		return nil, err
	}
	return &((*rpos)[0]), nil
}

// FindReportPosOrders finds report.pos.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPosOrders(criteria *Criteria, options *Options) (*ReportPosOrders, error) {
	rpos := &ReportPosOrders{}
	if err := c.SearchRead(ReportPosOrderModel, criteria, options, rpos); err != nil {
		return nil, err
	}
	return rpos, nil
}

// FindReportPosOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPosOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportPosOrderModel, criteria, options)
}

// FindReportPosOrderId finds record id by querying it with criteria.
func (c *Client) FindReportPosOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportPosOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
