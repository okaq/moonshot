// rugger bee nano game
// aq@okaq.com
// 2021-06-29
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "lito.html"
)


func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("okaq web localhost:8080")
}

func KitoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	fs0 := http.FileServer(http.Dir("json"))
	http.Handle("/json/", http.StripPrefix("/json/",fs0))
	http.HandleFunc("/", LitoHandler)
	http.ListenAndServe(":8080", nil)
}


