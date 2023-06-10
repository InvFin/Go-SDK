package api

import (
	"fmt"

	"github.com/InvFin/Go-SDK/client"
	"github.com/InvFin/Go-SDK/company"
	"github.com/InvFin/Go-SDK/exchange"
	"github.com/InvFin/Go-SDK/industry"
	"github.com/InvFin/Go-SDK/sector"
	"github.com/InvFin/Go-SDK/superinvestor"
	"github.com/InvFin/Go-SDK/term"
)

type API struct {
	Company       *company.Company
	Term          *term.Term
	Exchange      *exchange.Exchange
	Sector        *sector.Sector
	Industry      *industry.Industry
	SuperInvestor *superinvestor.SuperInvestor
}

func NewAPI(APIKey string) (*API, error) {
	if APIKey == "" {
		return nil, fmt.Errorf("You need to insert an API key")
	}
	api := &API{}
	client := client.Client{APIKey: APIKey}
	api.Company = &company.Company{Client: client}
	api.Term = &term.Term{Client: client}
	api.Exchange = &exchange.Exchange{Client: client}
	api.Sector = &sector.Sector{Client: client}
	api.Industry = &industry.Industry{Client: client}
	api.SuperInvestor = &superinvestor.SuperInvestor{Client: client}
	return api, nil
}
