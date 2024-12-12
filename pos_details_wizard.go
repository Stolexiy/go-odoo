package odoo

// PosDetailsWizard represents pos.details.wizard model.
type PosDetailsWizard struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	EndDate      *Time     `xmlrpc:"end_date,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	PosConfigIds *Relation `xmlrpc:"pos_config_ids,omitempty"`
	StartDate    *Time     `xmlrpc:"start_date,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosDetailsWizards represents array of pos.details.wizard model.
type PosDetailsWizards []PosDetailsWizard

// PosDetailsWizardModel is the odoo model name.
const PosDetailsWizardModel = "pos.details.wizard"

// Many2One convert PosDetailsWizard to *Many2One.
func (pdw *PosDetailsWizard) Many2One() *Many2One {
	return NewMany2One(pdw.Id.Get(), "")
}

// CreatePosDetailsWizard creates a new pos.details.wizard model and returns its id.
func (c *Client) CreatePosDetailsWizard(pdw *PosDetailsWizard) (int64, error) {
	ids, err := c.CreatePosDetailsWizards([]*PosDetailsWizard{pdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosDetailsWizard creates a new pos.details.wizard model and returns its id.
func (c *Client) CreatePosDetailsWizards(pdws []*PosDetailsWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pdws {
		vv = append(vv, v)
	}
	return c.Create(PosDetailsWizardModel, vv, nil)
}

// UpdatePosDetailsWizard updates an existing pos.details.wizard record.
func (c *Client) UpdatePosDetailsWizard(pdw *PosDetailsWizard) error {
	return c.UpdatePosDetailsWizards([]int64{pdw.Id.Get()}, pdw)
}

// UpdatePosDetailsWizards updates existing pos.details.wizard records.
// All records (represented by ids) will be updated by pdw values.
func (c *Client) UpdatePosDetailsWizards(ids []int64, pdw *PosDetailsWizard) error {
	return c.Update(PosDetailsWizardModel, ids, pdw, nil)
}

// DeletePosDetailsWizard deletes an existing pos.details.wizard record.
func (c *Client) DeletePosDetailsWizard(id int64) error {
	return c.DeletePosDetailsWizards([]int64{id})
}

// DeletePosDetailsWizards deletes existing pos.details.wizard records.
func (c *Client) DeletePosDetailsWizards(ids []int64) error {
	return c.Delete(PosDetailsWizardModel, ids)
}

// GetPosDetailsWizard gets pos.details.wizard existing record.
func (c *Client) GetPosDetailsWizard(id int64) (*PosDetailsWizard, error) {
	pdws, err := c.GetPosDetailsWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pdws)[0]), nil
}

// GetPosDetailsWizards gets pos.details.wizard existing records.
func (c *Client) GetPosDetailsWizards(ids []int64) (*PosDetailsWizards, error) {
	pdws := &PosDetailsWizards{}
	if err := c.Read(PosDetailsWizardModel, ids, nil, pdws); err != nil {
		return nil, err
	}
	return pdws, nil
}

// FindPosDetailsWizard finds pos.details.wizard record by querying it with criteria.
func (c *Client) FindPosDetailsWizard(criteria *Criteria) (*PosDetailsWizard, error) {
	pdws := &PosDetailsWizards{}
	if err := c.SearchRead(PosDetailsWizardModel, criteria, NewOptions().Limit(1), pdws); err != nil {
		return nil, err
	}
	return &((*pdws)[0]), nil
}

// FindPosDetailsWizards finds pos.details.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDetailsWizards(criteria *Criteria, options *Options) (*PosDetailsWizards, error) {
	pdws := &PosDetailsWizards{}
	if err := c.SearchRead(PosDetailsWizardModel, criteria, options, pdws); err != nil {
		return nil, err
	}
	return pdws, nil
}

// FindPosDetailsWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDetailsWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosDetailsWizardModel, criteria, options)
}

// FindPosDetailsWizardId finds record id by querying it with criteria.
func (c *Client) FindPosDetailsWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosDetailsWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
