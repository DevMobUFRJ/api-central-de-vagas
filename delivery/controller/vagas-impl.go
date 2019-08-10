package controller

type Resource struct {
	Service service.Vagas
}

func NewController(service service.Vagas) Vagas {
	return &Resource{
		Service: service,
	}
}
