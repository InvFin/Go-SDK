package company

import "github.com/InvFin/Go-SDK/client"

const (
	allCompaniesPath       string = "lista-empresas"
	singleCompanyFullPath  string = "empresa-completa"
	singleCompanyBasicPath string = "empresa-basica"

	INCOME_STATEMENT   = "estado-resultados"
	BALANCE_SHEET      = "balance-general"
	CASHFLOW_STATEMENT = "estado-flujo-efectivo"
)

type Company struct {
	Client client.Client
}

type BasicCompanyResponse struct {
	Ticker      string
	Name        string
	Currency    string
	Industry    string
	Sector      string
	Website     string
	State       string
	Country     string
	Ceo         string
	Image       string
	City        string
	Employees   string
	Address     string
	Zip_code    string
	Cik         string
	Exchange    string
	Cusip       string
	Isin        string
	Description string
	IpoDate     string
	Error       client.ErrorResponse
}

type AllCompaniesResponse struct {
	Companies []BasicCompanyResponse
	Error     client.ErrorResponse
}

type IncomeStatement struct {
	Revenue                                      float32
	Cost_of_revenue                              float32
	Gross_profit                                 float32
	Rd_expenses                                  float32
	General_administrative_expenses              float32
	Selling_marketing_expenses                   float32
	Sga_expenses                                 float32
	Other_expenses                               float32
	Operating_expenses                           float32
	Cost_and_expenses                            float32
	Interest_expense                             float32
	Depreciation_amortization                    float32
	Ebitda                                       float32
	Operating_income                             float32
	Net_total_other_income_expenses              float32
	Income_before_tax                            float32
	Income_tax_expenses                          float32
	Net_income                                   float32
	Weighted_average_shares_outstanding          float32
	Weighted_average_diluated_shares_outstanding float32
}

type BalanceSheet struct {
	Cash_and_cash_equivalents                   float32
	Short_term_investments                      float32
	Cash_and_short_term_investments             float32
	Net_receivables                             float32
	Inventory                                   float32
	Other_current_assets                        float32
	Total_current_assets                        float32
	Property_plant_equipment                    float32
	Goodwill                                    float32
	Intangible_assets                           float32
	Goodwill_and_intangible_assets              float32
	Long_term_investments                       float32
	Tax_assets                                  float32
	Other_non_current_assets                    float32
	Total_non_current_assets                    float32
	Other_assets                                float32
	Total_assets                                float32
	Accounts_payable                            float32
	Short_term_debt                             float32
	Tax_payables                                float32
	Deferred_revenue                            float32
	Other_current_liabilities                   float32
	Total_current_liabilities                   float32
	Long_term_debt                              float32
	Deferred_revenue_non_current                float32
	Deferred_tax_liabilities_non_current        float32
	Other_non_current_liabilities               float32
	Total_non_current_liabilities               float32
	Other_liabilities                           float32
	Total_liabilities                           float32
	Common_stocks                               float32
	Preferred_stocks                            float32
	Retained_earnings                           float32
	Accumulated_other_comprehensive_income_loss float32
	Othertotal_stockholders_equity              float32
	Total_stockholders_equity                   float32
	Total_liabilities_and_total_equity          float32
	Total_investments                           float32
	Total_debt                                  float32
	Net_debt                                    float32
}

type CashflowStatement struct {
	Net_income                           float32
	Depreciation_amortization            float32
	Deferred_income_tax                  float32
	Stock_based_compensation             float32
	Change_in_working_capital            float32
	Accounts_receivable                  float32
	Inventory                            float32
	Accounts_payable                     float32
	Other_working_capital                float32
	Other_non_cash_items                 float32
	Operating_activities_cf              float32
	Investments_property_plant_equipment float32
	Acquisitions_net                     float32
	Purchases_investments                float32
	Sales_maturities_investments         float32
	Other_investing_activites            float32
	Investing_activities_cf              float32
	Debt_repayment                       float32
	Common_stock_issued                  float32
	Common_stock_repurchased             float32
	Dividends_paid                       float32
	Other_financing_activities           float32
	Financing_activities_cf              float32
	Effect_forex_exchange                float32
	Net_change_cash                      float32
	Cash_end_period                      float32
	Cash_beginning_period                float32
	Operating_cf                         float32
	Capex                                float32
	Fcf                                  float32
}

type CompleteCompanyResponse struct {
	Basic              BasicCompanyResponse
	IncomeStatements   []IncomeStatement
	BalanceSheets      []BalanceSheet
	CashflowStatements []CashflowStatement
}

func (comp *Company) GetAllCompanies(params map[string]string) interface{} {
	return comp.Client.PerformRequest(allCompaniesPath, params, AllCompaniesResponse{})
}

func (comp *Company) GetBasicCompanyInformation(params map[string]string) interface{} {
	return comp.Client.PerformRequest(singleCompanyBasicPath, params, BasicCompanyResponse{})
}

func (comp *Company) GetCompleteCompanyInformation(params map[string]string) interface{} {
	return comp.Client.PerformRequest(singleCompanyFullPath, params, CompleteCompanyResponse{})
}
