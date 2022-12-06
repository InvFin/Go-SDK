package exchange

import "invfinsdk/client"

const AllExchangesPath string = "lista-exchanges"

type Exchange struct {
	Client client.Client
}

type SimpleExchangeResponse struct {
	Id int
	Name string
}

type AllExchangesResponse struct {
	[]SimpleExchangeResponse
}

func (exch *Exchange) GetAllExchanges() interface{}{
	return exch.Client.PerformRequest(AllExchangesPath, {"":""}, AllExchangesResponse)
}
