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

	//public routes
	r.Group(func(r chi.Router){
	//server health
	r.Get("/healthz",func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	})
	//auth
	r.Post("/login",server.LoginHandler)
	})

	//protected
	r.Group(func(r chi.Router){
	r.Use(server.AuthMiddleWare)
	//users
	r.Get("/user/{id}",server.GetUserHandler)
	r.Get("/users",server.GetUserListHandler)
	r.Post("/user",server.CreateUserHandler)
	r.Post("/user/{id}",server.UpdateUserHandler)
	r.Post("/user/password/{id}",server.UpdateUserPasswordHandler)
	r.Delete("/user/{id}",server.DeactivateUserHandler)

	//tickets
	r.Get("/ticket/{id}",server.GetTicketHandler)
	r.Get("/tickets",server.GetTicketListHandler)
	r.Post("/ticket",server.CreateTicketHandler)
	r.Post("/ticket/{id}",server.UpdateTicketHandler)
	
	//comments
	r.Post("/comment",server.createComment)
	r.Get("/comments/{id}",server.getTicketComments)
	r.Post("/comment/{id}",server.udpateComment)

	//categories
	r.Post("/categories",server.CreateCategoryHandler)
	r.Get("/categories",server.GetAllCategoriesHandler)
	})

	server.router = r

}

func (server *Server) Start(port string) error {
	return http.ListenAndServe(port,server.router)
}