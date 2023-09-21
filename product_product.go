package odoo

import (
	"fmt"
)

// ProductProduct represents product.product model.
type ProductProduct struct {
	LastUpdate                             *Time      `xmlrpc:"__last_update,omptempty"`
	AccountTagIds                          *Relation  `xmlrpc:"account_tag_ids,omptempty"`
	Active                                 *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline                   *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration            *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                  *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                            *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                          *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                        *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                       *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                         *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                         *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalProductTagIds                *Relation  `xmlrpc:"additional_product_tag_ids,omptempty"`
	AllProductTagIds                       *Relation  `xmlrpc:"all_product_tag_ids,omptempty"`
	AttributeLineIds                       *Relation  `xmlrpc:"attribute_line_ids,omptempty"`
	AvailableInPos                         *Bool      `xmlrpc:"available_in_pos,omptempty"`
	AvgCost                                *Float     `xmlrpc:"avg_cost,omptempty"`
	Barcode                                *String    `xmlrpc:"barcode,omptempty"`
	CanBeExpensed                          *Bool      `xmlrpc:"can_be_expensed,omptempty"`
	CanImage1024BeZoomed                   *Bool      `xmlrpc:"can_image_1024_be_zoomed,omptempty"`
	CanImageVariant1024BeZoomed            *Bool      `xmlrpc:"can_image_variant_1024_be_zoomed,omptempty"`
	CategId                                *Many2One  `xmlrpc:"categ_id,omptempty"`
	Code                                   *String    `xmlrpc:"code,omptempty"`
	Color                                  *Int       `xmlrpc:"color,omptempty"`
	CombinationIndices                     *String    `xmlrpc:"combination_indices,omptempty"`
	CompanyCurrencyId                      *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                              *Many2One  `xmlrpc:"company_id,omptempty"`
	CostCurrencyId                         *Many2One  `xmlrpc:"cost_currency_id,omptempty"`
	CostMethod                             *Selection `xmlrpc:"cost_method,omptempty"`
	CreateDate                             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                             *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateFrom                               *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                                 *Time      `xmlrpc:"date_to,omptempty"`
	DefaultCode                            *String    `xmlrpc:"default_code,omptempty"`
	Description                            *String    `xmlrpc:"description,omptempty"`
	DescriptionPicking                     *String    `xmlrpc:"description_picking,omptempty"`
	DescriptionPickingin                   *String    `xmlrpc:"description_pickingin,omptempty"`
	DescriptionPickingout                  *String    `xmlrpc:"description_pickingout,omptempty"`
	DescriptionPurchase                    *String    `xmlrpc:"description_purchase,omptempty"`
	DescriptionSale                        *String    `xmlrpc:"description_sale,omptempty"`
	DetailedType                           *Selection `xmlrpc:"detailed_type,omptempty"`
	DisplayName                            *String    `xmlrpc:"display_name,omptempty"`
	ExpectedMargin                         *Float     `xmlrpc:"expected_margin,omptempty"`
	ExpectedMarginRate                     *Float     `xmlrpc:"expected_margin_rate,omptempty"`
	ExpensePolicy                          *Selection `xmlrpc:"expense_policy,omptempty"`
	ExpensePolicyTooltip                   *String    `xmlrpc:"expense_policy_tooltip,omptempty"`
	FreeQty                                *Float     `xmlrpc:"free_qty,omptempty"`
	HasAvailableRouteIds                   *Bool      `xmlrpc:"has_available_route_ids,omptempty"`
	HasConfigurableAttributes              *Bool      `xmlrpc:"has_configurable_attributes,omptempty"`
	HasMessage                             *Bool      `xmlrpc:"has_message,omptempty"`
	Id                                     *Int       `xmlrpc:"id,omptempty"`
	Image1024                              *String    `xmlrpc:"image_1024,omptempty"`
	Image128                               *String    `xmlrpc:"image_128,omptempty"`
	Image1920                              *String    `xmlrpc:"image_1920,omptempty"`
	Image256                               *String    `xmlrpc:"image_256,omptempty"`
	Image512                               *String    `xmlrpc:"image_512,omptempty"`
	ImageVariant1024                       *String    `xmlrpc:"image_variant_1024,omptempty"`
	ImageVariant128                        *String    `xmlrpc:"image_variant_128,omptempty"`
	ImageVariant1920                       *String    `xmlrpc:"image_variant_1920,omptempty"`
	ImageVariant256                        *String    `xmlrpc:"image_variant_256,omptempty"`
	ImageVariant512                        *String    `xmlrpc:"image_variant_512,omptempty"`
	IncomingQty                            *Float     `xmlrpc:"incoming_qty,omptempty"`
	InvoicePolicy                          *Selection `xmlrpc:"invoice_policy,omptempty"`
	InvoiceState                           *Selection `xmlrpc:"invoice_state,omptempty"`
	IsProductVariant                       *Bool      `xmlrpc:"is_product_variant,omptempty"`
	ListPrice                              *Float     `xmlrpc:"list_price,omptempty"`
	LocationId                             *Many2One  `xmlrpc:"location_id,omptempty"`
	LstPrice                               *Float     `xmlrpc:"lst_price,omptempty"`
	MessageAttachmentCount                 *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds                     *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                        *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                 *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                     *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                             *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                      *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId                *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                      *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter               *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                      *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline                 *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                   *String    `xmlrpc:"name,omptempty"`
	NbrMovesIn                             *Int       `xmlrpc:"nbr_moves_in,omptempty"`
	NbrMovesOut                            *Int       `xmlrpc:"nbr_moves_out,omptempty"`
	NbrReorderingRules                     *Int       `xmlrpc:"nbr_reordering_rules,omptempty"`
	NormalCost                             *Float     `xmlrpc:"normal_cost,omptempty"`
	OptionalProductIds                     *Relation  `xmlrpc:"optional_product_ids,omptempty"`
	OrderpointIds                          *Relation  `xmlrpc:"orderpoint_ids,omptempty"`
	OutgoingQty                            *Float     `xmlrpc:"outgoing_qty,omptempty"`
	PackagingIds                           *Relation  `xmlrpc:"packaging_ids,omptempty"`
	PartnerRef                             *String    `xmlrpc:"partner_ref,omptempty"`
	PlanningEnabled                        *Bool      `xmlrpc:"planning_enabled,omptempty"`
	PlanningRoleId                         *Many2One  `xmlrpc:"planning_role_id,omptempty"`
	PosCategId                             *Many2One  `xmlrpc:"pos_categ_id,omptempty"`
	PriceExtra                             *Float     `xmlrpc:"price_extra,omptempty"`
	PricelistItemCount                     *Int       `xmlrpc:"pricelist_item_count,omptempty"`
	Priority                               *Selection `xmlrpc:"priority,omptempty"`
	ProductTagIds                          *Relation  `xmlrpc:"product_tag_ids,omptempty"`
	ProductTemplateAttributeValueIds       *Relation  `xmlrpc:"product_template_attribute_value_ids,omptempty"`
	ProductTemplateVariantValueIds         *Relation  `xmlrpc:"product_template_variant_value_ids,omptempty"`
	ProductTmplId                          *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductTooltip                         *String    `xmlrpc:"product_tooltip,omptempty"`
	ProductVariantCount                    *Int       `xmlrpc:"product_variant_count,omptempty"`
	ProductVariantId                       *Many2One  `xmlrpc:"product_variant_id,omptempty"`
	ProductVariantIds                      *Relation  `xmlrpc:"product_variant_ids,omptempty"`
	PropertyAccountCreditorPriceDifference *Many2One  `xmlrpc:"property_account_creditor_price_difference,omptempty"`
	PropertyAccountExpenseId               *Many2One  `xmlrpc:"property_account_expense_id,omptempty"`
	PropertyAccountIncomeId                *Many2One  `xmlrpc:"property_account_income_id,omptempty"`
	PropertyStockInventory                 *Many2One  `xmlrpc:"property_stock_inventory,omptempty"`
	PropertyStockProduction                *Many2One  `xmlrpc:"property_stock_production,omptempty"`
	PurchaseAvgPrice                       *Float     `xmlrpc:"purchase_avg_price,omptempty"`
	PurchaseGap                            *Float     `xmlrpc:"purchase_gap,omptempty"`
	PurchaseLineWarn                       *Selection `xmlrpc:"purchase_line_warn,omptempty"`
	PurchaseLineWarnMsg                    *String    `xmlrpc:"purchase_line_warn_msg,omptempty"`
	PurchaseMethod                         *Selection `xmlrpc:"purchase_method,omptempty"`
	PurchaseNumInvoiced                    *Float     `xmlrpc:"purchase_num_invoiced,omptempty"`
	PurchaseOk                             *Bool      `xmlrpc:"purchase_ok,omptempty"`
	PurchaseOrderLineIds                   *Relation  `xmlrpc:"purchase_order_line_ids,omptempty"`
	PurchasedProductQty                    *Float     `xmlrpc:"purchased_product_qty,omptempty"`
	PutawayRuleIds                         *Relation  `xmlrpc:"putaway_rule_ids,omptempty"`
	QtyAvailable                           *Float     `xmlrpc:"qty_available,omptempty"`
	QuantitySvl                            *Float     `xmlrpc:"quantity_svl,omptempty"`
	ReorderingMaxQty                       *Float     `xmlrpc:"reordering_max_qty,omptempty"`
	ReorderingMinQty                       *Float     `xmlrpc:"reordering_min_qty,omptempty"`
	ResponsibleId                          *Many2One  `xmlrpc:"responsible_id,omptempty"`
	RouteFromCategIds                      *Relation  `xmlrpc:"route_from_categ_ids,omptempty"`
	RouteIds                               *Relation  `xmlrpc:"route_ids,omptempty"`
	SaleAvgPrice                           *Float     `xmlrpc:"sale_avg_price,omptempty"`
	SaleDelay                              *Float     `xmlrpc:"sale_delay,omptempty"`
	SaleExpected                           *Float     `xmlrpc:"sale_expected,omptempty"`
	SaleLineWarn                           *Selection `xmlrpc:"sale_line_warn,omptempty"`
	SaleLineWarnMsg                        *String    `xmlrpc:"sale_line_warn_msg,omptempty"`
	SaleNumInvoiced                        *Float     `xmlrpc:"sale_num_invoiced,omptempty"`
	SaleOk                                 *Bool      `xmlrpc:"sale_ok,omptempty"`
	SalesCount                             *Float     `xmlrpc:"sales_count,omptempty"`
	SalesGap                               *Float     `xmlrpc:"sales_gap,omptempty"`
	SellerIds                              *Relation  `xmlrpc:"seller_ids,omptempty"`
	Sequence                               *Int       `xmlrpc:"sequence,omptempty"`
	ServiceToPurchase                      *Bool      `xmlrpc:"service_to_purchase,omptempty"`
	ServiceType                            *Selection `xmlrpc:"service_type,omptempty"`
	ShowForecastedQtyStatusButton          *Bool      `xmlrpc:"show_forecasted_qty_status_button,omptempty"`
	ShowOnHandQtyStatusButton              *Bool      `xmlrpc:"show_on_hand_qty_status_button,omptempty"`
	StandardPrice                          *Float     `xmlrpc:"standard_price,omptempty"`
	StockMoveIds                           *Relation  `xmlrpc:"stock_move_ids,omptempty"`
	StockQuantIds                          *Relation  `xmlrpc:"stock_quant_ids,omptempty"`
	StockValuationLayerIds                 *Relation  `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	StorageCategoryCapacityIds             *Relation  `xmlrpc:"storage_category_capacity_ids,omptempty"`
	SupplierTaxesId                        *Relation  `xmlrpc:"supplier_taxes_id,omptempty"`
	TaxString                              *String    `xmlrpc:"tax_string,omptempty"`
	TaxesId                                *Relation  `xmlrpc:"taxes_id,omptempty"`
	ToWeight                               *Bool      `xmlrpc:"to_weight,omptempty"`
	TotalCost                              *Float     `xmlrpc:"total_cost,omptempty"`
	TotalMargin                            *Float     `xmlrpc:"total_margin,omptempty"`
	TotalMarginRate                        *Float     `xmlrpc:"total_margin_rate,omptempty"`
	TotalValue                             *Float     `xmlrpc:"total_value,omptempty"`
	Tracking                               *Selection `xmlrpc:"tracking,omptempty"`
	Turnover                               *Float     `xmlrpc:"turnover,omptempty"`
	Type                                   *Selection `xmlrpc:"type,omptempty"`
	UomId                                  *Many2One  `xmlrpc:"uom_id,omptempty"`
	UomName                                *String    `xmlrpc:"uom_name,omptempty"`
	UomPoId                                *Many2One  `xmlrpc:"uom_po_id,omptempty"`
	ValidEan                               *Bool      `xmlrpc:"valid_ean,omptempty"`
	ValidProductTemplateAttributeLineIds   *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omptempty"`
	Valuation                              *Selection `xmlrpc:"valuation,omptempty"`
	ValueSvl                               *Float     `xmlrpc:"value_svl,omptempty"`
	VariantSellerIds                       *Relation  `xmlrpc:"variant_seller_ids,omptempty"`
	VirtualAvailable                       *Float     `xmlrpc:"virtual_available,omptempty"`
	VisibleExpensePolicy                   *Bool      `xmlrpc:"visible_expense_policy,omptempty"`
	VisibleQtyConfigurator                 *Bool      `xmlrpc:"visible_qty_configurator,omptempty"`
	Volume                                 *Float     `xmlrpc:"volume,omptempty"`
	VolumeUomName                          *String    `xmlrpc:"volume_uom_name,omptempty"`
	WarehouseId                            *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteMessageIds                      *Relation  `xmlrpc:"website_message_ids,omptempty"`
	Weight                                 *Float     `xmlrpc:"weight,omptempty"`
	WeightUomName                          *String    `xmlrpc:"weight_uom_name,omptempty"`
	WriteDate                              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                               *Many2One  `xmlrpc:"write_uid,omptempty"`
	XStudioArtists                         *String    `xmlrpc:"x_studio_artists,omptempty"`
	XStudioArtists3                        *Relation  `xmlrpc:"x_studio_artists3,omptempty"`
	XStudioArtists1                        *String    `xmlrpc:"x_studio_artists_1,omptempty"`
	XStudioCharFieldC6Cg6                  *String    `xmlrpc:"x_studio_char_field_C6Cg6,omptempty"`
	XStudioCharFieldGG7SU                  *String    `xmlrpc:"x_studio_char_field_GG7sU,omptempty"`
	XStudioCharFieldXgKQ0                  *String    `xmlrpc:"x_studio_char_field_XgKQ0,omptempty"`
	XStudioCharFieldYwxuy                  *String    `xmlrpc:"x_studio_char_field_Ywxuy,omptempty"`
	XStudioCharFieldJt0OE                  *String    `xmlrpc:"x_studio_char_field_jt0OE,omptempty"`
	XStudioCharFieldML2Po                  *String    `xmlrpc:"x_studio_char_field_mL2Po,omptempty"`
	XStudioCharFieldVa7IB                  *String    `xmlrpc:"x_studio_char_field_va7iB,omptempty"`
	XStudioCharFieldXZFU2                  *String    `xmlrpc:"x_studio_char_field_xZFU2,omptempty"`
	XStudioCollection                      *Relation  `xmlrpc:"x_studio_collection,omptempty"`
	XStudioCollectionId                    *String    `xmlrpc:"x_studio_collection_id,omptempty"`
	XStudioCountry                         *Relation  `xmlrpc:"x_studio_country,omptempty"`
	XStudioCountryRef                      *Relation  `xmlrpc:"x_studio_country_ref,omptempty"`
	XStudioFormatDescription               *String    `xmlrpc:"x_studio_format_description,omptempty"`
	XStudioGenre                           *Relation  `xmlrpc:"x_studio_genre,omptempty"`
	XStudioInstanceId                      *Int       `xmlrpc:"x_studio_instance_id,omptempty"`
	XStudioIntegerFieldYOcml               *Int       `xmlrpc:"x_studio_integer_field_YOcml,omptempty"`
	XStudioLabels                          *Relation  `xmlrpc:"x_studio_labels,omptempty"`
	XStudioMany2ManyFieldAF3BN             *Relation  `xmlrpc:"x_studio_many2many_field_AF3BN,omptempty"`
	XStudioMany2ManyFieldDkKij             *Relation  `xmlrpc:"x_studio_many2many_field_DkKij,omptempty"`
	XStudioMany2ManyFieldGj4OK             *Relation  `xmlrpc:"x_studio_many2many_field_Gj4OK,omptempty"`
	XStudioMany2ManyFieldLdYkA             *Relation  `xmlrpc:"x_studio_many2many_field_LdYkA,omptempty"`
	XStudioMany2ManyFieldMC7Xf             *Relation  `xmlrpc:"x_studio_many2many_field_MC7Xf,omptempty"`
	XStudioMany2ManyFieldPbGXA             *Relation  `xmlrpc:"x_studio_many2many_field_PbGXA,omptempty"`
	XStudioMany2ManyFieldT0UBA             *Relation  `xmlrpc:"x_studio_many2many_field_T0UBA,omptempty"`
	XStudioMany2ManyFieldFroCV             *Relation  `xmlrpc:"x_studio_many2many_field_froCV,omptempty"`
	XStudioMany2ManyFieldH9EBT             *Relation  `xmlrpc:"x_studio_many2many_field_h9EBT,omptempty"`
	XStudioMany2ManyFieldKoBkM             *Relation  `xmlrpc:"x_studio_many2many_field_koBkM,omptempty"`
	XStudioMany2ManyFieldVlIm5             *Relation  `xmlrpc:"x_studio_many2many_field_vlIm5,omptempty"`
	XStudioMediaCondition                  *Selection `xmlrpc:"x_studio_media_condition,omptempty"`
	XStudioMultilineNotes                  *String    `xmlrpc:"x_studio_multiline_notes,omptempty"`
	XStudioNotes                           *String    `xmlrpc:"x_studio_notes,omptempty"`
	XStudioNotes1                          *String    `xmlrpc:"x_studio_notes_1,omptempty"`
	XStudioOne2ManyFieldMkcRR              *Relation  `xmlrpc:"x_studio_one2many_field_MkcRR,omptempty"`
	XStudioProductFormat                   *Selection `xmlrpc:"x_studio_product_format,omptempty"`
	XStudioRelease                         *Relation  `xmlrpc:"x_studio_release,omptempty"`
	XStudioReleaseId                       *String    `xmlrpc:"x_studio_release_id,omptempty"`
	XStudioReleasePage                     *String    `xmlrpc:"x_studio_release_page,omptempty"`
	XStudioSelectionField8BqT6             *Selection `xmlrpc:"x_studio_selection_field_8BqT6,omptempty"`
	XStudioSelectionFieldXvIp6             *Selection `xmlrpc:"x_studio_selection_field_XvIp6,omptempty"`
	XStudioSelectionFieldNvsHl             *Selection `xmlrpc:"x_studio_selection_field_nvsHl,omptempty"`
	XStudioSleeveCondition                 *Selection `xmlrpc:"x_studio_sleeve_condition,omptempty"`
	XStudioStyles                          *Relation  `xmlrpc:"x_studio_styles,omptempty"`
	XStudioTextFieldHpNtT                  *String    `xmlrpc:"x_studio_text_field_HpNtT,omptempty"`
	XStudioTextFieldCt2Dp                  *String    `xmlrpc:"x_studio_text_field_ct2dp,omptempty"`
	XStudioYearOfTheRelease                *String    `xmlrpc:"x_studio_year_of_the_release,omptempty"`
	XStudioYearOfTheReleaseOfTheMaster     *String    `xmlrpc:"x_studio_year_of_the_release_of_the_master,omptempty"`
	XStudioYearRelease                     *String    `xmlrpc:"x_studio_year_release,omptempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	ids, err := c.CreateProductProducts([]*ProductProduct{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProducts(pps []*ProductProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductProductModel, vv)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.product not found", id)
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.product was not found with criteria %v", criteria)
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.product was not found with criteria %v and options %v", criteria, options)
}
