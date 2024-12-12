package odoo

// PosPreparationDisplayStage represents pos_preparation_display.stage model.
type PosPreparationDisplayStage struct {
	AlertTimer           *Int      `xmlrpc:"alert_timer,omitempty"`
	Color                *String   `xmlrpc:"color,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	Name                 *String   `xmlrpc:"name,omitempty"`
	PreparationDisplayId *Many2One `xmlrpc:"preparation_display_id,omitempty"`
	Sequence             *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPreparationDisplayStages represents array of pos_preparation_display.stage model.
type PosPreparationDisplayStages []PosPreparationDisplayStage

// PosPreparationDisplayStageModel is the odoo model name.
const PosPreparationDisplayStageModel = "pos_preparation_display.stage"

// Many2One convert PosPreparationDisplayStage to *Many2One.
func (ps *PosPreparationDisplayStage) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePosPreparationDisplayStage creates a new pos_preparation_display.stage model and returns its id.
func (c *Client) CreatePosPreparationDisplayStage(ps *PosPreparationDisplayStage) (int64, error) {
	ids, err := c.CreatePosPreparationDisplayStages([]*PosPreparationDisplayStage{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPreparationDisplayStage creates a new pos_preparation_display.stage model and returns its id.
func (c *Client) CreatePosPreparationDisplayStages(pss []*PosPreparationDisplayStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(PosPreparationDisplayStageModel, vv, nil)
}

// UpdatePosPreparationDisplayStage updates an existing pos_preparation_display.stage record.
func (c *Client) UpdatePosPreparationDisplayStage(ps *PosPreparationDisplayStage) error {
	return c.UpdatePosPreparationDisplayStages([]int64{ps.Id.Get()}, ps)
}

// UpdatePosPreparationDisplayStages updates existing pos_preparation_display.stage records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePosPreparationDisplayStages(ids []int64, ps *PosPreparationDisplayStage) error {
	return c.Update(PosPreparationDisplayStageModel, ids, ps, nil)
}

// DeletePosPreparationDisplayStage deletes an existing pos_preparation_display.stage record.
func (c *Client) DeletePosPreparationDisplayStage(id int64) error {
	return c.DeletePosPreparationDisplayStages([]int64{id})
}

// DeletePosPreparationDisplayStages deletes existing pos_preparation_display.stage records.
func (c *Client) DeletePosPreparationDisplayStages(ids []int64) error {
	return c.Delete(PosPreparationDisplayStageModel, ids)
}

// GetPosPreparationDisplayStage gets pos_preparation_display.stage existing record.
func (c *Client) GetPosPreparationDisplayStage(id int64) (*PosPreparationDisplayStage, error) {
	pss, err := c.GetPosPreparationDisplayStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// GetPosPreparationDisplayStages gets pos_preparation_display.stage existing records.
func (c *Client) GetPosPreparationDisplayStages(ids []int64) (*PosPreparationDisplayStages, error) {
	pss := &PosPreparationDisplayStages{}
	if err := c.Read(PosPreparationDisplayStageModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosPreparationDisplayStage finds pos_preparation_display.stage record by querying it with criteria.
func (c *Client) FindPosPreparationDisplayStage(criteria *Criteria) (*PosPreparationDisplayStage, error) {
	pss := &PosPreparationDisplayStages{}
	if err := c.SearchRead(PosPreparationDisplayStageModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// FindPosPreparationDisplayStages finds pos_preparation_display.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayStages(criteria *Criteria, options *Options) (*PosPreparationDisplayStages, error) {
	pss := &PosPreparationDisplayStages{}
	if err := c.SearchRead(PosPreparationDisplayStageModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosPreparationDisplayStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPreparationDisplayStageModel, criteria, options)
}

// FindPosPreparationDisplayStageId finds record id by querying it with criteria.
func (c *Client) FindPosPreparationDisplayStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPreparationDisplayStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
