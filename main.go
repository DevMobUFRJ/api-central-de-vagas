package main

import (
	"api-central-de-vagas/delivery/routes"
	"api-central-de-vagas/resources/config"
	"api-central-de-vagas/resources/injection"
	"io/ioutil"
	"os"
)

func init() {
	credentials := []byte(os.Getenv("FIREBASE_CREDENTIALS_JSON"))
	err := ioutil.WriteFile("./firebase-adminsdk.json", credentials, 0644)
	if err != nil {
		panic(err)
	}

	cfg := config.NewViperConfig()
	injection.Inject(cfg)
}

func main() {
	routes.Routes()
}
