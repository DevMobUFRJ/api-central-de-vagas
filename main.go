package main

import (
	"api-central-de-vagas/delivery/routes"
	"api-central-de-vagas/resources/injection"
)

func init() {
	injection.Inject()
}

func main() {
	routes.Routes()
}
