package term

import "github.com/InvFin/Go-SDK/client"

const (
	allTermsPath   = "lista-terminos"
	singleTermPath = "termino"
)

type Term struct {
	Client client.Client
}

type TermResponse struct {
	Id    int
	Title string
	Error client.ErrorResponse
}
type AllTermsResponse struct {
	Terms []TermResponse
	Error client.ErrorResponse
}

func (term *Term) GetAllTerms(params map[string]string) interface{} {
	return term.Client.PerformRequest(allTermsPath, params, AllTermsResponse{})
}

func (term *Term) GetTerm(params map[string]string) interface{} {
	return term.Client.PerformRequest(singleTermPath, params, TermResponse{})
}
