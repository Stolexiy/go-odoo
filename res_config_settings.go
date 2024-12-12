package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultCreditLimit                   *Float     `xmlrpc:"account_default_credit_limit,omitempty"`
	AccountDefaultPosReceivableAccountId        *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omitempty"`
	AccountDiscountExpenseAllocationId          *Many2One  `xmlrpc:"account_discount_expense_allocation_id,omitempty"`
	AccountDiscountIncomeAllocationId           *Many2One  `xmlrpc:"account_discount_income_allocation_id,omitempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omitempty"`
	AccountFolderId                             *Many2One  `xmlrpc:"account_folder_id,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	AccountPriceInclude                         *Selection `xmlrpc:"account_price_include,omitempty"`
	AccountReportsShowPerCompanySetting         *Bool      `xmlrpc:"account_reports_show_per_company_setting,omitempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omitempty"`
	AccountTaxPeriodicity                       *Selection `xmlrpc:"account_tax_periodicity,omitempty"`
	AccountTaxPeriodicityJournalId              *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omitempty"`
	AccountTaxPeriodicityReminderDay            *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omitempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omitempty"`
	ActiveUserCount                             *Int       `xmlrpc:"active_user_count,omitempty"`
	AliasDomainId                               *Many2One  `xmlrpc:"alias_domain_id,omitempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omitempty"`
	AuthSignupResetPassword                     *Bool      `xmlrpc:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId                    *Many2One  `xmlrpc:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                         *Selection `xmlrpc:"auth_signup_uninvited,omitempty"`
	AutomaticInvoice                            *Bool      `xmlrpc:"automatic_invoice,omitempty"`
	AutopostBills                               *Bool      `xmlrpc:"autopost_bills,omitempty"`
	BaPassword                                  *String    `xmlrpc:"ba_password,omitempty"`
	BaUsername                                  *String    `xmlrpc:"ba_username,omitempty"`
	BarcodeMaxTimeBetweenKeysInMs               *Int       `xmlrpc:"barcode_max_time_between_keys_in_ms,omitempty"`
	BarcodeNomenclatureId                       *Many2One  `xmlrpc:"barcode_nomenclature_id,omitempty"`
	BarcodeRfidBatchTime                        *Int       `xmlrpc:"barcode_rfid_batch_time,omitempty"`
	BarcodeSeparator                            *String    `xmlrpc:"barcode_separator,omitempty"`
	BarcodeSeparatorRegex                       *String    `xmlrpc:"barcode_separator_regex,omitempty"`
	BarcodelookupApiKey                         *String    `xmlrpc:"barcodelookup_api_key,omitempty"`
	CdnActivated                                *Bool      `xmlrpc:"cdn_activated,omitempty"`
	CdnFilters                                  *String    `xmlrpc:"cdn_filters,omitempty"`
	CdnUrl                                      *String    `xmlrpc:"cdn_url,omitempty"`
	ChartTemplate                               *Selection `xmlrpc:"chart_template,omitempty"`
	CheckAccountAuditTrail                      *Bool      `xmlrpc:"check_account_audit_trail,omitempty"`
	CompanyCount                                *Int       `xmlrpc:"company_count,omitempty"`
	CompanyCountryCode                          *String    `xmlrpc:"company_country_code,omitempty"`
	CompanyCountryId                            *Many2One  `xmlrpc:"company_country_id,omitempty"`
	CompanyCurrencyId                           *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyExpenseAllowedPaymentMethodLineIds   *Relation  `xmlrpc:"company_expense_allowed_payment_method_line_ids,omitempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyInformations                         *String    `xmlrpc:"company_informations,omitempty"`
	CompanyName                                 *String    `xmlrpc:"company_name,omitempty"`
	CompanySoTemplateId                         *Many2One  `xmlrpc:"company_so_template_id,omitempty"`
	ContractExpirationNoticePeriod              *Int       `xmlrpc:"contract_expiration_notice_period,omitempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmAutoAssignmentAction                     *Selection `xmlrpc:"crm_auto_assignment_action,omitempty"`
	CrmAutoAssignmentIntervalNumber             *Int       `xmlrpc:"crm_auto_assignment_interval_number,omitempty"`
	CrmAutoAssignmentIntervalType               *Selection `xmlrpc:"crm_auto_assignment_interval_type,omitempty"`
	CrmAutoAssignmentRunDatetime                *Time      `xmlrpc:"crm_auto_assignment_run_datetime,omitempty"`
	CrmUseAutoAssignment                        *Bool      `xmlrpc:"crm_use_auto_assignment,omitempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyIntervalUnit                        *Selection `xmlrpc:"currency_interval_unit,omitempty"`
	CurrencyNextExecutionDate                   *Time      `xmlrpc:"currency_next_execution_date,omitempty"`
	CurrencyProvider                            *Selection `xmlrpc:"currency_provider,omitempty"`
	DaysToPurchase                              *Float     `xmlrpc:"days_to_purchase,omitempty"`
	DefaultInvoicePolicy                        *Selection `xmlrpc:"default_invoice_policy,omitempty"`
	DefaultPickingPolicy                        *Selection `xmlrpc:"default_picking_policy,omitempty"`
	DefaultPurchaseMethod                       *Selection `xmlrpc:"default_purchase_method,omitempty"`
	DeferredExpenseAccountId                    *Many2One  `xmlrpc:"deferred_expense_account_id,omitempty"`
	DeferredExpenseAmountComputationMethod      *Selection `xmlrpc:"deferred_expense_amount_computation_method,omitempty"`
	DeferredExpenseJournalId                    *Many2One  `xmlrpc:"deferred_expense_journal_id,omitempty"`
	DeferredRevenueAccountId                    *Many2One  `xmlrpc:"deferred_revenue_account_id,omitempty"`
	DeferredRevenueAmountComputationMethod      *Selection `xmlrpc:"deferred_revenue_amount_computation_method,omitempty"`
	DeferredRevenueJournalId                    *Many2One  `xmlrpc:"deferred_revenue_journal_id,omitempty"`
	DeletionDelay                               *Int       `xmlrpc:"deletion_delay,omitempty"`
	DigestEmails                                *Bool      `xmlrpc:"digest_emails,omitempty"`
	DigestId                                    *Many2One  `xmlrpc:"digest_id,omitempty"`
	DisableRedirectFirebaseDynamicLink          *Bool      `xmlrpc:"disable_redirect_firebase_dynamic_link,omitempty"`
	DisplayInvoiceAmountTotalWords              *Bool      `xmlrpc:"display_invoice_amount_total_words,omitempty"`
	DisplayInvoiceTaxCompanyCurrency            *Bool      `xmlrpc:"display_invoice_tax_company_currency,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	DocumentSpreadsheetFolderId                 *Many2One  `xmlrpc:"document_spreadsheet_folder_id,omitempty"`
	DocumentsAccountSettings                    *Bool      `xmlrpc:"documents_account_settings,omitempty"`
	DocumentsHrContractsTags                    *Relation  `xmlrpc:"documents_hr_contracts_tags,omitempty"`
	DocumentsHrFolder                           *Many2One  `xmlrpc:"documents_hr_folder,omitempty"`
	DocumentsHrSettings                         *Bool      `xmlrpc:"documents_hr_settings,omitempty"`
	DocumentsProductSettings                    *Bool      `xmlrpc:"documents_product_settings,omitempty"`
	Domain                                      *String    `xmlrpc:"domain,omitempty"`
	EmailPrimaryColor                           *String    `xmlrpc:"email_primary_color,omitempty"`
	EmailSecondaryColor                         *String    `xmlrpc:"email_secondary_color,omitempty"`
	Enable                                      *Bool      `xmlrpc:"enable,omitempty"`
	EnableOcn                                   *Bool      `xmlrpc:"enable_ocn,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExpenseJournalId                            *Many2One  `xmlrpc:"expense_journal_id,omitempty"`
	ExpenseOutstandingAccountId                 *Many2One  `xmlrpc:"expense_outstanding_account_id,omitempty"`
	ExternalEmailServerDefault                  *Bool      `xmlrpc:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omitempty"`
	ExtractBankStatementDigitalizationMode      *Selection `xmlrpc:"extract_bank_statement_digitalization_mode,omitempty"`
	ExtractInInvoiceDigitalizationMode          *Selection `xmlrpc:"extract_in_invoice_digitalization_mode,omitempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection `xmlrpc:"extract_out_invoice_digitalization_mode,omitempty"`
	ExtractSingleLinePerTax                     *Bool      `xmlrpc:"extract_single_line_per_tax,omitempty"`
	FailCounter                                 *Int       `xmlrpc:"fail_counter,omitempty"`
	Favicon                                     *String    `xmlrpc:"favicon,omitempty"`
	FirstProviderLabel                          *String    `xmlrpc:"first_provider_label,omitempty"`
	FiscalyearLastDay                           *Int       `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth                         *Selection `xmlrpc:"fiscalyear_last_month,omitempty"`
	GenerateDeferredExpenseEntriesMethod        *Selection `xmlrpc:"generate_deferred_expense_entries_method,omitempty"`
	GenerateDeferredRevenueEntriesMethod        *Selection `xmlrpc:"generate_deferred_revenue_entries_method,omitempty"`
	GoogleAnalyticsKey                          *String    `xmlrpc:"google_analytics_key,omitempty"`
	GoogleGmailClientIdentifier                 *String    `xmlrpc:"google_gmail_client_identifier,omitempty"`
	GoogleGmailClientSecret                     *String    `xmlrpc:"google_gmail_client_secret,omitempty"`
	GoogleSearchConsole                         *String    `xmlrpc:"google_search_console,omitempty"`
	GoogleTranslateApiKey                       *String    `xmlrpc:"google_translate_api_key,omitempty"`
	GroupAnalyticAccounting                     *Bool      `xmlrpc:"group_analytic_accounting,omitempty"`
	GroupAutoDoneSetting                        *Bool      `xmlrpc:"group_auto_done_setting,omitempty"`
	GroupBarcodeCountEntireLocation             *Bool      `xmlrpc:"group_barcode_count_entire_location,omitempty"`
	GroupBarcodeShowQuantityCount               *Bool      `xmlrpc:"group_barcode_show_quantity_count,omitempty"`
	GroupCashRounding                           *Bool      `xmlrpc:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                      *Bool      `xmlrpc:"group_discount_per_so_line,omitempty"`
	GroupFiscalYear                             *Bool      `xmlrpc:"group_fiscal_year,omitempty"`
	GroupLotOnDeliverySlip                      *Bool      `xmlrpc:"group_lot_on_delivery_slip,omitempty"`
	GroupLotOnInvoice                           *Bool      `xmlrpc:"group_lot_on_invoice,omitempty"`
	GroupMultiCurrency                          *Bool      `xmlrpc:"group_multi_currency,omitempty"`
	GroupMultiWebsite                           *Bool      `xmlrpc:"group_multi_website,omitempty"`
	GroupProductPricelist                       *Bool      `xmlrpc:"group_product_pricelist,omitempty"`
	GroupProductVariant                         *Bool      `xmlrpc:"group_product_variant,omitempty"`
	GroupProformaSales                          *Bool      `xmlrpc:"group_proforma_sales,omitempty"`
	GroupProjectMilestone                       *Bool      `xmlrpc:"group_project_milestone,omitempty"`
	GroupProjectRating                          *Bool      `xmlrpc:"group_project_rating,omitempty"`
	GroupProjectRecurringTasks                  *Bool      `xmlrpc:"group_project_recurring_tasks,omitempty"`
	GroupProjectStages                          *Bool      `xmlrpc:"group_project_stages,omitempty"`
	GroupProjectTaskDependencies                *Bool      `xmlrpc:"group_project_task_dependencies,omitempty"`
	GroupSaleDeliveryAddress                    *Bool      `xmlrpc:"group_sale_delivery_address,omitempty"`
	GroupSaleOrderTemplate                      *Bool      `xmlrpc:"group_sale_order_template,omitempty"`
	GroupSendReminder                           *Bool      `xmlrpc:"group_send_reminder,omitempty"`
	GroupShowPurchaseReceipts                   *Bool      `xmlrpc:"group_show_purchase_receipts,omitempty"`
	GroupShowSaleReceipts                       *Bool      `xmlrpc:"group_show_sale_receipts,omitempty"`
	GroupStockAccountingAutomatic               *Bool      `xmlrpc:"group_stock_accounting_automatic,omitempty"`
	GroupStockAdvLocation                       *Bool      `xmlrpc:"group_stock_adv_location,omitempty"`
	GroupStockLotPrintGs1                       *Bool      `xmlrpc:"group_stock_lot_print_gs1,omitempty"`
	GroupStockMultiLocations                    *Bool      `xmlrpc:"group_stock_multi_locations,omitempty"`
	GroupStockPackaging                         *Bool      `xmlrpc:"group_stock_packaging,omitempty"`
	GroupStockProductionLot                     *Bool      `xmlrpc:"group_stock_production_lot,omitempty"`
	GroupStockReceptionReport                   *Bool      `xmlrpc:"group_stock_reception_report,omitempty"`
	GroupStockSignDelivery                      *Bool      `xmlrpc:"group_stock_sign_delivery,omitempty"`
	GroupStockTrackingLot                       *Bool      `xmlrpc:"group_stock_tracking_lot,omitempty"`
	GroupStockTrackingOwner                     *Bool      `xmlrpc:"group_stock_tracking_owner,omitempty"`
	GroupUom                                    *Bool      `xmlrpc:"group_uom,omitempty"`
	GroupUseLead                                *Bool      `xmlrpc:"group_use_lead,omitempty"`
	GroupUseRecurringRevenues                   *Bool      `xmlrpc:"group_use_recurring_revenues,omitempty"`
	GroupWarningAccount                         *Bool      `xmlrpc:"group_warning_account,omitempty"`
	GroupWarningPurchase                        *Bool      `xmlrpc:"group_warning_purchase,omitempty"`
	GroupWarningSale                            *Bool      `xmlrpc:"group_warning_sale,omitempty"`
	GroupWarningStock                           *Bool      `xmlrpc:"group_warning_stock,omitempty"`
	HasAccountingEntries                        *Bool      `xmlrpc:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                          *Bool      `xmlrpc:"has_chart_of_accounts,omitempty"`
	HasCountryRegions                           *Bool      `xmlrpc:"has_country_regions,omitempty"`
	HasDefaultShareImage                        *Bool      `xmlrpc:"has_default_share_image,omitempty"`
	HasGoogleAnalytics                          *Bool      `xmlrpc:"has_google_analytics,omitempty"`
	HasGoogleSearchConsole                      *Bool      `xmlrpc:"has_google_search_console,omitempty"`
	HasPlausibleSharedKey                       *Bool      `xmlrpc:"has_plausible_shared_key,omitempty"`
	HrEmployeeSelfEdit                          *Bool      `xmlrpc:"hr_employee_self_edit,omitempty"`
	HrExpenseAliasPrefix                        *String    `xmlrpc:"hr_expense_alias_prefix,omitempty"`
	HrExpenseUseMailgateway                     *Bool      `xmlrpc:"hr_expense_use_mailgateway,omitempty"`
	HrPresenceControlEmail                      *Bool      `xmlrpc:"hr_presence_control_email,omitempty"`
	HrPresenceControlEmailAmount                *Int       `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIp                         *Bool      `xmlrpc:"hr_presence_control_ip,omitempty"`
	HrPresenceControlIpList                     *String    `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	HrPresenceControlLogin                      *Bool      `xmlrpc:"hr_presence_control_login,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	IntrastatDefaultInvoiceTransactionCodeId    *Many2One  `xmlrpc:"intrastat_default_invoice_transaction_code_id,omitempty"`
	IntrastatDefaultRefundTransactionCodeId     *Many2One  `xmlrpc:"intrastat_default_refund_transaction_code_id,omitempty"`
	IntrastatRegionId                           *Many2One  `xmlrpc:"intrastat_region_id,omitempty"`
	InvoiceMailTemplateId                       *Many2One  `xmlrpc:"invoice_mail_template_id,omitempty"`
	InvoicePolicy                               *Bool      `xmlrpc:"invoice_policy,omitempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omitempty"`
	InvoicedTimesheet                           *Selection `xmlrpc:"invoiced_timesheet,omitempty"`
	InvoicingSwitchThreshold                    *Time      `xmlrpc:"invoicing_switch_threshold,omitempty"`
	IsAccountPeppolEligible                     *Bool      `xmlrpc:"is_account_peppol_eligible,omitempty"`
	IsEncodeUomDays                             *Bool      `xmlrpc:"is_encode_uom_days,omitempty"`
	IsInstalledSale                             *Bool      `xmlrpc:"is_installed_sale,omitempty"`
	IsKioskMode                                 *Bool      `xmlrpc:"is_kiosk_mode,omitempty"`
	IsMembershipMulti                           *Bool      `xmlrpc:"is_membership_multi,omitempty"`
	IsRootCompany                               *Bool      `xmlrpc:"is_root_company,omitempty"`
	IsStripeSupportedCountry                    *Bool      `xmlrpc:"is_stripe_supported_country,omitempty"`
	Iso20022InitiatingPartyName                 *String    `xmlrpc:"iso20022_initiating_party_name,omitempty"`
	Iso20022OrgidId                             *String    `xmlrpc:"iso20022_orgid_id,omitempty"`
	Iso20022OrgidIssr                           *String    `xmlrpc:"iso20022_orgid_issr,omitempty"`
	L10NFrIntrastatEnvelopeId                   *String    `xmlrpc:"l10n_fr_intrastat_envelope_id,omitempty"`
	L10NFrRoundingDifferenceLossAccountId       *Many2One  `xmlrpc:"l10n_fr_rounding_difference_loss_account_id,omitempty"`
	L10NFrRoundingDifferenceProfitAccountId     *Many2One  `xmlrpc:"l10n_fr_rounding_difference_profit_account_id,omitempty"`
	LanguageCount                               *Int       `xmlrpc:"language_count,omitempty"`
	LanguageIds                                 *Relation  `xmlrpc:"language_ids,omitempty"`
	LeadEnrichAuto                              *Selection `xmlrpc:"lead_enrich_auto,omitempty"`
	LeadMiningInPipeline                        *Bool      `xmlrpc:"lead_mining_in_pipeline,omitempty"`
	LockConfirmedPo                             *Bool      `xmlrpc:"lock_confirmed_po,omitempty"`
	MapBoxToken                                 *String    `xmlrpc:"map_box_token,omitempty"`
	ModuleAccount3WayMatch                      *Bool      `xmlrpc:"module_account_3way_match,omitempty"`
	ModuleAccountAccountant                     *Bool      `xmlrpc:"module_account_accountant,omitempty"`
	ModuleAccountBankStatementExtract           *Bool      `xmlrpc:"module_account_bank_statement_extract,omitempty"`
	ModuleAccountBankStatementImportCamt        *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv         *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx         *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif         *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment                   *Bool      `xmlrpc:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                         *Bool      `xmlrpc:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting                  *Bool      `xmlrpc:"module_account_check_printing,omitempty"`
	ModuleAccountExtract                        *Bool      `xmlrpc:"module_account_extract,omitempty"`
	ModuleAccountInterCompanyRules              *Bool      `xmlrpc:"module_account_inter_company_rules,omitempty"`
	ModuleAccountIntrastat                      *Bool      `xmlrpc:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract                 *Bool      `xmlrpc:"module_account_invoice_extract,omitempty"`
	ModuleAccountIso20022                       *Bool      `xmlrpc:"module_account_iso20022,omitempty"`
	ModuleAccountPayment                        *Bool      `xmlrpc:"module_account_payment,omitempty"`
	ModuleAccountPeppol                         *Bool      `xmlrpc:"module_account_peppol,omitempty"`
	ModuleAccountReports                        *Bool      `xmlrpc:"module_account_reports,omitempty"`
	ModuleAccountSepaDirectDebit                *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty"`
	ModuleAuthLdap                              *Bool      `xmlrpc:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                             *Bool      `xmlrpc:"module_auth_oauth,omitempty"`
	ModuleBaseGeolocalize                       *Bool      `xmlrpc:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                            *Bool      `xmlrpc:"module_base_import,omitempty"`
	ModuleCrmIapEnrich                          *Bool      `xmlrpc:"module_crm_iap_enrich,omitempty"`
	ModuleCrmIapMine                            *Bool      `xmlrpc:"module_crm_iap_mine,omitempty"`
	ModuleCurrencyRateLive                      *Bool      `xmlrpc:"module_currency_rate_live,omitempty"`
	ModuleDelivery                              *Bool      `xmlrpc:"module_delivery,omitempty"`
	ModuleDeliveryBpost                         *Bool      `xmlrpc:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                           *Bool      `xmlrpc:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                      *Bool      `xmlrpc:"module_delivery_easypost,omitempty"`
	ModuleDeliveryFedex                         *Bool      `xmlrpc:"module_delivery_fedex,omitempty"`
	ModuleDeliverySendcloud                     *Bool      `xmlrpc:"module_delivery_sendcloud,omitempty"`
	ModuleDeliveryShiprocket                    *Bool      `xmlrpc:"module_delivery_shiprocket,omitempty"`
	ModuleDeliveryStarshipit                    *Bool      `xmlrpc:"module_delivery_starshipit,omitempty"`
	ModuleDeliveryUps                           *Bool      `xmlrpc:"module_delivery_ups,omitempty"`
	ModuleDeliveryUsps                          *Bool      `xmlrpc:"module_delivery_usps,omitempty"`
	ModuleGoogleCalendar                        *Bool      `xmlrpc:"module_google_calendar,omitempty"`
	ModuleGoogleGmail                           *Bool      `xmlrpc:"module_google_gmail,omitempty"`
	ModuleGoogleRecaptcha                       *Bool      `xmlrpc:"module_google_recaptcha,omitempty"`
	ModuleHrAttendance                          *Bool      `xmlrpc:"module_hr_attendance,omitempty"`
	ModuleHrExpenseExtract                      *Bool      `xmlrpc:"module_hr_expense_extract,omitempty"`
	ModuleHrHomeworking                         *Bool      `xmlrpc:"module_hr_homeworking,omitempty"`
	ModuleHrPayrollExpense                      *Bool      `xmlrpc:"module_hr_payroll_expense,omitempty"`
	ModuleHrPresence                            *Bool      `xmlrpc:"module_hr_presence,omitempty"`
	ModuleHrSkills                              *Bool      `xmlrpc:"module_hr_skills,omitempty"`
	ModuleHrTimesheet                           *Bool      `xmlrpc:"module_hr_timesheet,omitempty"`
	ModuleL10NEuOss                             *Bool      `xmlrpc:"module_l10n_eu_oss,omitempty"`
	ModuleLoyalty                               *Bool      `xmlrpc:"module_loyalty,omitempty"`
	ModuleMailPlugin                            *Bool      `xmlrpc:"module_mail_plugin,omitempty"`
	ModuleMarketingAutomation                   *Bool      `xmlrpc:"module_marketing_automation,omitempty"`
	ModuleMicrosoftCalendar                     *Bool      `xmlrpc:"module_microsoft_calendar,omitempty"`
	ModuleMicrosoftOutlook                      *Bool      `xmlrpc:"module_microsoft_outlook,omitempty"`
	ModulePartnerAutocomplete                   *Bool      `xmlrpc:"module_partner_autocomplete,omitempty"`
	ModulePosAdyen                              *Bool      `xmlrpc:"module_pos_adyen,omitempty"`
	ModulePosMercadoPago                        *Bool      `xmlrpc:"module_pos_mercado_pago,omitempty"`
	ModulePosPaytm                              *Bool      `xmlrpc:"module_pos_paytm,omitempty"`
	ModulePosPreparationDisplay                 *Bool      `xmlrpc:"module_pos_preparation_display,omitempty"`
	ModulePosPricer                             *Bool      `xmlrpc:"module_pos_pricer,omitempty"`
	ModulePosRazorpay                           *Bool      `xmlrpc:"module_pos_razorpay,omitempty"`
	ModulePosSix                                *Bool      `xmlrpc:"module_pos_six,omitempty"`
	ModulePosStripe                             *Bool      `xmlrpc:"module_pos_stripe,omitempty"`
	ModulePosVivaWallet                         *Bool      `xmlrpc:"module_pos_viva_wallet,omitempty"`
	ModuleProductEmailTemplate                  *Bool      `xmlrpc:"module_product_email_template,omitempty"`
	ModuleProductExpiry                         *Bool      `xmlrpc:"module_product_expiry,omitempty"`
	ModuleProductImages                         *Bool      `xmlrpc:"module_product_images,omitempty"`
	ModuleProductMargin                         *Bool      `xmlrpc:"module_product_margin,omitempty"`
	ModuleProjectForecast                       *Bool      `xmlrpc:"module_project_forecast,omitempty"`
	ModuleProjectTimesheetHolidays              *Bool      `xmlrpc:"module_project_timesheet_holidays,omitempty"`
	ModulePurchaseProductMatrix                 *Bool      `xmlrpc:"module_purchase_product_matrix,omitempty"`
	ModulePurchaseRequisition                   *Bool      `xmlrpc:"module_purchase_requisition,omitempty"`
	ModuleQualityControl                        *Bool      `xmlrpc:"module_quality_control,omitempty"`
	ModuleQualityControlWorksheet               *Bool      `xmlrpc:"module_quality_control_worksheet,omitempty"`
	ModuleSaleAmazon                            *Bool      `xmlrpc:"module_sale_amazon,omitempty"`
	ModuleSaleCommission                        *Bool      `xmlrpc:"module_sale_commission,omitempty"`
	ModuleSaleLoyalty                           *Bool      `xmlrpc:"module_sale_loyalty,omitempty"`
	ModuleSaleMargin                            *Bool      `xmlrpc:"module_sale_margin,omitempty"`
	ModuleSalePdfQuoteBuilder                   *Bool      `xmlrpc:"module_sale_pdf_quote_builder,omitempty"`
	ModuleSaleProductMatrix                     *Bool      `xmlrpc:"module_sale_product_matrix,omitempty"`
	ModuleSign                                  *Bool      `xmlrpc:"module_sign,omitempty"`
	ModuleSms                                   *Bool      `xmlrpc:"module_sms,omitempty"`
	ModuleSnailmailAccount                      *Bool      `xmlrpc:"module_snailmail_account,omitempty"`
	ModuleStockBarcode                          *Bool      `xmlrpc:"module_stock_barcode,omitempty"`
	ModuleStockBarcodeBarcodelookup             *Bool      `xmlrpc:"module_stock_barcode_barcodelookup,omitempty"`
	ModuleStockDropshipping                     *Bool      `xmlrpc:"module_stock_dropshipping,omitempty"`
	ModuleStockFleet                            *Bool      `xmlrpc:"module_stock_fleet,omitempty"`
	ModuleStockLandedCosts                      *Bool      `xmlrpc:"module_stock_landed_costs,omitempty"`
	ModuleStockPickingBatch                     *Bool      `xmlrpc:"module_stock_picking_batch,omitempty"`
	ModuleStockSms                              *Bool      `xmlrpc:"module_stock_sms,omitempty"`
	ModuleVoip                                  *Bool      `xmlrpc:"module_voip,omitempty"`
	ModuleWebUnsplash                           *Bool      `xmlrpc:"module_web_unsplash,omitempty"`
	ModuleWebsiteCfTurnstile                    *Bool      `xmlrpc:"module_website_cf_turnstile,omitempty"`
	ModuleWebsiteCrmIapReveal                   *Bool      `xmlrpc:"module_website_crm_iap_reveal,omitempty"`
	ModuleWebsiteLivechat                       *Bool      `xmlrpc:"module_website_livechat,omitempty"`
	PartnerAutocompleteInsufficientCredit       *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omitempty"`
	PayInvoicesOnline                           *Bool      `xmlrpc:"pay_invoices_online,omitempty"`
	PlanningEmployeeUnavailabilities            *Selection `xmlrpc:"planning_employee_unavailabilities,omitempty"`
	PlanningGenerationInterval                  *Int       `xmlrpc:"planning_generation_interval,omitempty"`
	PlanningSelfUnassignDaysBefore              *Int       `xmlrpc:"planning_self_unassign_days_before,omitempty"`
	PlausibleSharedKey                          *String    `xmlrpc:"plausible_shared_key,omitempty"`
	PlausibleSite                               *String    `xmlrpc:"plausible_site,omitempty"`
	PoDoubleValidation                          *Selection `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float     `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                                      *Float     `xmlrpc:"po_lead,omitempty"`
	PoLock                                      *Selection `xmlrpc:"po_lock,omitempty"`
	PoOrderApproval                             *Bool      `xmlrpc:"po_order_approval,omitempty"`
	PointOfSaleTicketPortalUrlDisplayMode       *Selection `xmlrpc:"point_of_sale_ticket_portal_url_display_mode,omitempty"`
	PointOfSaleTicketUniqueCode                 *Bool      `xmlrpc:"point_of_sale_ticket_unique_code,omitempty"`
	PointOfSaleUseTicketQrCode                  *Bool      `xmlrpc:"point_of_sale_use_ticket_qr_code,omitempty"`
	PortalAllowApiKeys                          *Bool      `xmlrpc:"portal_allow_api_keys,omitempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omitempty"`
	PosAdvancedEmployeeIds                      *Relation  `xmlrpc:"pos_advanced_employee_ids,omitempty"`
	PosAllowedPricelistIds                      *Relation  `xmlrpc:"pos_allowed_pricelist_ids,omitempty"`
	PosAmountAuthorizedDiff                     *Float     `xmlrpc:"pos_amount_authorized_diff,omitempty"`
	PosAutoValidateTerminalPayment              *Bool      `xmlrpc:"pos_auto_validate_terminal_payment,omitempty"`
	PosAvailablePricelistIds                    *Relation  `xmlrpc:"pos_available_pricelist_ids,omitempty"`
	PosBasicEmployeeIds                         *Relation  `xmlrpc:"pos_basic_employee_ids,omitempty"`
	PosBasicReceipt                             *Bool      `xmlrpc:"pos_basic_receipt,omitempty"`
	PosCashControl                              *Bool      `xmlrpc:"pos_cash_control,omitempty"`
	PosCashRounding                             *Bool      `xmlrpc:"pos_cash_rounding,omitempty"`
	PosCompanyHasTemplate                       *Bool      `xmlrpc:"pos_company_has_template,omitempty"`
	PosConfigId                                 *Many2One  `xmlrpc:"pos_config_id,omitempty"`
	PosCrmTeamId                                *Many2One  `xmlrpc:"pos_crm_team_id,omitempty"`
	PosCustomerDisplayBgImg                     *String    `xmlrpc:"pos_customer_display_bg_img,omitempty"`
	PosCustomerDisplayBgImgName                 *String    `xmlrpc:"pos_customer_display_bg_img_name,omitempty"`
	PosCustomerDisplayType                      *Selection `xmlrpc:"pos_customer_display_type,omitempty"`
	PosDefaultBillIds                           *Relation  `xmlrpc:"pos_default_bill_ids,omitempty"`
	PosDefaultFiscalPositionId                  *Many2One  `xmlrpc:"pos_default_fiscal_position_id,omitempty"`
	PosDiscountPc                               *Float     `xmlrpc:"pos_discount_pc,omitempty"`
	PosDiscountProductId                        *Many2One  `xmlrpc:"pos_discount_product_id,omitempty"`
	PosDownPaymentProductId                     *Many2One  `xmlrpc:"pos_down_payment_product_id,omitempty"`
	PosEpsonPrinterIp                           *String    `xmlrpc:"pos_epson_printer_ip,omitempty"`
	PosFiscalPositionIds                        *Relation  `xmlrpc:"pos_fiscal_position_ids,omitempty"`
	PosHasActiveSession                         *Bool      `xmlrpc:"pos_has_active_session,omitempty"`
	PosIfaceAvailableCategIds                   *Relation  `xmlrpc:"pos_iface_available_categ_ids,omitempty"`
	PosIfaceBigScrollbars                       *Bool      `xmlrpc:"pos_iface_big_scrollbars,omitempty"`
	PosIfaceCashdrawer                          *Bool      `xmlrpc:"pos_iface_cashdrawer,omitempty"`
	PosIfaceElectronicScale                     *Bool      `xmlrpc:"pos_iface_electronic_scale,omitempty"`
	PosIfacePrintAuto                           *Bool      `xmlrpc:"pos_iface_print_auto,omitempty"`
	PosIfacePrintSkipScreen                     *Bool      `xmlrpc:"pos_iface_print_skip_screen,omitempty"`
	PosIfacePrintViaProxy                       *Bool      `xmlrpc:"pos_iface_print_via_proxy,omitempty"`
	PosIfaceScanViaProxy                        *Bool      `xmlrpc:"pos_iface_scan_via_proxy,omitempty"`
	PosIfaceTaxIncluded                         *Selection `xmlrpc:"pos_iface_tax_included,omitempty"`
	PosIfaceTipproduct                          *Bool      `xmlrpc:"pos_iface_tipproduct,omitempty"`
	PosInvoiceJournalId                         *Many2One  `xmlrpc:"pos_invoice_journal_id,omitempty"`
	PosIsClosingEntryByProduct                  *Bool      `xmlrpc:"pos_is_closing_entry_by_product,omitempty"`
	PosIsHeaderOrFooter                         *Bool      `xmlrpc:"pos_is_header_or_footer,omitempty"`
	PosIsMarginsCostsAccessibleToEveryUser      *Bool      `xmlrpc:"pos_is_margins_costs_accessible_to_every_user,omitempty"`
	PosIsOrderPrinter                           *Bool      `xmlrpc:"pos_is_order_printer,omitempty"`
	PosIsPosbox                                 *Bool      `xmlrpc:"pos_is_posbox,omitempty"`
	PosJournalId                                *Many2One  `xmlrpc:"pos_journal_id,omitempty"`
	PosLimitCategories                          *Bool      `xmlrpc:"pos_limit_categories,omitempty"`
	PosManualDiscount                           *Bool      `xmlrpc:"pos_manual_discount,omitempty"`
	PosModulePosAvatax                          *Bool      `xmlrpc:"pos_module_pos_avatax,omitempty"`
	PosModulePosDiscount                        *Bool      `xmlrpc:"pos_module_pos_discount,omitempty"`
	PosModulePosHr                              *Bool      `xmlrpc:"pos_module_pos_hr,omitempty"`
	PosModulePosRestaurant                      *Bool      `xmlrpc:"pos_module_pos_restaurant,omitempty"`
	PosModulePosSms                             *Bool      `xmlrpc:"pos_module_pos_sms,omitempty"`
	PosModulePosUrbanPiper                      *Bool      `xmlrpc:"pos_module_pos_urban_piper,omitempty"`
	PosNoteIds                                  *Relation  `xmlrpc:"pos_note_ids,omitempty"`
	PosOnlyRoundCashMethod                      *Bool      `xmlrpc:"pos_only_round_cash_method,omitempty"`
	PosOrderEditTracking                        *Bool      `xmlrpc:"pos_order_edit_tracking,omitempty"`
	PosOrderlinesSequenceInCartByCategory       *Bool      `xmlrpc:"pos_orderlines_sequence_in_cart_by_category,omitempty"`
	PosOtherDevices                             *Bool      `xmlrpc:"pos_other_devices,omitempty"`
	PosPaymentMethodIds                         *Relation  `xmlrpc:"pos_payment_method_ids,omitempty"`
	PosPickingPolicy                            *Selection `xmlrpc:"pos_picking_policy,omitempty"`
	PosPickingTypeId                            *Many2One  `xmlrpc:"pos_picking_type_id,omitempty"`
	PosPricelistId                              *Many2One  `xmlrpc:"pos_pricelist_id,omitempty"`
	PosPrinterIds                               *Relation  `xmlrpc:"pos_printer_ids,omitempty"`
	PosProxyIp                                  *String    `xmlrpc:"pos_proxy_ip,omitempty"`
	PosReceiptFooter                            *String    `xmlrpc:"pos_receipt_footer,omitempty"`
	PosReceiptHeader                            *String    `xmlrpc:"pos_receipt_header,omitempty"`
	PosRestrictPriceControl                     *Bool      `xmlrpc:"pos_restrict_price_control,omitempty"`
	PosRoundingMethod                           *Many2One  `xmlrpc:"pos_rounding_method,omitempty"`
	PosRouteId                                  *Many2One  `xmlrpc:"pos_route_id,omitempty"`
	PosSelectableCategIds                       *Relation  `xmlrpc:"pos_selectable_categ_ids,omitempty"`
	PosSequenceId                               *Many2One  `xmlrpc:"pos_sequence_id,omitempty"`
	PosSetMaximumDifference                     *Bool      `xmlrpc:"pos_set_maximum_difference,omitempty"`
	PosShipLater                                *Bool      `xmlrpc:"pos_ship_later,omitempty"`
	PosShowCategoryImages                       *Bool      `xmlrpc:"pos_show_category_images,omitempty"`
	PosShowProductImages                        *Bool      `xmlrpc:"pos_show_product_images,omitempty"`
	PosTaxRegimeSelection                       *Bool      `xmlrpc:"pos_tax_regime_selection,omitempty"`
	PosTipProductId                             *Many2One  `xmlrpc:"pos_tip_product_id,omitempty"`
	PosTrustedConfigIds                         *Relation  `xmlrpc:"pos_trusted_config_ids,omitempty"`
	PosUsePricelist                             *Bool      `xmlrpc:"pos_use_pricelist,omitempty"`
	PosWarehouseId                              *Many2One  `xmlrpc:"pos_warehouse_id,omitempty"`
	PredictBillProduct                          *Bool      `xmlrpc:"predict_bill_product,omitempty"`
	PredictiveLeadScoringFieldLabels            *String    `xmlrpc:"predictive_lead_scoring_field_labels,omitempty"`
	PredictiveLeadScoringFields                 *Relation  `xmlrpc:"predictive_lead_scoring_fields,omitempty"`
	PredictiveLeadScoringFieldsStr              *String    `xmlrpc:"predictive_lead_scoring_fields_str,omitempty"`
	PredictiveLeadScoringStartDate              *Time      `xmlrpc:"predictive_lead_scoring_start_date,omitempty"`
	PredictiveLeadScoringStartDateStr           *String    `xmlrpc:"predictive_lead_scoring_start_date_str,omitempty"`
	PrepaymentPercent                           *Float     `xmlrpc:"prepayment_percent,omitempty"`
	PreviewReady                                *Bool      `xmlrpc:"preview_ready,omitempty"`
	ProductFolderId                             *Many2One  `xmlrpc:"product_folder_id,omitempty"`
	ProductTagIds                               *Relation  `xmlrpc:"product_tag_ids,omitempty"`
	ProductVolumeVolumeInCubicFeet              *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                          *Selection `xmlrpc:"product_weight_in_lbs,omitempty"`
	ProfilingEnabledUntil                       *Time      `xmlrpc:"profiling_enabled_until,omitempty"`
	ProjectTimeModeId                           *Many2One  `xmlrpc:"project_time_mode_id,omitempty"`
	PropertyAccountExpenseCategId               *Many2One  `xmlrpc:"property_account_expense_categ_id,omitempty"`
	PropertyAccountIncomeCategId                *Many2One  `xmlrpc:"property_account_income_categ_id,omitempty"`
	PropertyStockAccountInputCategId            *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId           *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockJournal                        *Many2One  `xmlrpc:"property_stock_journal,omitempty"`
	PropertyStockValuationAccountId             *Many2One  `xmlrpc:"property_stock_valuation_account_id,omitempty"`
	ProvidersState                              *Selection `xmlrpc:"providers_state,omitempty"`
	PurchaseTaxId                               *Many2One  `xmlrpc:"purchase_tax_id,omitempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omitempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omitempty"`
	RecaptchaMinScore                           *Float     `xmlrpc:"recaptcha_min_score,omitempty"`
	RecaptchaPrivateKey                         *String    `xmlrpc:"recaptcha_private_key,omitempty"`
	RecaptchaPublicKey                          *String    `xmlrpc:"recaptcha_public_key,omitempty"`
	ReminderAllow                               *Bool      `xmlrpc:"reminder_allow,omitempty"`
	ReminderDelay                               *Int       `xmlrpc:"reminder_delay,omitempty"`
	ReminderInterval                            *Selection `xmlrpc:"reminder_interval,omitempty"`
	ReminderUserAllow                           *Bool      `xmlrpc:"reminder_user_allow,omitempty"`
	ReminderUserDelay                           *Int       `xmlrpc:"reminder_user_delay,omitempty"`
	ReminderUserInterval                        *Selection `xmlrpc:"reminder_user_interval,omitempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omitempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	RestrictTemplateRendering                   *Bool      `xmlrpc:"restrict_template_rendering,omitempty"`
	SaleTaxId                                   *Many2One  `xmlrpc:"sale_tax_id,omitempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omitempty"`
	SfuServerKey                                *String    `xmlrpc:"sfu_server_key,omitempty"`
	SfuServerUrl                                *String    `xmlrpc:"sfu_server_url,omitempty"`
	SharedUserAccount                           *Bool      `xmlrpc:"shared_user_account,omitempty"`
	ShowBarcodeNomenclature                     *Bool      `xmlrpc:"show_barcode_nomenclature,omitempty"`
	ShowEffect                                  *Bool      `xmlrpc:"show_effect,omitempty"`
	SignInvoice                                 *Bool      `xmlrpc:"sign_invoice,omitempty"`
	SigningUser                                 *Many2One  `xmlrpc:"signing_user,omitempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailCoverReadonly                      *Bool      `xmlrpc:"snailmail_cover_readonly,omitempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omitempty"`
	SocialDefaultImage                          *String    `xmlrpc:"social_default_image,omitempty"`
	StockBarcodeDemoActive                      *Bool      `xmlrpc:"stock_barcode_demo_active,omitempty"`
	StockBarcodeMuteSoundNotifications          *Bool      `xmlrpc:"stock_barcode_mute_sound_notifications,omitempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TenorApiKey                                 *String    `xmlrpc:"tenor_api_key,omitempty"`
	TenorContentFilter                          *Selection `xmlrpc:"tenor_content_filter,omitempty"`
	TenorGifLimit                               *Int       `xmlrpc:"tenor_gif_limit,omitempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omitempty"`
	TimesheetEncodeMethod                       *Selection `xmlrpc:"timesheet_encode_method,omitempty"`
	TimesheetMinDuration                        *Int       `xmlrpc:"timesheet_min_duration,omitempty"`
	TimesheetRounding                           *Int       `xmlrpc:"timesheet_rounding,omitempty"`
	TimesheetShowLeaderboard                    *Bool      `xmlrpc:"timesheet_show_leaderboard,omitempty"`
	TimesheetShowRates                          *Bool      `xmlrpc:"timesheet_show_rates,omitempty"`
	TotalsBelowSections                         *Bool      `xmlrpc:"totals_below_sections,omitempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omitempty"`
	TwilioAccountSid                            *String    `xmlrpc:"twilio_account_sid,omitempty"`
	TwilioAccountToken                          *String    `xmlrpc:"twilio_account_token,omitempty"`
	UnsplashAccessKey                           *String    `xmlrpc:"unsplash_access_key,omitempty"`
	UnsplashAppId                               *String    `xmlrpc:"unsplash_app_id,omitempty"`
	UpdateStockQuantities                       *Selection `xmlrpc:"update_stock_quantities,omitempty"`
	UseAngloSaxon                               *Bool      `xmlrpc:"use_anglo_saxon,omitempty"`
	UseInvoiceTerms                             *Bool      `xmlrpc:"use_invoice_terms,omitempty"`
	UsePoLead                                   *Bool      `xmlrpc:"use_po_lead,omitempty"`
	UseSecurityLead                             *Bool      `xmlrpc:"use_security_lead,omitempty"`
	UseTwilioRtcServers                         *Bool      `xmlrpc:"use_twilio_rtc_servers,omitempty"`
	UserDefaultRights                           *Bool      `xmlrpc:"user_default_rights,omitempty"`
	VatCheckVies                                *Bool      `xmlrpc:"vat_check_vies,omitempty"`
	WebAppName                                  *String    `xmlrpc:"web_app_name,omitempty"`
	WebsiteBlockThirdPartyDomains               *Bool      `xmlrpc:"website_block_third_party_domains,omitempty"`
	WebsiteCompanyId                            *Many2One  `xmlrpc:"website_company_id,omitempty"`
	WebsiteCookiesBar                           *Bool      `xmlrpc:"website_cookies_bar,omitempty"`
	WebsiteDefaultLangCode                      *String    `xmlrpc:"website_default_lang_code,omitempty"`
	WebsiteDefaultLangId                        *Many2One  `xmlrpc:"website_default_lang_id,omitempty"`
	WebsiteDomain                               *String    `xmlrpc:"website_domain,omitempty"`
	WebsiteHomepageUrl                          *String    `xmlrpc:"website_homepage_url,omitempty"`
	WebsiteId                                   *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteLanguageCount                        *Int       `xmlrpc:"website_language_count,omitempty"`
	WebsiteLogo                                 *String    `xmlrpc:"website_logo,omitempty"`
	WebsiteName                                 *String    `xmlrpc:"website_name,omitempty"`
	WorkPermitExpirationNoticePeriod            *Int       `xmlrpc:"work_permit_expiration_notice_period,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	ids, err := c.CreateResConfigSettingss([]*ResConfigSettings{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettingss(rcss []*ResConfigSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResConfigSettingsModel, vv, nil)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs, nil)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResConfigSettingsModel, criteria, options)
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
