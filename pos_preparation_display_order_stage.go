package odoo

// PosPreparationDisplayOrderStage represents pos_preparation_display.order.stage model.
type PosPreparationDisplayOrderStage struct {
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Done                 *Bool     `xmlrpc:"done,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	OrderId              *Many2One `xmlrpc:"order_id,omitempty"`
	PreparationDisplayId *Many2One `xmlrpc:"preparation_display_id,omitempty"`
	StageId              *Many2One `xmlrpc:"stage_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPreparationDisplayOrderStages represents array of pos_preparation_display.order.stage model.
type PosPreparationDisplayOrderStages []PosPreparationDisplayOrderStage

// PosPreparationDisplayOrderStageModel is the odoo model name.
const PosPreparationDisplayOrderStageModel = "pos_preparation_display.order.stage"

// Many2One convert PosPreparationDisplayOrderStage to *Many2One.
func (pos *PosPreparationDisplayOrderStage) Many2One() *Many2One {
	return NewMany2One(pos.Id.Get(), "")
}

// CreatePosPreparationDisplayOrderStage creates a new pos_preparation_display.order.stage model and returns its id.
func (c *Client) CreatePosPreparationDisplayOrderStage(pos *PosPreparationDisplayOrderStage) (int64, error) {
	ids, err := c.CreatePosPreparationDisplayOrderStages([]*PosPreparationDisplayOrderStage{pos})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPreparationDisplayOrderStage creates a new pos_preparation_display.order.stage model and returns its id.
func (c *Client) CreatePosPreparationDisplayOrderStages(poss []*PosPreparationDisplayOrderStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range poss {
		vv = append(vv, v)
	}
	return c.Create(PosPreparationDisplayOrderStageModel, vv, nil)
}

// UpdatePosPreparationDisplayOrderStage updates an existing pos_preparation_display.order.stage record.
func (c *Client) UpdatePosPreparationDisplayOrderStage(pos *PosPreparationDisplayOrderStage) error {
	return c.UpdatePosPreparationDisplayOrderStages([]int64{pos.Id.Get()}, pos)
}

// UpdatePosPreparationDisplayOrderStages updates existing pos_preparation_display.order.stage records.
// All records (represented by ids) will be updated by pos values.
func (c *Client) UpdatePosPreparationDisplayOrderStages(ids []int64, pos *PosPreparationDisplayOrderStage) error {
	return c.Update(PosPreparationDisplayOrderStageModel, ids, pos, nil)
}

// DeletePosPreparationDisplayOrderStage deletes an existing pos_preparation_display.order.stage record.
func (c *Client) DeletePosPreparationDisplayOrderStage(id int64) error {
	return c.DeletePosPreparationDisplayOrderStages([]int64{id})
}

// DeletePosPreparationDisplayOrderStages deletes existing pos_preparation_display.order.stage records.
func (c *Client) DeletePosPreparationDisplayOrderStages(ids []int64) error {
	return c.Delete(PosPreparationDisplayOrderStageModel, ids)
}

// GetPosPreparationDisplayOrderStage gets pos_preparation_display.order.stage existing record.
func (c *Client) GetPosPreparationDisplayOrderStage(id int64) (*PosPreparationDisplayOrderStage, error) {
	poss, err := c.GetPosPreparationDisplayOrderStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*poss)[0]), nil
}

// GetPosPreparationDisplayOrderStages gets pos_preparation_display.order.stage existing records.
func (c *Client) GetPosPreparationDisplayOrderStages(ids []int64) (*PosPreparationDisplayOrderStages, error) {
	poss := &PosPreparationDisplayOrderStages{}
	if err := c.Read(PosPreparationDisplayOrderStageModel, ids, nil, poss); err != nil {
		return nil, err
	}
	return poss, nil
}

// FindPosPreparationDisplayOrderStage finds pos_preparation_display.order.stage record by querying it with criteria.
func (c *Client) FindPosPreparationDisplayOrderStage(criteria *Criteria) (*PosPreparationDisplayOrderStage, error) {
	poss := &PosPreparationDisplayOrderStages{}
	if err := c.SearchRead(PosPreparationDisplayOrderStageModel, criteria, NewOptions().Limit(1), poss); err != nil {
		return nil, err
	}
	return &((*poss)[0]), nil
}

// FindPosPreparationDisplayOrderStages finds pos_preparation_display.order.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayOrderStages(criteria *Criteria, options *Options) (*PosPreparationDisplayOrderStages, error) {
	poss := &PosPreparationDisplayOrderStages{}
	if err := c.SearchRead(PosPreparationDisplayOrderStageModel, criteria, options, poss); err != nil {
		return nil, err
	}
	return poss, nil
}

// FindPosPreparationDisplayOrderStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayOrderStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPreparationDisplayOrderStageModel, criteria, options)
}

// FindPosPreparationDisplayOrderStageId finds record id by querying it with criteria.
func (c *Client) FindPosPreparationDisplayOrderStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPreparationDisplayOrderStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
