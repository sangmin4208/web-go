package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe("192.168.0.8:8080", nil)
}
func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	http.SetCookie(w, &http.Cookie{
		Name:    "my-name",
		Value:   "Nacho",
		Path:    "/",
		Expires: time.Now().Add(10 * time.Second),
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	// c1, err := req.Cookie("my-cookie")
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	// }
	cookies := req.Cookies()
	fmt.Fprintln(w, "YOUR COOKIE:", cookies)
}
