package main

import "net/http"

func main(){
	http.HandleFunc("/hello",func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("Hello world"))
	})
	http.HandleFunc("/bye",func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("Good-bye world"))
	})
   http.ListenAndServe(":8080",nil)
}
