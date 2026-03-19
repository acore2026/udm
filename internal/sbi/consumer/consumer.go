package consumer

import (
	Nnrf_NFDiscovery "github.com/acore2026/openapi/nrf/NFDiscovery"
	Nnrf_NFManagement "github.com/acore2026/openapi/nrf/NFManagement"
	Nudm_SubscriberDataManagement "github.com/acore2026/openapi/udm/SubscriberDataManagement"
	Nudm_UEContextManagement "github.com/acore2026/openapi/udm/UEContextManagement"
	Nudr_DataRepository "github.com/acore2026/openapi/udr/DataRepository"
	"github.com/acore2026/udm/pkg/app"
)

type ConsumerUdm interface {
	app.App
}

type Consumer struct {
	ConsumerUdm

	// consumer services
	*nnrfService
	*nudrService
	*nudmService
}

func NewConsumer(udm ConsumerUdm) (*Consumer, error) {
	c := &Consumer{
		ConsumerUdm: udm,
	}

	c.nnrfService = &nnrfService{
		consumer:        c,
		nfMngmntClients: make(map[string]*Nnrf_NFManagement.APIClient),
		nfDiscClients:   make(map[string]*Nnrf_NFDiscovery.APIClient),
	}

	c.nudrService = &nudrService{
		consumer:    c,
		nfDRClients: make(map[string]*Nudr_DataRepository.APIClient),
	}

	c.nudmService = &nudmService{
		consumer:      c,
		nfSDMClients:  make(map[string]*Nudm_SubscriberDataManagement.APIClient),
		nfUECMClients: make(map[string]*Nudm_UEContextManagement.APIClient),
	}
	return c, nil
}
