package odoo

// PosNote represents pos.note model.
type PosNote struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosNotes represents array of pos.note model.
type PosNotes []PosNote

// PosNoteModel is the odoo model name.
const PosNoteModel = "pos.note"

// Many2One convert PosNote to *Many2One.
func (pn *PosNote) Many2One() *Many2One {
	return NewMany2One(pn.Id.Get(), "")
}

// CreatePosNote creates a new pos.note model and returns its id.
func (c *Client) CreatePosNote(pn *PosNote) (int64, error) {
	ids, err := c.CreatePosNotes([]*PosNote{pn})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosNote creates a new pos.note model and returns its id.
func (c *Client) CreatePosNotes(pns []*PosNote) ([]int64, error) {
	var vv []interface{}
	for _, v := range pns {
		vv = append(vv, v)
	}
	return c.Create(PosNoteModel, vv, nil)
}

// UpdatePosNote updates an existing pos.note record.
func (c *Client) UpdatePosNote(pn *PosNote) error {
	return c.UpdatePosNotes([]int64{pn.Id.Get()}, pn)
}

// UpdatePosNotes updates existing pos.note records.
// All records (represented by ids) will be updated by pn values.
func (c *Client) UpdatePosNotes(ids []int64, pn *PosNote) error {
	return c.Update(PosNoteModel, ids, pn, nil)
}

// DeletePosNote deletes an existing pos.note record.
func (c *Client) DeletePosNote(id int64) error {
	return c.DeletePosNotes([]int64{id})
}

// DeletePosNotes deletes existing pos.note records.
func (c *Client) DeletePosNotes(ids []int64) error {
	return c.Delete(PosNoteModel, ids)
}

// GetPosNote gets pos.note existing record.
func (c *Client) GetPosNote(id int64) (*PosNote, error) {
	pns, err := c.GetPosNotes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pns)[0]), nil
}

// GetPosNotes gets pos.note existing records.
func (c *Client) GetPosNotes(ids []int64) (*PosNotes, error) {
	pns := &PosNotes{}
	if err := c.Read(PosNoteModel, ids, nil, pns); err != nil {
		return nil, err
	}
	return pns, nil
}

// FindPosNote finds pos.note record by querying it with criteria.
func (c *Client) FindPosNote(criteria *Criteria) (*PosNote, error) {
	pns := &PosNotes{}
	if err := c.SearchRead(PosNoteModel, criteria, NewOptions().Limit(1), pns); err != nil {
		return nil, err
	}
	return &((*pns)[0]), nil
}

// FindPosNotes finds pos.note records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosNotes(criteria *Criteria, options *Options) (*PosNotes, error) {
	pns := &PosNotes{}
	if err := c.SearchRead(PosNoteModel, criteria, options, pns); err != nil {
		return nil, err
	}
	return pns, nil
}

// FindPosNoteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosNoteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosNoteModel, criteria, options)
}

// FindPosNoteId finds record id by querying it with criteria.
func (c *Client) FindPosNoteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosNoteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
