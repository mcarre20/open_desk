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



func(server *Server) CreateTicketHandler (w http.ResponseWriter, r *http.Request){
	type ticketDataReq struct{
		User_id uuid.UUID `json:"user_id"`
		Title string `json:"title"`
		Description string `json:"description"`
	}
	defer r.Body.Close()

	ticket := ticketDataReq{}

	err := util.JsonDecode(r.Body,&ticket)
	if err != nil{
		msg:= "error reading request body"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}

	//create ticket in db
	dbTicket, err := server.store.CreateTicket(r.Context(),db.CreateTicketParams{
		UserID: ticket.User_id,
		Title: ticket.Title,
		Description: ticket.Description,
	})
	if err != nil {
		msg := "error creating ticket"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}
	err = util.RespondWithJson(w,http.StatusOK,dbTicket)
	if err != nil{
		util.RespondWithError(w,"server error",http.StatusInternalServerError)
	}
}

func(server *Server) GetTicketHandler (w http.ResponseWriter,r *http.Request){
	//user id
	ticketId := chi.URLParam(r,"id")
	ticketIdInt, err := strconv.ParseInt(ticketId,0,64)
	if err != nil{
		msg:= "unrecognized ticket id"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}

	//get ticket from db
	dbticket, err := server.store.GetTicket(r.Context(),ticketIdInt)
	if err != nil {
		msg := "error retrieving ticket"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	err = util.RespondWithJson(w,http.StatusOK,dbticket)
	if err != nil{
		util.RespondWithError(w,"server error",http.StatusInternalServerError)
	}
}

func(server *Server) GetTicketListHandler (w http.ResponseWriter,r *http.Request){
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
	dbTickets, err := server.store.GetTicketList(r.Context(),db.GetTicketListParams{
		Limit: int32(pageSize),
		Offset: int32(dbOffset),
	})

	if err != nil {
		msg:= "error retrieving data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}
	

	type ticketListResponse struct{
		Tickets []db.Ticket `json:"tickets"`
	}
	err = util.RespondWithJson(w,http.StatusAccepted,ticketListResponse{
		Tickets: dbTickets,
	})

	if err != nil{
		msg := "error sending data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}
	
}

func(server *Server) UpdateTicketHandler (w http.ResponseWriter,r *http.Request){
	ticketId := chi.URLParam(r,"id")
	ticketIdInt,err := strconv.ParseInt(ticketId,0,64)
	if err != nil{
		msg:= "unrecognized ticket id"
		util.RespondWithError(w,msg,http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	ticketDataReq := db.Ticket{}
	err = util.JsonDecode(r.Body,&ticketDataReq)
	if err != nil{
		msg := "error reading json"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//update ticket
	dbticket, err := server.store.UpdateTicket(r.Context(),db.UpdateTicketParams{
		ID: ticketIdInt,
		AssignedTo: ticketDataReq.AssignedTo,
		Status: ticketDataReq.Status,
		Priority: ticketDataReq.Priority,
		CategoryID: ticketDataReq.CategoryID,
	})
	if err != nil {
		log.Println(err)
		msg := "error updating the ticket"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
		return
	}

	//response
	err = util.RespondWithJson(w, http.StatusOK,dbticket)
	if err != nil {
		msg := "error sending data"
		util.RespondWithError(w,msg,http.StatusInternalServerError)
	}

}