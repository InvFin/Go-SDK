package api

import (
	"fmt"
	"invfinsdk/client"
	"invfinsdk/company"
	"invfinsdk/exchange"
	"invfinsdk/industry"
	"invfinsdk/sector"
	"invfinsdk/superinvestor"
	"invfinsdk/term"
)

type API struct {
	Company       *company.Company
	Term          *term.Term
	Exchange      *exchange.Exchange
	Sector        *sector.Sector
	Industry      *industry.Industry
	SuperInvestor *superinvestor.SuperInvestor
}

func NewAPI(APIKey string) (API, err) {
	if APIKey == "" {
		fmt.Sprintf("You need to insert an API key")
	}
	api := API{}
	client := client.Client{APIKey: APIKey}
	api.Company = &company.Company{Client: client}
	api.Term = &term.Term{Client: client}
	api.Exchange = &exchange.Exchange{Client: client}
	api.Sector = &sector.Sector{Client: client}
	api.Industry = &industry.Industry{Client: client}
	api.SuperInvestor = &superinvestor.SuperInvestor{Client: client}
	return api, nil
}
