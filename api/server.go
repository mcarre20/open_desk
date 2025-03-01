package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	db "github.com/mcarre20/open_desk/db/sqlc"
	"github.com/mcarre20/open_desk/util"
)


type Server struct{
	config util.Config
	router *chi.Mux
	store *db.Queries
}

// NewServer function takes in a config strutc that contains server config
// function creates connection to dabase and setup route
// and returns pointer to a server struct
func NewServer(config util.Config, store *db.Queries) (*Server, error){
	//config server
	server := &Server{
		config: config,
		store: store,
	}
	//load routes
	server.setupRouter()
	return server,nil
}

func (server *Server) setupRouter(){
	r :=chi.NewRouter()

	//server health
	r.Get("/healthz",func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	})

	//users
	r.Get("/user/{id}",server.GetUserHandler)
	r.Get("/users",server.GetUserListHandler)
	r.Post("/user",server.CreateUserHandler)
	r.Post("/user/{id}",server.UpdateUserHandler)
	r.Delete("/user/{id}",server.DeactivateUserHandler)

	//tickets
	

	server.router = r

}

func (server *Server) Start(port string) error {
	return http.ListenAndServe(port,server.router)
}