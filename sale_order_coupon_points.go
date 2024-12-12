package odoo

// SaleOrderCouponPoints represents sale.order.coupon.points model.
type SaleOrderCouponPoints struct {
	CouponId    *Many2One `xmlrpc:"coupon_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	OrderId     *Many2One `xmlrpc:"order_id,omitempty"`
	Points      *Float    `xmlrpc:"points,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderCouponPointss represents array of sale.order.coupon.points model.
type SaleOrderCouponPointss []SaleOrderCouponPoints

// SaleOrderCouponPointsModel is the odoo model name.
const SaleOrderCouponPointsModel = "sale.order.coupon.points"

// Many2One convert SaleOrderCouponPoints to *Many2One.
func (socp *SaleOrderCouponPoints) Many2One() *Many2One {
	return NewMany2One(socp.Id.Get(), "")
}

// CreateSaleOrderCouponPoints creates a new sale.order.coupon.points model and returns its id.
func (c *Client) CreateSaleOrderCouponPoints(socp *SaleOrderCouponPoints) (int64, error) {
	ids, err := c.CreateSaleOrderCouponPointss([]*SaleOrderCouponPoints{socp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderCouponPoints creates a new sale.order.coupon.points model and returns its id.
func (c *Client) CreateSaleOrderCouponPointss(socps []*SaleOrderCouponPoints) ([]int64, error) {
	var vv []interface{}
	for _, v := range socps {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderCouponPointsModel, vv, nil)
}

// UpdateSaleOrderCouponPoints updates an existing sale.order.coupon.points record.
func (c *Client) UpdateSaleOrderCouponPoints(socp *SaleOrderCouponPoints) error {
	return c.UpdateSaleOrderCouponPointss([]int64{socp.Id.Get()}, socp)
}

// UpdateSaleOrderCouponPointss updates existing sale.order.coupon.points records.
// All records (represented by ids) will be updated by socp values.
func (c *Client) UpdateSaleOrderCouponPointss(ids []int64, socp *SaleOrderCouponPoints) error {
	return c.Update(SaleOrderCouponPointsModel, ids, socp, nil)
}

// DeleteSaleOrderCouponPoints deletes an existing sale.order.coupon.points record.
func (c *Client) DeleteSaleOrderCouponPoints(id int64) error {
	return c.DeleteSaleOrderCouponPointss([]int64{id})
}

// DeleteSaleOrderCouponPointss deletes existing sale.order.coupon.points records.
func (c *Client) DeleteSaleOrderCouponPointss(ids []int64) error {
	return c.Delete(SaleOrderCouponPointsModel, ids)
}

// GetSaleOrderCouponPoints gets sale.order.coupon.points existing record.
func (c *Client) GetSaleOrderCouponPoints(id int64) (*SaleOrderCouponPoints, error) {
	socps, err := c.GetSaleOrderCouponPointss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*socps)[0]), nil
}

// GetSaleOrderCouponPointss gets sale.order.coupon.points existing records.
func (c *Client) GetSaleOrderCouponPointss(ids []int64) (*SaleOrderCouponPointss, error) {
	socps := &SaleOrderCouponPointss{}
	if err := c.Read(SaleOrderCouponPointsModel, ids, nil, socps); err != nil {
		return nil, err
	}
	return socps, nil
}

// FindSaleOrderCouponPoints finds sale.order.coupon.points record by querying it with criteria.
func (c *Client) FindSaleOrderCouponPoints(criteria *Criteria) (*SaleOrderCouponPoints, error) {
	socps := &SaleOrderCouponPointss{}
	if err := c.SearchRead(SaleOrderCouponPointsModel, criteria, NewOptions().Limit(1), socps); err != nil {
		return nil, err
	}
	return &((*socps)[0]), nil
}

// FindSaleOrderCouponPointss finds sale.order.coupon.points records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCouponPointss(criteria *Criteria, options *Options) (*SaleOrderCouponPointss, error) {
	socps := &SaleOrderCouponPointss{}
	if err := c.SearchRead(SaleOrderCouponPointsModel, criteria, options, socps); err != nil {
		return nil, err
	}
	return socps, nil
}

// FindSaleOrderCouponPointsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCouponPointsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderCouponPointsModel, criteria, options)
}

// FindSaleOrderCouponPointsId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderCouponPointsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderCouponPointsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
