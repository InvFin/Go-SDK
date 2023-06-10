package exchange

import "github.com/InvFin/Go-SDK/client"

const AllExchangesPath string = "lista-exchanges"

type Exchange struct {
	Client client.Client
}

type SimpleExchangeResponse struct {
	Id   int
	Name string
}

type AllExchangesResponse struct {
	Exchanges []SimpleExchangeResponse
	Error     client.ErrorResponse
}

func (exch *Exchange) GetAllExchanges() interface{} {
	return exch.Client.PerformRequest(AllExchangesPath, nil, AllExchangesResponse{})
}
