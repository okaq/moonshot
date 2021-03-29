// calder archive
// 2021-03-29
// <aq@okaq.com>
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
	fmt.Println(time.Now().String())
	fmt.Println("web server running on localhost:8080")
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// save image to disk
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/a", SaveHandler)
	http.ListenAndServe(":8080", nil)
}


