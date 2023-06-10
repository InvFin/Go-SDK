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
}

func (exch *Exchange) GetAllExchanges() interface{} {
	return exch.Client.PerformRequest(AllExchangesPath, map[string]string{"": ""}, AllExchangesResponse{})
}
