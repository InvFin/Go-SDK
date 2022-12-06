package term

import "invfinsdk/client"

const (
	allTermsPath   = "lista-terminos"
	singleTermPath = "termino"
)

type Term struct {
	Client client.Client
}

func (term *Term) GetAllTerms() interface{}{
	return term.Client.PerformRequest(allTermsPath, {"":""}, AllCompaniesResponse)
}

func (term *Term) GetTerm(params map[string]string) interface{}{
	return term.Client.PerformRequest(singleTermPath, params, AllCompaniesResponse)
}