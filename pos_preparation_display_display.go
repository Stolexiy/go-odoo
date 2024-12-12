package odoo

// PosPreparationDisplayDisplay represents pos_preparation_display.display model.
type PosPreparationDisplayDisplay struct {
	AccessToken           *String   `xmlrpc:"access_token,omitempty"`
	AverageTime           *Int      `xmlrpc:"average_time,omitempty"`
	CategoryIds           *Relation `xmlrpc:"category_ids,omitempty"`
	CompanyId             *Many2One `xmlrpc:"company_id,omitempty"`
	ContainsBarRestaurant *Bool     `xmlrpc:"contains_bar_restaurant,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	Name                  *String   `xmlrpc:"name,omitempty"`
	OrderCount            *Int      `xmlrpc:"order_count,omitempty"`
	PosConfigIds          *Relation `xmlrpc:"pos_config_ids,omitempty"`
	StageIds              *Relation `xmlrpc:"stage_ids,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosPreparationDisplayDisplays represents array of pos_preparation_display.display model.
type PosPreparationDisplayDisplays []PosPreparationDisplayDisplay

// PosPreparationDisplayDisplayModel is the odoo model name.
const PosPreparationDisplayDisplayModel = "pos_preparation_display.display"

// Many2One convert PosPreparationDisplayDisplay to *Many2One.
func (pd *PosPreparationDisplayDisplay) Many2One() *Many2One {
	return NewMany2One(pd.Id.Get(), "")
}

// CreatePosPreparationDisplayDisplay creates a new pos_preparation_display.display model and returns its id.
func (c *Client) CreatePosPreparationDisplayDisplay(pd *PosPreparationDisplayDisplay) (int64, error) {
	ids, err := c.CreatePosPreparationDisplayDisplays([]*PosPreparationDisplayDisplay{pd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosPreparationDisplayDisplay creates a new pos_preparation_display.display model and returns its id.
func (c *Client) CreatePosPreparationDisplayDisplays(pds []*PosPreparationDisplayDisplay) ([]int64, error) {
	var vv []interface{}
	for _, v := range pds {
		vv = append(vv, v)
	}
	return c.Create(PosPreparationDisplayDisplayModel, vv, nil)
}

// UpdatePosPreparationDisplayDisplay updates an existing pos_preparation_display.display record.
func (c *Client) UpdatePosPreparationDisplayDisplay(pd *PosPreparationDisplayDisplay) error {
	return c.UpdatePosPreparationDisplayDisplays([]int64{pd.Id.Get()}, pd)
}

// UpdatePosPreparationDisplayDisplays updates existing pos_preparation_display.display records.
// All records (represented by ids) will be updated by pd values.
func (c *Client) UpdatePosPreparationDisplayDisplays(ids []int64, pd *PosPreparationDisplayDisplay) error {
	return c.Update(PosPreparationDisplayDisplayModel, ids, pd, nil)
}

// DeletePosPreparationDisplayDisplay deletes an existing pos_preparation_display.display record.
func (c *Client) DeletePosPreparationDisplayDisplay(id int64) error {
	return c.DeletePosPreparationDisplayDisplays([]int64{id})
}

// DeletePosPreparationDisplayDisplays deletes existing pos_preparation_display.display records.
func (c *Client) DeletePosPreparationDisplayDisplays(ids []int64) error {
	return c.Delete(PosPreparationDisplayDisplayModel, ids)
}

// GetPosPreparationDisplayDisplay gets pos_preparation_display.display existing record.
func (c *Client) GetPosPreparationDisplayDisplay(id int64) (*PosPreparationDisplayDisplay, error) {
	pds, err := c.GetPosPreparationDisplayDisplays([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pds)[0]), nil
}

// GetPosPreparationDisplayDisplays gets pos_preparation_display.display existing records.
func (c *Client) GetPosPreparationDisplayDisplays(ids []int64) (*PosPreparationDisplayDisplays, error) {
	pds := &PosPreparationDisplayDisplays{}
	if err := c.Read(PosPreparationDisplayDisplayModel, ids, nil, pds); err != nil {
		return nil, err
	}
	return pds, nil
}

// FindPosPreparationDisplayDisplay finds pos_preparation_display.display record by querying it with criteria.
func (c *Client) FindPosPreparationDisplayDisplay(criteria *Criteria) (*PosPreparationDisplayDisplay, error) {
	pds := &PosPreparationDisplayDisplays{}
	if err := c.SearchRead(PosPreparationDisplayDisplayModel, criteria, NewOptions().Limit(1), pds); err != nil {
		return nil, err
	}
	return &((*pds)[0]), nil
}

// FindPosPreparationDisplayDisplays finds pos_preparation_display.display records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayDisplays(criteria *Criteria, options *Options) (*PosPreparationDisplayDisplays, error) {
	pds := &PosPreparationDisplayDisplays{}
	if err := c.SearchRead(PosPreparationDisplayDisplayModel, criteria, options, pds); err != nil {
		return nil, err
	}
	return pds, nil
}

// FindPosPreparationDisplayDisplayIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPreparationDisplayDisplayIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosPreparationDisplayDisplayModel, criteria, options)
}

// FindPosPreparationDisplayDisplayId finds record id by querying it with criteria.
func (c *Client) FindPosPreparationDisplayDisplayId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPreparationDisplayDisplayModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
