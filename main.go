package main

import (

     "fmt"
     "net/http"

)

func main(){

http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
fmt.Fprint(w, "Hello Femin, you're connected to mac air now!")
})

fmt.Println("If you see this, the server is set upp and running! Dattebayo")
http.ListenAndServe("0.0.0.0:8080", nil)
}
