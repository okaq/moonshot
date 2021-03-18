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
)

func motd() {
	fmt.Println("font load and display...")
	fmt.Println(time.Now().String())
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	// http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img/"))))
	http.ListenAndServe(":8080", nil)
}


