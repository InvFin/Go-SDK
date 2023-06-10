package superinvestor

import (
	"github.com/InvFin/Go-SDK/client"
)

const (
	AllSuperInvestorsPath      string = "lista-superinversores"
	SuperInvestorHistorialPath string = "lista-historial"
	SuperInvestorMovementsPath string = "lista-movimientos"
)

type SuperInvestor struct {
	Client client.Client
}

type SuperInvestorResponse struct {
	Id   int
	Name string
}

type SuperInvestorsHistorial struct {
	Id   int
	Name string
}

type SuperInvestorsMovement struct {
	Id   int
	Name string
}
type AllSuperInvestorsResponse struct {
	SuperInvestors []SuperInvestorResponse
	Error          client.ErrorResponse
}

type SuperInvestorsHistorialResponse struct {
	Historial []SuperInvestorsHistorial
	Error     client.ErrorResponse
}

type SuperInvestorsMovementsResponse struct {
	Movements []SuperInvestorsMovement
	Error     client.ErrorResponse
}

func (spi *SuperInvestor) GetAllSuperinvestors(params map[string]string) interface{} {
	return spi.Client.PerformRequest(AllSuperInvestorsPath, params, AllSuperInvestorsResponse{})
}

func (spi *SuperInvestor) GetSuperinvestorHistorial(params map[string]string) interface{} {
	return spi.Client.PerformRequest(SuperInvestorHistorialPath, params, SuperInvestorsHistorialResponse{})
}

func (spi *SuperInvestor) GetSuperinvestorMovements(params map[string]string) interface{} {
	return spi.Client.PerformRequest(SuperInvestorMovementsPath, params, SuperInvestorsMovementsResponse{})
}
