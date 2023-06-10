package industry

import "github.com/InvFin/Go-SDK/client"

const AllIndustriesPath string = "lista-industrias"

type Industry struct {
	Client client.Client
}

type SimpleIndustryResponse struct {
	Id   int
	Name string
}

type AllIndustriesResponse struct {
	Industries []SimpleIndustryResponse
}

func (inds *Industry) GetAllIndustries() interface{} {
	return inds.Client.PerformRequest(AllIndustriesPath, map[string]string{"": ""}, AllIndustriesResponse{})
}
