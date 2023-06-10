package sector

import "github.com/InvFin/Go-SDK/client"

const AllSectorsPath string = "lista-sectores"

type Sector struct {
	Client client.Client
}

type SimpleSectorResponse struct {
	Id   int
	Name string
}

type AllSectorssResponse struct {
	Sectors []SimpleSectorResponse
	Error   client.ErrorResponse
}

func (sct *Sector) GetAllSectors() interface{} {
	return sct.Client.PerformRequest(AllSectorsPath, map[string]string{"": ""}, AllSectorssResponse{})
}
