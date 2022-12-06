package sector

import "invfinsdk/client"

const AllSectorsPath string = "lista-sectores"

type Sector struct {
	Client client.Client
}

type SimpleSectorResponse struct {
	Id int
	Name string
}

type AllSectorssResponse struct {
	[]SimpleSectorResponse
}

func (sct *Sector) GetAllSectors() interface{}{
	return sct.Client.PerformRequest(AllSectorsPath, {"":""}, AllSectorssResponse)
}