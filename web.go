// p5js scratch server
// created by AQ <aq@okaq.com>
// 2020-12-02
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
	fmt.Println("how do sir?")
	fmt.Printf("%s\n", time.Now().String())
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	http.Handle("/p5/", http.StripPrefix("/p5", http.FileServer(http.Dir("p5/"))))
	http.ListenAndServe(":8080", nil)
}



