package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mcarre20/open_desk/api"
	db "github.com/mcarre20/open_desk/db/sqlc"
	"github.com/mcarre20/open_desk/util"
)

func main(){

	//Load app config
	log.Println("loading config...")
	config, err := util.LoadConfig("./.env")
	if err != nil{
		log.Fatal(err)
	}

	//connect to db
	d, err := sql.Open("postgres",config.DBurl)
	if err != nil{
		log.Fatal(err)
	}
	store := db.New(d)

	//setup server
	log.Println("setting up server...")
	s, err :=api.NewServer(config,store)
	if err != nil{
		log.Fatal(err)
	}


	//start server
	log.Printf("server starting and listening on Port %v",config.ServerPort)
	err = s.Start(":"+config.ServerPort)
	if err != nil {
		log.Fatal(err)
	}
}