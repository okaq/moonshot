// xonix full screen
// AQ <aq@okaq.com>
// 2021-04-02
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "xox.html"
	INDEX2 = "xox2.html"
	INDEX3 = "xox3.html"
)

func motd() {
	fmt.Println(time.Now().String())
}

func XonixHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func Xonix2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX2)
}

func Xonix3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX3)
}


func main() {
	motd()
	http.HandleFunc("/", XonixHandler)
	http.HandleFunc("/a", Xonix2Handler)
	http.HandleFunc("/b", Xonix3Handler)
	http.ListenAndServe(":8080", nil)
}


