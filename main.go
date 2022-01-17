package main

import (
	"log"
	"net/http"
	"os"
)

func main(){
	log.Print("Starting server")
	http.HandleFunc("/hello",func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("Hello world"))
	})
	http.HandleFunc("/bye",func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("Good-bye world"))
	})
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}
	log.Printf("listening on port %s",port)

   err:=http.ListenAndServe(":"+port,nil)
   if err!=nil {
   		log.Print("couldn't start server ",err)
   }
}
