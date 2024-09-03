package api

import (
	"log"
	"net/http"
	"recruitment/service/employee"
	"recruitment/service/employer"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiServer struct {
	addr string
	db   *mongo.Database
}

func NewApiServer(addr string, db *mongo.Database) *ApiServer {
	return &ApiServer{addr: addr, db: db}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api").Subrouter()

	employeeStore := employee.NewStore(s.db.Collection("employees"))

	employeeHandler := employee.NewEmployeeHandler(employeeStore)
	employeeHandler.RegisterRoutes(subRouter)

	employerStore := employer.NewStore(s.db.Collection("employers"))

	employerHandler := employer.NewEmployerHandler(employerStore)
	employerHandler.RegisterRoutes(subRouter)


	log.Fatal("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
