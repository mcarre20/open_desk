package server

import (
	"database/sql"
	"fmt"
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
func NewServer(config util.Config) (*Server, error){
	//connect to database
	fmt.Println(config.DBurl)
	d, err := sql.Open("mysql",config.DBurl)
	if err != nil{
		return &Server{}, fmt.Errorf("error connecting to database:\n %v",err)
	}
	store := db.New(d)
	
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

	r.Get("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World!"))
	})

	server.router = r

}

func (server *Server) Start(port string) error {
	return http.ListenAndServe(port,server.router)
}