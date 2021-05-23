// RW 1D fun
// 2021-05-24
// aq@okaq.com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "vami.html"
)

func motd() {
	fmt.Println("okaq web on localhost:8080")
	fmt.Println(time.Now().String())
}

func VamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", VamiHandler)
	http.ListenAndServe(":8080", nil)
}


