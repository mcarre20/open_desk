package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	db "github.com/mcarre20/open_desk/db/sqlc"
	"github.com/mcarre20/open_desk/util"
)

const pageDefaultLimit = 20

type userResponse struct{
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	UserRole          int32     `json:"user_role"`
}

func (server *Server)CreateUserHandler(w http.ResponseWriter,r *http.Request){
	//get user for request
	type userReq struct{
		Username          string    `json:"username"`
		FirstName         string    `json:"first_name"`
		LastName          string    `json:"last_name"`
		Email             string    `json:"email"`
		Password		  string    `json:"password"`
	}
	defer r.Body.Close()
	user := userReq{}
	err := util.JsonDecode(r.Body,&user)
	if err != nil{
		log.Print("error decoding json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//to-do hash password before saving in database

	//create user in db
	dbUser, err := server.store.CreateUser(r.Context(),db.CreateUserParams{
		Username: user.Username,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		HashedPassword: user.Password,
		UserRole: 1,
	})

	if err != nil{
		log.Println(err)
		msg:="error creating user"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	
	//respond with user data
	err = util.RespondWithJson(w,http.StatusCreated,userResponse{
		ID: dbUser.ID,
		Username: dbUser.Username,
		FirstName: dbUser.FirstName,
		LastName: dbUser.LastName,
		Email: dbUser.Email,
		UserRole: dbUser.UserRole,
	})
	if err != nil{
		util.RespondWithError(w,"server error",http.StatusInternalServerError)
	}

}

func (server *Server) GetUserHandler(w http.ResponseWriter, r *http.Request){
	//get user id
	userId := chi.URLParam(r,"id")
	userUUUID, err := uuid.Parse(userId)
	if err != nil{
		log.Println("error converting id to uuid")
		msg:= "unrecognized user id"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}
	//get user from db
	dbUser, err := server.store.GetUser(r.Context(), userUUUID)
	if err != nil{
		msg:="error get user"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}

	//response
	err = util.RespondWithJson(w,http.StatusOK,userResponse{
		ID: dbUser.ID,
		Username: dbUser.Username,
		FirstName: dbUser.FirstName,
		LastName: dbUser.LastName,
		Email: dbUser.Email,
		UserRole: dbUser.UserRole,
	})
	if err != nil{
		msg:="error sending response"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}
}

func (server *Server) GetUserListHandler(w http.ResponseWriter, r *http.Request){
	var pageSize int
	var dbOffset int
	pageSizeReq := r.URL.Query().Get("page_size")
	pageNumberReq:= r.URL.Query().Get("page_number")
	//set default if string is empty 
	if pageSizeReq == "" {
		pageSize = pageDefaultLimit
	}else{
		pageSizeInt, err := strconv.Atoi(pageSizeReq)
		if err != nil {
			msg := "page_size must be a number"
			util.RespondWithError(w,msg,http.	StatusBadRequest)
			return
		}
		pageSize = pageSizeInt
	}
	
	if pageNumberReq == "" {
		dbOffset = 0
	}else{
		pageNumber,err := strconv.Atoi(pageNumberReq)
		if err != nil {
			msg:= "page_number must be a number"
			util.RespondWithError(w, msg,http.StatusBadRequest)
			return
		} 
		dbOffset = (pageNumber - 1)*pageSize
	}
	
	//Get user list
	dbUser, err := server.store.GetUserList(r.Context(),db.GetUserListParams{
		Limit: int32(pageSize),
		Offset: int32(dbOffset),
	})

	if err != nil {
		msg:= "error retrieving data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}
	
	//clearn up DB data and send response
	userListClean := make([]userResponse,0,len(dbUser))
	for _,user := range dbUser {
		userListClean = append(userListClean, userResponse{
			ID: user.ID,
			Username: user.Username,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Email: user.Email,
			UserRole: user.UserRole,
		})
	}
	log.Println(userListClean)
	type userListResponse struct{
		Users []userResponse `json:"users"`
	}
	err = util.RespondWithJson(w,http.StatusAccepted,userListResponse{
		Users: userListClean,
	})

	if err != nil{
		msg := "error sending data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

}

func (server *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request){
	userId := chi.URLParam(r,"id")
	userUUID,err := uuid.Parse(userId)
	if err != nil{
		log.Println("error converting id to uuid")
		msg:= "unrecognized user id"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	user := userResponse{}
	err = util.JsonDecode(r.Body,&user)
	if err != nil{
		msg := "error reading json"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//update user
	dbUser, err := server.store.UpdateUserInfo(r.Context(),db.UpdateUserInfoParams{
		ID: userUUID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		UserRole: user.UserRole,
	})
	if err != nil {
		msg := "error updating the user"
		log.Println(msg)
		log.Println(err)
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//response
	err = util.RespondWithJson(w, http.StatusAccepted,userResponse{
		ID: dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName: dbUser.LastName,
		Email: dbUser.Email,
		UserRole: dbUser.UserRole,
	})

	if err != nil {
		msg := "error sending data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
	}

}

func (server *Server) UpdateUserPassword(w http.ResponseWriter, r *http.Request){}

func (server *Server) DeactivateUserHandler(w http.ResponseWriter,r *http.Request){
	userId := chi.URLParam(r,"id")
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		msg := "error with user ID"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}

	//update database
	err = server.store.DeactivateUser(r.Context(),userUUID)
	if err != nil{
		msg := "error deactivating user"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}