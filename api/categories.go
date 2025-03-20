package api

import (
	"log"
	"net/http"

	"github.com/mcarre20/open_desk/util"
)

func (server *Server) CreateCategoryHandler(w http.ResponseWriter,r *http.Request){
	type req struct{
		Category string `json:"category"`
	}

	defer r.Body.Close()

	category := req{}
	err := util.JsonDecode(r.Body, &category)
	if err != nil{
		log.Print("error decoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//create category in db
	dbCategory, err := server.store.CreateCategory(r.Context(),category.Category)
	if err != nil{
		log.Println(err)
		msg:="error creating user"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}
	
	//response
	err = util.RespondWithJson(w,http.StatusAccepted,dbCategory)
	if err != nil{
		util.RespondWithError(w,"server error",http.StatusInternalServerError)
	}
}

func (server *Server) GetAllCategoriesHandler(w http.ResponseWriter,r *http.Request){
	dbCategories, err := server.store.GetAllCategories(r.Context())
	if err != nil{
		log.Println(err)
		msg:="error creating user"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//response
	err = util.RespondWithJson(w,http.StatusAccepted,dbCategories)
	if err != nil{
		util.RespondWithError(w,"server error",http.StatusInternalServerError)
	}
}