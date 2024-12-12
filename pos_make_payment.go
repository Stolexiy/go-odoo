package odoo

// PosMakePayment represents pos.make.payment model.
type PosMakePayment struct {
	Amount          *Float    `xmlrpc:"amount,omitempty"`
	ConfigId        *Many2One `xmlrpc:"config_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	PaymentDate     *Time     `xmlrpc:"payment_date,omitempty"`
	PaymentMethodId *Many2One `xmlrpc:"payment_method_id,omitempty"`
	PaymentName     *String   `xmlrpc:"payment_name,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosMakePayments represents array of pos.make.payment model.
type PosMakePayments []PosMakePayment

// PosMakePaymentModel is the odoo model name.
const PosMakePaymentModel = "pos.make.payment"

// Many2One convert PosMakePayment to *Many2One.
func (pmp *PosMakePayment) Many2One() *Many2One {
	return NewMany2One(pmp.Id.Get(), "")
}

// CreatePosMakePayment creates a new pos.make.payment model and returns its id.
func (c *Client) CreatePosMakePayment(pmp *PosMakePayment) (int64, error) {
	ids, err := c.CreatePosMakePayments([]*PosMakePayment{pmp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosMakePayment creates a new pos.make.payment model and returns its id.
func (c *Client) CreatePosMakePayments(pmps []*PosMakePayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range pmps {
		vv = append(vv, v)
	}
	return c.Create(PosMakePaymentModel, vv, nil)
}

// UpdatePosMakePayment updates an existing pos.make.payment record.
func (c *Client) UpdatePosMakePayment(pmp *PosMakePayment) error {
	return c.UpdatePosMakePayments([]int64{pmp.Id.Get()}, pmp)
}

// UpdatePosMakePayments updates existing pos.make.payment records.
// All records (represented by ids) will be updated by pmp values.
func (c *Client) UpdatePosMakePayments(ids []int64, pmp *PosMakePayment) error {
	return c.Update(PosMakePaymentModel, ids, pmp, nil)
}

// DeletePosMakePayment deletes an existing pos.make.payment record.
func (c *Client) DeletePosMakePayment(id int64) error {
	return c.DeletePosMakePayments([]int64{id})
}

// DeletePosMakePayments deletes existing pos.make.payment records.
func (c *Client) DeletePosMakePayments(ids []int64) error {
	return c.Delete(PosMakePaymentModel, ids)
}

// GetPosMakePayment gets pos.make.payment existing record.
func (c *Client) GetPosMakePayment(id int64) (*PosMakePayment, error) {
	pmps, err := c.GetPosMakePayments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pmps)[0]), nil
}

// GetPosMakePayments gets pos.make.payment existing records.
func (c *Client) GetPosMakePayments(ids []int64) (*PosMakePayments, error) {
	pmps := &PosMakePayments{}
	if err := c.Read(PosMakePaymentModel, ids, nil, pmps); err != nil {
		return nil, err
	}
	return pmps, nil
}

// FindPosMakePayment finds pos.make.payment record by querying it with criteria.
func (c *Client) FindPosMakePayment(criteria *Criteria) (*PosMakePayment, error) {
	pmps := &PosMakePayments{}
	if err := c.SearchRead(PosMakePaymentModel, criteria, NewOptions().Limit(1), pmps); err != nil {
		return nil, err
	}
	return &((*pmps)[0]), nil
}

// FindPosMakePayments finds pos.make.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMakePayments(criteria *Criteria, options *Options) (*PosMakePayments, error) {
	pmps := &PosMakePayments{}
	if err := c.SearchRead(PosMakePaymentModel, criteria, options, pmps); err != nil {
		return nil, err
	}
	return pmps, nil
}

// FindPosMakePaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMakePaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosMakePaymentModel, criteria, options)
}

// FindPosMakePaymentId finds record id by querying it with criteria.
func (c *Client) FindPosMakePaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosMakePaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
