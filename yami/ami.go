// CFP deploy
// 2021-05-02
// aq@okaq.com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "zami.html"
)

func motd() {
	fmt.Println("okaq web serve on localhost:8080")
	fmt.Println(time.Now().String())
}

func YamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func main() {
	motd()
	http.HandleFunc("/", YamiHandler)
	http.ListenAndServe(":8080", nil)
}

// set up map[string]string cache
// json api to read, write, update
// message queue using channel

