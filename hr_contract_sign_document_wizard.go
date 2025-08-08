package odoo

// HrContractSignDocumentWizard represents hr.contract.sign.document.wizard model.
type HrContractSignDocumentWizard struct {
	AttachmentIds              *Relation  `xmlrpc:"attachment_ids,omitempty"`
	CcPartnerIds               *Relation  `xmlrpc:"cc_partner_ids,omitempty"`
	ContractId                 *Many2One  `xmlrpc:"contract_id,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	EmployeeIds                *Relation  `xmlrpc:"employee_ids,omitempty"`
	EmployeeRoleId             *Many2One  `xmlrpc:"employee_role_id,omitempty"`
	HasBothTemplate            *Bool      `xmlrpc:"has_both_template,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	MailDisplayed              *String    `xmlrpc:"mail_displayed,omitempty"`
	MailTo                     *Selection `xmlrpc:"mail_to,omitempty"`
	Message                    *String    `xmlrpc:"message,omitempty"`
	PossibleTemplateIds        *Relation  `xmlrpc:"possible_template_ids,omitempty"`
	ResponsibleId              *Many2One  `xmlrpc:"responsible_id,omitempty"`
	SignTemplateIds            *Relation  `xmlrpc:"sign_template_ids,omitempty"`
	SignTemplateResponsibleIds *Relation  `xmlrpc:"sign_template_responsible_ids,omitempty"`
	Subject                    *String    `xmlrpc:"subject,omitempty"`
	TemplateWarning            *String    `xmlrpc:"template_warning,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrContractSignDocumentWizards represents array of hr.contract.sign.document.wizard model.
type HrContractSignDocumentWizards []HrContractSignDocumentWizard

// HrContractSignDocumentWizardModel is the odoo model name.
const HrContractSignDocumentWizardModel = "hr.contract.sign.document.wizard"

// Many2One convert HrContractSignDocumentWizard to *Many2One.
func (hcsdw *HrContractSignDocumentWizard) Many2One() *Many2One {
	return NewMany2One(hcsdw.Id.Get(), "")
}

// CreateHrContractSignDocumentWizard creates a new hr.contract.sign.document.wizard model and returns its id.
func (c *Client) CreateHrContractSignDocumentWizard(hcsdw *HrContractSignDocumentWizard) (int64, error) {
	ids, err := c.CreateHrContractSignDocumentWizards([]*HrContractSignDocumentWizard{hcsdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrContractSignDocumentWizard creates a new hr.contract.sign.document.wizard model and returns its id.
func (c *Client) CreateHrContractSignDocumentWizards(hcsdws []*HrContractSignDocumentWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcsdws {
		vv = append(vv, v)
	}
	return c.Create(HrContractSignDocumentWizardModel, vv, nil)
}

// UpdateHrContractSignDocumentWizard updates an existing hr.contract.sign.document.wizard record.
func (c *Client) UpdateHrContractSignDocumentWizard(hcsdw *HrContractSignDocumentWizard) error {
	return c.UpdateHrContractSignDocumentWizards([]int64{hcsdw.Id.Get()}, hcsdw)
}

// UpdateHrContractSignDocumentWizards updates existing hr.contract.sign.document.wizard records.
// All records (represented by ids) will be updated by hcsdw values.
func (c *Client) UpdateHrContractSignDocumentWizards(ids []int64, hcsdw *HrContractSignDocumentWizard) error {
	return c.Update(HrContractSignDocumentWizardModel, ids, hcsdw, nil)
}

// DeleteHrContractSignDocumentWizard deletes an existing hr.contract.sign.document.wizard record.
func (c *Client) DeleteHrContractSignDocumentWizard(id int64) error {
	return c.DeleteHrContractSignDocumentWizards([]int64{id})
}

// DeleteHrContractSignDocumentWizards deletes existing hr.contract.sign.document.wizard records.
func (c *Client) DeleteHrContractSignDocumentWizards(ids []int64) error {
	return c.Delete(HrContractSignDocumentWizardModel, ids)
}

// GetHrContractSignDocumentWizard gets hr.contract.sign.document.wizard existing record.
func (c *Client) GetHrContractSignDocumentWizard(id int64) (*HrContractSignDocumentWizard, error) {
	hcsdws, err := c.GetHrContractSignDocumentWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hcsdws)[0]), nil
}

// GetHrContractSignDocumentWizards gets hr.contract.sign.document.wizard existing records.
func (c *Client) GetHrContractSignDocumentWizards(ids []int64) (*HrContractSignDocumentWizards, error) {
	hcsdws := &HrContractSignDocumentWizards{}
	if err := c.Read(HrContractSignDocumentWizardModel, ids, nil, hcsdws); err != nil {
		return nil, err
	}
	return hcsdws, nil
}

// FindHrContractSignDocumentWizard finds hr.contract.sign.document.wizard record by querying it with criteria.
func (c *Client) FindHrContractSignDocumentWizard(criteria *Criteria) (*HrContractSignDocumentWizard, error) {
	hcsdws := &HrContractSignDocumentWizards{}
	if err := c.SearchRead(HrContractSignDocumentWizardModel, criteria, NewOptions().Limit(1), hcsdws); err != nil {
		return nil, err
	}
	return &((*hcsdws)[0]), nil
}

// FindHrContractSignDocumentWizards finds hr.contract.sign.document.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractSignDocumentWizards(criteria *Criteria, options *Options) (*HrContractSignDocumentWizards, error) {
	hcsdws := &HrContractSignDocumentWizards{}
	if err := c.SearchRead(HrContractSignDocumentWizardModel, criteria, options, hcsdws); err != nil {
		return nil, err
	}
	return hcsdws, nil
}

// FindHrContractSignDocumentWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractSignDocumentWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrContractSignDocumentWizardModel, criteria, options)
}

// FindHrContractSignDocumentWizardId finds record id by querying it with criteria.
func (c *Client) FindHrContractSignDocumentWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractSignDocumentWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
