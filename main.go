package main

import (
	"log"

	"github.com/mcarre20/open_desk/server"
	"github.com/mcarre20/open_desk/util"
)

func main(){

	//Load app config
	log.Println("loading config...")
	config, err := util.LoadConfig("./.env")
	if err != nil{
		log.Fatal(err)
	}

	//setup server
	log.Println("setting up server...")
	s, err :=server.NewServer(config)
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