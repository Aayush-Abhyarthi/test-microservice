package main

import "net/http"

func handleHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func main(){
	
	http.HandleFunc("/", handleHome)
	http.ListenAndServe(":8000", nil)

}