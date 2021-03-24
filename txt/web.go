// async font render
// AQ <aq@okaq.com>
// 2021-03-18
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "web.html"
	INDEX2 = "web2.html"
	DATA = "txt_0.json"
)

func motd() {
	fmt.Println("font load and display...")
	fmt.Println(time.Now().String())
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,DATA)
}

func Web2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX2)
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/a", DataHandler)
	http.HandleFunc("/b", Web2Handler)
	// http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img/"))))
	http.ListenAndServe(":8080", nil)
}


