// calder archive
// 2021-03-29
// <aq@okaq.com>
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	INDEX = "web.html"
	ARCHIVE = "https://calder.org/archive/all/works/"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("web server running on localhost:8080")
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// save image to disk
}

func ArchiveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	res, err := http.Get(ARCHIVE)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(body)
}

func main() {
	motd()
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/a", SaveHandler)
	http.HandleFunc("/b", ArchiveHandler)
	http.ListenAndServe(":8080", nil)
}


