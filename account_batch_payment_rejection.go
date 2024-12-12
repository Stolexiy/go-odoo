package odoo

// AccountBatchPaymentRejection represents account.batch.payment.rejection model.
type AccountBatchPaymentRejection struct {
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	InReconcilePaymentIds *Relation `xmlrpc:"in_reconcile_payment_ids,omitempty"`
	NbBatchPaymentIds     *Int      `xmlrpc:"nb_batch_payment_ids,omitempty"`
	NbRejectedPaymentIds  *Int      `xmlrpc:"nb_rejected_payment_ids,omitempty"`
	RejectedPaymentIds    *Relation `xmlrpc:"rejected_payment_ids,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountBatchPaymentRejections represents array of account.batch.payment.rejection model.
type AccountBatchPaymentRejections []AccountBatchPaymentRejection

// AccountBatchPaymentRejectionModel is the odoo model name.
const AccountBatchPaymentRejectionModel = "account.batch.payment.rejection"

// Many2One convert AccountBatchPaymentRejection to *Many2One.
func (abpr *AccountBatchPaymentRejection) Many2One() *Many2One {
	return NewMany2One(abpr.Id.Get(), "")
}

// CreateAccountBatchPaymentRejection creates a new account.batch.payment.rejection model and returns its id.
func (c *Client) CreateAccountBatchPaymentRejection(abpr *AccountBatchPaymentRejection) (int64, error) {
	ids, err := c.CreateAccountBatchPaymentRejections([]*AccountBatchPaymentRejection{abpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchPaymentRejection creates a new account.batch.payment.rejection model and returns its id.
func (c *Client) CreateAccountBatchPaymentRejections(abprs []*AccountBatchPaymentRejection) ([]int64, error) {
	var vv []interface{}
	for _, v := range abprs {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchPaymentRejectionModel, vv, nil)
}

// UpdateAccountBatchPaymentRejection updates an existing account.batch.payment.rejection record.
func (c *Client) UpdateAccountBatchPaymentRejection(abpr *AccountBatchPaymentRejection) error {
	return c.UpdateAccountBatchPaymentRejections([]int64{abpr.Id.Get()}, abpr)
}

// UpdateAccountBatchPaymentRejections updates existing account.batch.payment.rejection records.
// All records (represented by ids) will be updated by abpr values.
func (c *Client) UpdateAccountBatchPaymentRejections(ids []int64, abpr *AccountBatchPaymentRejection) error {
	return c.Update(AccountBatchPaymentRejectionModel, ids, abpr, nil)
}

// DeleteAccountBatchPaymentRejection deletes an existing account.batch.payment.rejection record.
func (c *Client) DeleteAccountBatchPaymentRejection(id int64) error {
	return c.DeleteAccountBatchPaymentRejections([]int64{id})
}

// DeleteAccountBatchPaymentRejections deletes existing account.batch.payment.rejection records.
func (c *Client) DeleteAccountBatchPaymentRejections(ids []int64) error {
	return c.Delete(AccountBatchPaymentRejectionModel, ids)
}

// GetAccountBatchPaymentRejection gets account.batch.payment.rejection existing record.
func (c *Client) GetAccountBatchPaymentRejection(id int64) (*AccountBatchPaymentRejection, error) {
	abprs, err := c.GetAccountBatchPaymentRejections([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abprs)[0]), nil
}

// GetAccountBatchPaymentRejections gets account.batch.payment.rejection existing records.
func (c *Client) GetAccountBatchPaymentRejections(ids []int64) (*AccountBatchPaymentRejections, error) {
	abprs := &AccountBatchPaymentRejections{}
	if err := c.Read(AccountBatchPaymentRejectionModel, ids, nil, abprs); err != nil {
		return nil, err
	}
	return abprs, nil
}

// FindAccountBatchPaymentRejection finds account.batch.payment.rejection record by querying it with criteria.
func (c *Client) FindAccountBatchPaymentRejection(criteria *Criteria) (*AccountBatchPaymentRejection, error) {
	abprs := &AccountBatchPaymentRejections{}
	if err := c.SearchRead(AccountBatchPaymentRejectionModel, criteria, NewOptions().Limit(1), abprs); err != nil {
		return nil, err
	}
	return &((*abprs)[0]), nil
}

// FindAccountBatchPaymentRejections finds account.batch.payment.rejection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPaymentRejections(criteria *Criteria, options *Options) (*AccountBatchPaymentRejections, error) {
	abprs := &AccountBatchPaymentRejections{}
	if err := c.SearchRead(AccountBatchPaymentRejectionModel, criteria, options, abprs); err != nil {
		return nil, err
	}
	return abprs, nil
}

// FindAccountBatchPaymentRejectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPaymentRejectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountBatchPaymentRejectionModel, criteria, options)
}

// FindAccountBatchPaymentRejectionId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchPaymentRejectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchPaymentRejectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
