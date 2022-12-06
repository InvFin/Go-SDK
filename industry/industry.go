package industry

import "invfinsdk/client"

const AllIndustriesPath string = "lista-industrias"

type Industry struct {
	Client client.Client
}

type SimpleIndustryResponse struct {
	Id int
	Name string
}

type AllIndustriesResponse struct {
	[]SimpleIndustryResponse
}

func (inds *Industry) GetAllIndustries() interface{}{
	return inds.Client.PerformRequest(AllIndustriesPath, {"":""}, AllIndustriesResponse)
}