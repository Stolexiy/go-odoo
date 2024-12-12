package odoo

// AsyncRequestModel represents async.request.model model.
type AsyncRequestModel struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AsyncRequestModels represents array of async.request.model model.
type AsyncRequestModels []AsyncRequestModel

// AsyncRequestModelModel is the odoo model name.
const AsyncRequestModelModel = "async.request.model"

// Many2One convert AsyncRequestModel to *Many2One.
func (arm *AsyncRequestModel) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAsyncRequestModel creates a new async.request.model model and returns its id.
func (c *Client) CreateAsyncRequestModel(arm *AsyncRequestModel) (int64, error) {
	ids, err := c.CreateAsyncRequestModels([]*AsyncRequestModel{arm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAsyncRequestModel creates a new async.request.model model and returns its id.
func (c *Client) CreateAsyncRequestModels(arms []*AsyncRequestModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range arms {
		vv = append(vv, v)
	}
	return c.Create(AsyncRequestModelModel, vv, nil)
}

// UpdateAsyncRequestModel updates an existing async.request.model record.
func (c *Client) UpdateAsyncRequestModel(arm *AsyncRequestModel) error {
	return c.UpdateAsyncRequestModels([]int64{arm.Id.Get()}, arm)
}

// UpdateAsyncRequestModels updates existing async.request.model records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAsyncRequestModels(ids []int64, arm *AsyncRequestModel) error {
	return c.Update(AsyncRequestModelModel, ids, arm, nil)
}

// DeleteAsyncRequestModel deletes an existing async.request.model record.
func (c *Client) DeleteAsyncRequestModel(id int64) error {
	return c.DeleteAsyncRequestModels([]int64{id})
}

// DeleteAsyncRequestModels deletes existing async.request.model records.
func (c *Client) DeleteAsyncRequestModels(ids []int64) error {
	return c.Delete(AsyncRequestModelModel, ids)
}

// GetAsyncRequestModel gets async.request.model existing record.
func (c *Client) GetAsyncRequestModel(id int64) (*AsyncRequestModel, error) {
	arms, err := c.GetAsyncRequestModels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// GetAsyncRequestModels gets async.request.model existing records.
func (c *Client) GetAsyncRequestModels(ids []int64) (*AsyncRequestModels, error) {
	arms := &AsyncRequestModels{}
	if err := c.Read(AsyncRequestModelModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAsyncRequestModel finds async.request.model record by querying it with criteria.
func (c *Client) FindAsyncRequestModel(criteria *Criteria) (*AsyncRequestModel, error) {
	arms := &AsyncRequestModels{}
	if err := c.SearchRead(AsyncRequestModelModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// FindAsyncRequestModels finds async.request.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAsyncRequestModels(criteria *Criteria, options *Options) (*AsyncRequestModels, error) {
	arms := &AsyncRequestModels{}
	if err := c.SearchRead(AsyncRequestModelModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAsyncRequestModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAsyncRequestModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AsyncRequestModelModel, criteria, options)
}

// FindAsyncRequestModelId finds record id by querying it with criteria.
func (c *Client) FindAsyncRequestModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AsyncRequestModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
