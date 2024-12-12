package odoo

// PosPreparationDisplayResetWizard represents pos_preparation_display.reset.wizard model.
type PosPreparationDisplayResetWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPreparationDisplayResetWizards represents array of pos_preparation_display.reset.wizard model.
type PosPreparationDisplayResetWizards []PosPreparationDisplayResetWizard

// PosPreparationDisplayResetWizardModel is the odoo model name.
const PosPreparationDisplayResetWizardModel = "pos_preparation_display.reset.wizard"

// Many2One convert PosPreparationDisplayResetWizard to *Many2One.
func (prw *PosPreparationDisplayResetWizard) Many2One() *Many2One {
	return NewMany2One(prw.Id.Get(), "")
}

// CreatePosPreparationDisplayResetWizard creates a new pos_preparation_display.reset.wizard model and returns its id.
func (c *Client) CreatePosPreparationDisplayResetWizard(prw *PosPreparationDisplayResetWizard) (int64, error) {
	ids, err := c.CreatePosPreparationDisplayResetWizards([]*PosPreparationDisplayResetWizard{prw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPreparationDisplayResetWizard creates a new pos_preparation_display.reset.wizard model and returns its id.
func (c *Client) CreatePosPreparationDisplayResetWizards(prws []*PosPreparationDisplayResetWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range prws {
		vv = append(vv, v)
	}
	return c.Create(PosPreparationDisplayResetWizardModel, vv, nil)
}

// UpdatePosPreparationDisplayResetWizard updates an existing pos_preparation_display.reset.wizard record.
func (c *Client) UpdatePosPreparationDisplayResetWizard(prw *PosPreparationDisplayResetWizard) error {
	return c.UpdatePosPreparationDisplayResetWizards([]int64{prw.Id.Get()}, prw)
}

// UpdatePosPreparationDisplayResetWizards updates existing pos_preparation_display.reset.wizard records.
// All records (represented by ids) will be updated by prw values.
func (c *Client) UpdatePosPreparationDisplayResetWizards(ids []int64, prw *PosPreparationDisplayResetWizard) error {
	return c.Update(PosPreparationDisplayResetWizardModel, ids, prw, nil)
}

// DeletePosPreparationDisplayResetWizard deletes an existing pos_preparation_display.reset.wizard record.
func (c *Client) DeletePosPreparationDisplayResetWizard(id int64) error {
	return c.DeletePosPreparationDisplayResetWizards([]int64{id})
}

// DeletePosPreparationDisplayResetWizards deletes existing pos_preparation_display.reset.wizard records.
func (c *Client) DeletePosPreparationDisplayResetWizards(ids []int64) error {
	return c.Delete(PosPreparationDisplayResetWizardModel, ids)
}

// GetPosPreparationDisplayResetWizard gets pos_preparation_display.reset.wizard existing record.
func (c *Client) GetPosPreparationDisplayResetWizard(id int64) (*PosPreparationDisplayResetWizard, error) {
	prws, err := c.GetPosPreparationDisplayResetWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prws)[0]), nil
}

// GetPosPreparationDisplayResetWizards gets pos_preparation_display.reset.wizard existing records.
func (c *Client) GetPosPreparationDisplayResetWizards(ids []int64) (*PosPreparationDisplayResetWizards, error) {
	prws := &PosPreparationDisplayResetWizards{}
	if err := c.Read(PosPreparationDisplayResetWizardModel, ids, nil, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPosPreparationDisplayResetWizard finds pos_preparation_display.reset.wizard record by querying it with criteria.
func (c *Client) FindPosPreparationDisplayResetWizard(criteria *Criteria) (*PosPreparationDisplayResetWizard, error) {
	prws := &PosPreparationDisplayResetWizards{}
	if err := c.SearchRead(PosPreparationDisplayResetWizardModel, criteria, NewOptions().Limit(1), prws); err != nil {
		return nil, err
	}
	return &((*prws)[0]), nil
}

// FindPosPreparationDisplayResetWizards finds pos_preparation_display.reset.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayResetWizards(criteria *Criteria, options *Options) (*PosPreparationDisplayResetWizards, error) {
	prws := &PosPreparationDisplayResetWizards{}
	if err := c.SearchRead(PosPreparationDisplayResetWizardModel, criteria, options, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPosPreparationDisplayResetWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayResetWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPreparationDisplayResetWizardModel, criteria, options)
}

// FindPosPreparationDisplayResetWizardId finds record id by querying it with criteria.
func (c *Client) FindPosPreparationDisplayResetWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPreparationDisplayResetWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
