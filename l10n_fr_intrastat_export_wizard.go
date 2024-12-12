package odoo

// L10NFrIntrastatExportWizard represents l10n_fr_intrastat.export.wizard model.
type L10NFrIntrastatExportWizard struct {
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	EmebiFlow                  *Selection `xmlrpc:"emebi_flow,omitempty"`
	EmebiFlowVisible           *Bool      `xmlrpc:"emebi_flow_visible,omitempty"`
	ExportType                 *Selection `xmlrpc:"export_type,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	WarningIncompatibleOptions *Bool      `xmlrpc:"warning_incompatible_options,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// L10NFrIntrastatExportWizards represents array of l10n_fr_intrastat.export.wizard model.
type L10NFrIntrastatExportWizards []L10NFrIntrastatExportWizard

// L10NFrIntrastatExportWizardModel is the odoo model name.
const L10NFrIntrastatExportWizardModel = "l10n_fr_intrastat.export.wizard"

// Many2One convert L10NFrIntrastatExportWizard to *Many2One.
func (lew *L10NFrIntrastatExportWizard) Many2One() *Many2One {
	return NewMany2One(lew.Id.Get(), "")
}

// CreateL10NFrIntrastatExportWizard creates a new l10n_fr_intrastat.export.wizard model and returns its id.
func (c *Client) CreateL10NFrIntrastatExportWizard(lew *L10NFrIntrastatExportWizard) (int64, error) {
	ids, err := c.CreateL10NFrIntrastatExportWizards([]*L10NFrIntrastatExportWizard{lew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NFrIntrastatExportWizard creates a new l10n_fr_intrastat.export.wizard model and returns its id.
func (c *Client) CreateL10NFrIntrastatExportWizards(lews []*L10NFrIntrastatExportWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lews {
		vv = append(vv, v)
	}
	return c.Create(L10NFrIntrastatExportWizardModel, vv, nil)
}

// UpdateL10NFrIntrastatExportWizard updates an existing l10n_fr_intrastat.export.wizard record.
func (c *Client) UpdateL10NFrIntrastatExportWizard(lew *L10NFrIntrastatExportWizard) error {
	return c.UpdateL10NFrIntrastatExportWizards([]int64{lew.Id.Get()}, lew)
}

// UpdateL10NFrIntrastatExportWizards updates existing l10n_fr_intrastat.export.wizard records.
// All records (represented by ids) will be updated by lew values.
func (c *Client) UpdateL10NFrIntrastatExportWizards(ids []int64, lew *L10NFrIntrastatExportWizard) error {
	return c.Update(L10NFrIntrastatExportWizardModel, ids, lew, nil)
}

// DeleteL10NFrIntrastatExportWizard deletes an existing l10n_fr_intrastat.export.wizard record.
func (c *Client) DeleteL10NFrIntrastatExportWizard(id int64) error {
	return c.DeleteL10NFrIntrastatExportWizards([]int64{id})
}

// DeleteL10NFrIntrastatExportWizards deletes existing l10n_fr_intrastat.export.wizard records.
func (c *Client) DeleteL10NFrIntrastatExportWizards(ids []int64) error {
	return c.Delete(L10NFrIntrastatExportWizardModel, ids)
}

// GetL10NFrIntrastatExportWizard gets l10n_fr_intrastat.export.wizard existing record.
func (c *Client) GetL10NFrIntrastatExportWizard(id int64) (*L10NFrIntrastatExportWizard, error) {
	lews, err := c.GetL10NFrIntrastatExportWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lews)[0]), nil
}

// GetL10NFrIntrastatExportWizards gets l10n_fr_intrastat.export.wizard existing records.
func (c *Client) GetL10NFrIntrastatExportWizards(ids []int64) (*L10NFrIntrastatExportWizards, error) {
	lews := &L10NFrIntrastatExportWizards{}
	if err := c.Read(L10NFrIntrastatExportWizardModel, ids, nil, lews); err != nil {
		return nil, err
	}
	return lews, nil
}

// FindL10NFrIntrastatExportWizard finds l10n_fr_intrastat.export.wizard record by querying it with criteria.
func (c *Client) FindL10NFrIntrastatExportWizard(criteria *Criteria) (*L10NFrIntrastatExportWizard, error) {
	lews := &L10NFrIntrastatExportWizards{}
	if err := c.SearchRead(L10NFrIntrastatExportWizardModel, criteria, NewOptions().Limit(1), lews); err != nil {
		return nil, err
	}
	return &((*lews)[0]), nil
}

// FindL10NFrIntrastatExportWizards finds l10n_fr_intrastat.export.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrIntrastatExportWizards(criteria *Criteria, options *Options) (*L10NFrIntrastatExportWizards, error) {
	lews := &L10NFrIntrastatExportWizards{}
	if err := c.SearchRead(L10NFrIntrastatExportWizardModel, criteria, options, lews); err != nil {
		return nil, err
	}
	return lews, nil
}

// FindL10NFrIntrastatExportWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrIntrastatExportWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NFrIntrastatExportWizardModel, criteria, options)
}

// FindL10NFrIntrastatExportWizardId finds record id by querying it with criteria.
func (c *Client) FindL10NFrIntrastatExportWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NFrIntrastatExportWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
