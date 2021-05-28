// alpha pastel cloud
// 2021-05-27
// aq@okaq.com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "uami.html"
)

func motd() {
	fmt.Println("okaq web server localhost:8080")
	fmt.Println(time.Now().String())
}

func UamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", UamiHandler)
	http.ListenAndServe(":8080", nil)
}


