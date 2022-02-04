package main

import (
	"fmt"
	"net/http"
)


type nacho int
func (n nacho) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Nacho-Key", "Nacho-Value")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main(){
err := http.ListenAndServe(":8080", nacho(0))
if err != nil {
	fmt.Println(err)
}
}
