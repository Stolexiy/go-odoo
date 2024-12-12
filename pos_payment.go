package odoo

// PosPayment represents pos.payment model.
type PosPayment struct {
	AccountMoveId            *Many2One `xmlrpc:"account_move_id,omitempty"`
	Amount                   *Float    `xmlrpc:"amount,omitempty"`
	CardBrand                *String   `xmlrpc:"card_brand,omitempty"`
	CardNo                   *String   `xmlrpc:"card_no,omitempty"`
	CardType                 *String   `xmlrpc:"card_type,omitempty"`
	CardholderName           *String   `xmlrpc:"cardholder_name,omitempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omitempty"`
	CurrencyRate             *Float    `xmlrpc:"currency_rate,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId               *Many2One `xmlrpc:"employee_id,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	IsChange                 *Bool     `xmlrpc:"is_change,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	OnlineAccountPaymentId   *Many2One `xmlrpc:"online_account_payment_id,omitempty"`
	PartnerId                *Many2One `xmlrpc:"partner_id,omitempty"`
	PaymentDate              *Time     `xmlrpc:"payment_date,omitempty"`
	PaymentMethodAuthcode    *String   `xmlrpc:"payment_method_authcode,omitempty"`
	PaymentMethodId          *Many2One `xmlrpc:"payment_method_id,omitempty"`
	PaymentMethodIssuerBank  *String   `xmlrpc:"payment_method_issuer_bank,omitempty"`
	PaymentMethodPaymentMode *String   `xmlrpc:"payment_method_payment_mode,omitempty"`
	PaymentRefNo             *String   `xmlrpc:"payment_ref_no,omitempty"`
	PaymentStatus            *String   `xmlrpc:"payment_status,omitempty"`
	PosOrderId               *Many2One `xmlrpc:"pos_order_id,omitempty"`
	SessionId                *Many2One `xmlrpc:"session_id,omitempty"`
	Ticket                   *String   `xmlrpc:"ticket,omitempty"`
	TransactionId            *String   `xmlrpc:"transaction_id,omitempty"`
	UserId                   *Many2One `xmlrpc:"user_id,omitempty"`
	Uuid                     *String   `xmlrpc:"uuid,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPayments represents array of pos.payment model.
type PosPayments []PosPayment

// PosPaymentModel is the odoo model name.
const PosPaymentModel = "pos.payment"

// Many2One convert PosPayment to *Many2One.
func (pp *PosPayment) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePosPayment creates a new pos.payment model and returns its id.
func (c *Client) CreatePosPayment(pp *PosPayment) (int64, error) {
	ids, err := c.CreatePosPayments([]*PosPayment{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPayment creates a new pos.payment model and returns its id.
func (c *Client) CreatePosPayments(pps []*PosPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(PosPaymentModel, vv, nil)
}

// UpdatePosPayment updates an existing pos.payment record.
func (c *Client) UpdatePosPayment(pp *PosPayment) error {
	return c.UpdatePosPayments([]int64{pp.Id.Get()}, pp)
}

// UpdatePosPayments updates existing pos.payment records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePosPayments(ids []int64, pp *PosPayment) error {
	return c.Update(PosPaymentModel, ids, pp, nil)
}

// DeletePosPayment deletes an existing pos.payment record.
func (c *Client) DeletePosPayment(id int64) error {
	return c.DeletePosPayments([]int64{id})
}

// DeletePosPayments deletes existing pos.payment records.
func (c *Client) DeletePosPayments(ids []int64) error {
	return c.Delete(PosPaymentModel, ids)
}

// GetPosPayment gets pos.payment existing record.
func (c *Client) GetPosPayment(id int64) (*PosPayment, error) {
	pps, err := c.GetPosPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetPosPayments gets pos.payment existing records.
func (c *Client) GetPosPayments(ids []int64) (*PosPayments, error) {
	pps := &PosPayments{}
	if err := c.Read(PosPaymentModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPayment finds pos.payment record by querying it with criteria.
func (c *Client) FindPosPayment(criteria *Criteria) (*PosPayment, error) {
	pps := &PosPayments{}
	if err := c.SearchRead(PosPaymentModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindPosPayments finds pos.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPayments(criteria *Criteria, options *Options) (*PosPayments, error) {
	pps := &PosPayments{}
	if err := c.SearchRead(PosPaymentModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPaymentModel, criteria, options)
}

// FindPosPaymentId finds record id by querying it with criteria.
func (c *Client) FindPosPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
