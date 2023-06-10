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
}
type AllTermsResponse struct {
	Terms []TermResponse
}

func (term *Term) GetAllTerms() interface{} {
	return term.Client.PerformRequest(allTermsPath, map[string]string{"": ""}, AllTermsResponse{})
}

func (term *Term) GetTerm(params map[string]string) interface{} {
	return term.Client.PerformRequest(singleTermPath, params, TermResponse{})
}
