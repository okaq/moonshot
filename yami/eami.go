// alpha pastel big cloud
// 2021-05-31
// aq@okaq.com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "tami.html"
)

func motd() {
	fmt.Println("okaq web localhost:8080")
	fmt.Println(time.Now().String())
}

func TamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", TamiHandler)
	http.ListenAndServe(":8080", nil)
}


