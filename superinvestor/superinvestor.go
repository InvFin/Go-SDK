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
}

type SuperInvestorsHistorialResponse struct {
	Historial []SuperInvestorsHistorial
}

type SuperInvestorsMovementsResponse struct {
	Movements []SuperInvestorsMovement
}

func (spi *SuperInvestor) GetAllSuperinvestors() interface{} {
	return spi.Client.PerformRequest(AllSuperInvestorsPath, map[string]string{"": ""}, AllSuperInvestorsResponse{})
}

func (spi *SuperInvestor) GetSuperinvestorHistorial(params map[string]string) interface{} {
	return spi.Client.PerformRequest(SuperInvestorHistorialPath, params, SuperInvestorsHistorialResponse{})
}

func (spi *SuperInvestor) GetSuperinvestorMovements(params map[string]string) interface{} {
	return spi.Client.PerformRequest(SuperInvestorMovementsPath, params, SuperInvestorsMovementsResponse{})
}
