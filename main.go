package main

import (
	"api-central-de-vagas/delivery/routes"
	"api-central-de-vagas/resources/config"
	"api-central-de-vagas/resources/injection"
)

func init() {
	cfg := config.NewViperConfig()
	injection.Inject(cfg)
}

func main() {
	routes.Routes()
}
