// yami = "dark" (jp)
// port OKAQ V01D to webgl 2.0
// 2021-04-24
// aq@okaq.com
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INDEX = "yami.html"
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

// to be hosted on okaq.gtihub.io
// no need to write log files for now


