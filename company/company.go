package company

import "invfinsdk/client"

const (
	allCompaniesPath       string = "lista-empresas"
	singleCompanyFullPath  string = "empresa-completa"
	singleCompanyBasicPath string = "empresa-basica"

	INCOME_STATEMENT = "estado-resultados"
BALANCE_SHEET = "balance-general"
CASHFLOW_STATEMENT = "estado-flujo-efectivo"
)

type Company struct {
	Client client.Client
}

type AllCompaniesResponse struct {
	CompanyResponse[string]
}

func (comp *Company) GetAllCompanies() interface{}{
	return comp.Client.PerformRequest(allCompaniesPath, {"":""}, AllCompaniesResponse)
}

func (comp *Company) GetSimpleCompanyInformation(params map[string]string) interface{}{
	return comp.Client.PerformRequest(singleCompanyBasicPath, params, AllCompaniesResponse)
}

func (comp *Company) GetCompleteCompanyInformation(params map[string]string) interface{}{
	return comp.Client.PerformRequest(singleCompanyFullPath, params, AllCompaniesResponse)
}
