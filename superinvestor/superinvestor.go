package superinvestor

import "invfinsdk/client"

const (
	AllSuperInvestorsPath string = "lista-superinversores"
	SuperInvestorHistorialPath string = "lista-historial"
	SuperInvestorMovementsPath string = "lista-movimientos"
)

type SuperInvestor struct {
	Client client.Client
}

type SuperInvestorResponse struct {
	Id int
	Name string
}

type AllSuperInvestorsResponse struct {
	[]SuperInvestorResponse
}


func (spi *SuperInvestor) GetAllSuperinvestors() interface{}{
	return spi.Client.PerformRequest(AllSuperInvestorsPath, {"":""}, AllSuperInvestorsResponse)
}

func (spi *SuperInvestor) GetSuperinvestorHistorial(params map[string]string) interface{}{
	return spi.Client.PerformRequest(SuperInvestorHistorialPath, params, AllCompaniesResponse)
}

func (spi *SuperInvestor) GetSuperinvestorMovements(params map[string]string) interface{}{
	return spi.Client.PerformRequest(SuperInvestorMovementsPath, params, AllCompaniesResponse)
}