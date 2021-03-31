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
)

func motd() {
	fmt.Println(time.Now().String())
}

func XonixHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", XonixHandler)
	http.ListenAndServe(":8080", nil)
}


