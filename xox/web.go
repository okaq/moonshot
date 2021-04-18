// okaq web serve
// gae golang deploy
// aq@okaq.com
// 2021-04-20
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
	fmt.Println("okaq web serve running on localhost:8080")
	fmt.Println(time.Now().String())
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	http.ListenAndServe(":8080", nil)
}


