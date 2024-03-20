package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func respondWithJSON(w http.ResponseWriter,code int,payload interface{}){
	data,err:=json.Marshal(payload)
	if err!=nil{
		log.Printf("Failed to marshal respond to %v",payload)
		w.WriteHeader(500)
		
	}
	w.Header().Add("content type","JSON ")
	w.Write(data)
	w.WriteHeader(code)
    

}