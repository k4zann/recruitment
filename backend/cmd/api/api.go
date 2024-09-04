package api

import (
	"log"
	"net/http"
	"recruitment/service/employee"
	"recruitment/service/employer"
	"recruitment/service/resume"
	"recruitment/service/vacancy"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiServer struct {
	addr string
	db   *mongo.Database
}

func NewApiServer(addr string, db *mongo.Database) *ApiServer {
	if db == nil {
		log.Fatalf("Failed to initialize ApiServer: Database is nil")
	}
	return &ApiServer{addr: addr, db: db}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api").Subrouter()

	// employee routes
	employeeStore := employee.NewStore(s.db.Collection("employee"))
	employeeHandler := employee.NewEmployeeHandler(employeeStore)
	employeeHandler.RegisterRoutes(subRouter)

	// employer routes
	employerStore := employer.NewStore(s.db.Collection("employers"))
	employerHandler := employer.NewEmployerHandler(employerStore)
	employerHandler.RegisterRoutes(subRouter)

	// resume routes
	resumeStore := resume.NewStore(s.db.Collection("resume"))
	resumeHandler := resume.NewResumeHandler(resumeStore)
	resumeHandler.RegisterRoutes(subRouter)

	// vacancy routes
	vacancyStore := vacancy.NewStore(s.db.Collection("vacancy"))
	vacancyHandler := vacancy.NewVacancyHandler(vacancyStore)
	vacancyHandler.RegisterRoutes(subRouter)

	log.Println("Listening on ", s.addr)

	if err := http.ListenAndServe(s.addr, router); err != nil {
		log.Fatal("Server failed to start: ", err)
	}

	return nil
}
