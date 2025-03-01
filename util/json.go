package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

//decodes request body to JSON.
//takes an io.Reader such as r.body and store in pointer v
func JsonDecode(r io.Reader,v any)error{
	decoder := json.NewDecoder(r)
	err := decoder.Decode(v)
	if err != nil{
		return err
	}
	return nil
}

//Takes http ResponseWriter, http status code, and payload
// convert payload to json and sends http response to client
func RespondWithJson(w http.ResponseWriter,c int,p any) error{
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(p)
	if err != nil {
		log.Println("error marshalling data")
		return err
	}
	w.WriteHeader(c)
	w.Write(data)
	return nil
}

func RespondWithError(w http.ResponseWriter,m string,c int){
	type errMsg struct{
		Error string `json:"error"`
	}
	w.Header().Set("Content-Type", "application/json")

	data,err := json.Marshal(errMsg{
		Error: m,
	})
	if err != nil{
		log.Println("error responding with error message")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(c)
	w.Write(data)
}