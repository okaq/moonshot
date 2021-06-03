// txt lay out
// 2021-06-04
// aq @ okaq . com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "sami.html"
)

func motd() {
	fmt.Println("okaq web localhost:8080")
	fmt.Println(time.Now().String())
}

func SamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", SamiHandler)
	http.ListenAndServe(":8080", nil)
}


