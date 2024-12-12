package odoo

// PosPaymentMethod represents pos.payment.method model.
type PosPaymentMethod struct {
	Active                     *Bool      `xmlrpc:"active,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	ConfigIds                  *Relation  `xmlrpc:"config_ids,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultQr                  *String    `xmlrpc:"default_qr,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	HasAnOnlinePaymentProvider *Bool      `xmlrpc:"has_an_online_payment_provider,omitempty"`
	HideQrCodeMethod           *Bool      `xmlrpc:"hide_qr_code_method,omitempty"`
	HideUsePaymentTerminal     *Bool      `xmlrpc:"hide_use_payment_terminal,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	Image                      *String    `xmlrpc:"image,omitempty"`
	IsCashCount                *Bool      `xmlrpc:"is_cash_count,omitempty"`
	IsOnlinePayment            *Bool      `xmlrpc:"is_online_payment,omitempty"`
	JournalId                  *Many2One  `xmlrpc:"journal_id,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	OnlinePaymentProviderIds   *Relation  `xmlrpc:"online_payment_provider_ids,omitempty"`
	OpenSessionIds             *Relation  `xmlrpc:"open_session_ids,omitempty"`
	OutstandingAccountId       *Many2One  `xmlrpc:"outstanding_account_id,omitempty"`
	PaymentMethodType          *Selection `xmlrpc:"payment_method_type,omitempty"`
	QrCodeMethod               *Selection `xmlrpc:"qr_code_method,omitempty"`
	ReceivableAccountId        *Many2One  `xmlrpc:"receivable_account_id,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	SplitTransactions          *Bool      `xmlrpc:"split_transactions,omitempty"`
	StripeSerialNumber         *String    `xmlrpc:"stripe_serial_number,omitempty"`
	Type                       *Selection `xmlrpc:"type,omitempty"`
	UsePaymentTerminal         *Selection `xmlrpc:"use_payment_terminal,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PosPaymentMethods represents array of pos.payment.method model.
type PosPaymentMethods []PosPaymentMethod

// PosPaymentMethodModel is the odoo model name.
const PosPaymentMethodModel = "pos.payment.method"

// Many2One convert PosPaymentMethod to *Many2One.
func (ppm *PosPaymentMethod) Many2One() *Many2One {
	return NewMany2One(ppm.Id.Get(), "")
}

// CreatePosPaymentMethod creates a new pos.payment.method model and returns its id.
func (c *Client) CreatePosPaymentMethod(ppm *PosPaymentMethod) (int64, error) {
	ids, err := c.CreatePosPaymentMethods([]*PosPaymentMethod{ppm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPaymentMethod creates a new pos.payment.method model and returns its id.
func (c *Client) CreatePosPaymentMethods(ppms []*PosPaymentMethod) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppms {
		vv = append(vv, v)
	}
	return c.Create(PosPaymentMethodModel, vv, nil)
}

// UpdatePosPaymentMethod updates an existing pos.payment.method record.
func (c *Client) UpdatePosPaymentMethod(ppm *PosPaymentMethod) error {
	return c.UpdatePosPaymentMethods([]int64{ppm.Id.Get()}, ppm)
}

// UpdatePosPaymentMethods updates existing pos.payment.method records.
// All records (represented by ids) will be updated by ppm values.
func (c *Client) UpdatePosPaymentMethods(ids []int64, ppm *PosPaymentMethod) error {
	return c.Update(PosPaymentMethodModel, ids, ppm, nil)
}

// DeletePosPaymentMethod deletes an existing pos.payment.method record.
func (c *Client) DeletePosPaymentMethod(id int64) error {
	return c.DeletePosPaymentMethods([]int64{id})
}

// DeletePosPaymentMethods deletes existing pos.payment.method records.
func (c *Client) DeletePosPaymentMethods(ids []int64) error {
	return c.Delete(PosPaymentMethodModel, ids)
}

// GetPosPaymentMethod gets pos.payment.method existing record.
func (c *Client) GetPosPaymentMethod(id int64) (*PosPaymentMethod, error) {
	ppms, err := c.GetPosPaymentMethods([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ppms)[0]), nil
}

// GetPosPaymentMethods gets pos.payment.method existing records.
func (c *Client) GetPosPaymentMethods(ids []int64) (*PosPaymentMethods, error) {
	ppms := &PosPaymentMethods{}
	if err := c.Read(PosPaymentMethodModel, ids, nil, ppms); err != nil {
		return nil, err
	}
	return ppms, nil
}

// FindPosPaymentMethod finds pos.payment.method record by querying it with criteria.
func (c *Client) FindPosPaymentMethod(criteria *Criteria) (*PosPaymentMethod, error) {
	ppms := &PosPaymentMethods{}
	if err := c.SearchRead(PosPaymentMethodModel, criteria, NewOptions().Limit(1), ppms); err != nil {
		return nil, err
	}
	return &((*ppms)[0]), nil
}

// FindPosPaymentMethods finds pos.payment.method records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPaymentMethods(criteria *Criteria, options *Options) (*PosPaymentMethods, error) {
	ppms := &PosPaymentMethods{}
	if err := c.SearchRead(PosPaymentMethodModel, criteria, options, ppms); err != nil {
		return nil, err
	}
	return ppms, nil
}

// FindPosPaymentMethodIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPaymentMethodIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPaymentMethodModel, criteria, options)
}

// FindPosPaymentMethodId finds record id by querying it with criteria.
func (c *Client) FindPosPaymentMethodId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPaymentMethodModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
