package injection

import (
	"api-central-de-vagas/delivery/controller"
	"api-central-de-vagas/persistance/repository"
	"api-central-de-vagas/resources/database"
	"api-central-de-vagas/usecase/service"
	"firebase.google.com/go/auth"
	"github.com/globalsign/mgo"
	"github.com/karlkfi/inject"
	"os"
)

var (
	graph        inject.Graph
	Repository   repository.Vagas
	Service      service.Vagas
	Controller   controller.Vagas
	MongoSession *mgo.Session
	FirebaseAuth *auth.Client
)

func Inject() {

	MongoSession = database.MongoDBConnect(os.Getenv("MONGO_DB"))
	panic("asd")
	FirebaseAuth = database.FirebaseAuthConnect()

	graph = inject.NewGraph()

	graph.Define(&Controller, inject.NewProvider(controller.NewController, &Service))
	graph.Define(&Service, inject.NewProvider(service.NewService, &Repository, &FirebaseAuth))
	graph.Define(&Repository, inject.NewProvider(repository.NewRepository, &MongoSession))

	graph.ResolveAll()
}
