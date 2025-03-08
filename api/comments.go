package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	db "github.com/mcarre20/open_desk/db/sqlc"
	"github.com/mcarre20/open_desk/util"
)
type commnetReq struct{
	UserID uuid.UUID `json:"user_id"`
	TicketID int64 `json:"ticket_id"`
	Comment string `json:"comment"`
	CustomerVisible bool `json:"customer_visible"`
}

func(server *Server) createComment (w http.ResponseWriter,r *http.Request){
	
	defer r.Body.Close()
	comment := commnetReq{}
	err := util.JsonDecode(r.Body,&comment)
	if err != nil {
		msg := "error reading request data"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}

	//create comment in db
	dbComment, err := server.store.CreateComment(r.Context(),db.CreateCommentParams{
		UserID: comment.UserID,
		TicketID: comment.TicketID,
		Comments: comment.Comment,
		CustomerVisible: comment.CustomerVisible,
	})
	if err != nil {
		msg := "error creating comment"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//send response
	err = util.RespondWithJson(w,http.StatusOK,dbComment)
	if err != nil {
		msg := "error sending data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
	}

}

func (server *Server) getTicketComments(w http.ResponseWriter,r *http.Request){
	ticketId := chi.URLParam(r,"id")
	ticketIdInt,err := strconv.ParseInt(ticketId,0,64)
	if err != nil{
		msg:= "unrecognized ticket id"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}

	//get comments from db
	dbComments, err := server.store.GetTicketComments(r.Context(),ticketIdInt)
	if err != nil {
		msg := "error get comments"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//respond with data
	err = util.RespondWithJson(w, http.StatusOK,dbComments)
	if err != nil {
		msg := "error sending data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
	}
}

func (server *Server) udpateComment(w http.ResponseWriter,r *http.Request){
	commentId := chi.URLParam(r,"id")
	commentIdInt,err := strconv.ParseInt(commentId,0,64)
	if err != nil{
		msg:= "unrecognized ticket id"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	comment := commnetReq{}
	err = util.JsonDecode(r.Body,&comment)
	if err != nil {
		msg := "error reading request data"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}

	//update comment in db
	dbComment, err := server.store.UpdateComment(r.Context(),db.UpdateCommentParams{
		ID: commentIdInt,
		Comments: comment.Comment,
		CustomerVisible: comment.CustomerVisible,
	})
	if err != nil {
		msg := "error updating comment"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//respond with data
	err = util.RespondWithJson(w,http.StatusOK,dbComment)
	if err != nil {
		msg := "error sending data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
	}
}