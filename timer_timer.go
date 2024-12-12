package odoo

// TimerTimer represents timer.timer model.
type TimerTimer struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	IsTimerRunning *Bool     `xmlrpc:"is_timer_running,omitempty"`
	ResId          *Int      `xmlrpc:"res_id,omitempty"`
	ResModel       *String   `xmlrpc:"res_model,omitempty"`
	TimerPause     *Time     `xmlrpc:"timer_pause,omitempty"`
	TimerStart     *Time     `xmlrpc:"timer_start,omitempty"`
	UserId         *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// TimerTimers represents array of timer.timer model.
type TimerTimers []TimerTimer

// TimerTimerModel is the odoo model name.
const TimerTimerModel = "timer.timer"

// Many2One convert TimerTimer to *Many2One.
func (tt *TimerTimer) Many2One() *Many2One {
	return NewMany2One(tt.Id.Get(), "")
}

// CreateTimerTimer creates a new timer.timer model and returns its id.
func (c *Client) CreateTimerTimer(tt *TimerTimer) (int64, error) {
	ids, err := c.CreateTimerTimers([]*TimerTimer{tt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimerTimer creates a new timer.timer model and returns its id.
func (c *Client) CreateTimerTimers(tts []*TimerTimer) ([]int64, error) {
	var vv []interface{}
	for _, v := range tts {
		vv = append(vv, v)
	}
	return c.Create(TimerTimerModel, vv, nil)
}

// UpdateTimerTimer updates an existing timer.timer record.
func (c *Client) UpdateTimerTimer(tt *TimerTimer) error {
	return c.UpdateTimerTimers([]int64{tt.Id.Get()}, tt)
}

// UpdateTimerTimers updates existing timer.timer records.
// All records (represented by ids) will be updated by tt values.
func (c *Client) UpdateTimerTimers(ids []int64, tt *TimerTimer) error {
	return c.Update(TimerTimerModel, ids, tt, nil)
}

// DeleteTimerTimer deletes an existing timer.timer record.
func (c *Client) DeleteTimerTimer(id int64) error {
	return c.DeleteTimerTimers([]int64{id})
}

// DeleteTimerTimers deletes existing timer.timer records.
func (c *Client) DeleteTimerTimers(ids []int64) error {
	return c.Delete(TimerTimerModel, ids)
}

// GetTimerTimer gets timer.timer existing record.
func (c *Client) GetTimerTimer(id int64) (*TimerTimer, error) {
	tts, err := c.GetTimerTimers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tts)[0]), nil
}

// GetTimerTimers gets timer.timer existing records.
func (c *Client) GetTimerTimers(ids []int64) (*TimerTimers, error) {
	tts := &TimerTimers{}
	if err := c.Read(TimerTimerModel, ids, nil, tts); err != nil {
		return nil, err
	}
	return tts, nil
}

// FindTimerTimer finds timer.timer record by querying it with criteria.
func (c *Client) FindTimerTimer(criteria *Criteria) (*TimerTimer, error) {
	tts := &TimerTimers{}
	if err := c.SearchRead(TimerTimerModel, criteria, NewOptions().Limit(1), tts); err != nil {
		return nil, err
	}
	return &((*tts)[0]), nil
}

// FindTimerTimers finds timer.timer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimerTimers(criteria *Criteria, options *Options) (*TimerTimers, error) {
	tts := &TimerTimers{}
	if err := c.SearchRead(TimerTimerModel, criteria, options, tts); err != nil {
		return nil, err
	}
	return tts, nil
}

// FindTimerTimerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimerTimerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TimerTimerModel, criteria, options)
}

// FindTimerTimerId finds record id by querying it with criteria.
func (c *Client) FindTimerTimerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimerTimerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
