// CFP init
// 2021-05-04
// aq@okaq.com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "wami.html"
)

func motd() {
	fmt.Println("okaq web on localhost:8080")
	fmt.Println(time.Now().String())
}

func WamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", WamiHandler)
	http.ListenAndServe(":8080", nil)
}

// cache map[string]string
// chan message queue
// json api


