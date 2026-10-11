package main

import (

     "fmt"
     "net/http"
     "encoding/json"

)
// program to haggle internationally lol
// struct for saving the base common marketplace val
type Product struct{
	ID 	      string 'json:"id"'
	TITLE         string 'json:"title"'
	PRICE         string 'json:"price"'
	DESC          string 'json:"desc"'
	LOCATION      string 'json:"location"'
	CURRENCY      string 'json:"currency"' // this should be dyn



}
func main(){

http.HandleFunc("/api/v1/products",func(w http.ResponseWriter, r *http.Request){
fmt.Fprint(w, "Hello Femin, you're connected to mac air now!")
w.Header().Set("Content-Type", "application/json")
exampleItem := Product{
			ID:          "0000001",
			Title:       "Retro Goku Figurine (mint)",
			Price:       4500.00,
			Currency:    "JPY",
			Description: "Authentic item bought in Akihabara. Global shipping available!",
			Location:    "Tokyo, Japan",
		}

json.NewEncoder(w).Encode(exampleItem)


})

fmt.Println("If you see this, the server is set upp and running! Dattebayo")
http.ListenAndServe("0.0.0.0:8080", nil)
}
