// fumi bit maps
// 2021-06-06
// aq @ okaq . com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "fumi.html"
)

func motd() {
	fmt.Println("okaq web localhost:8080")
	fmt.Println(time.Now().String())
}

func FumiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", FumiHandler)
	http.ListenAndServe(":8080", nil)
}


