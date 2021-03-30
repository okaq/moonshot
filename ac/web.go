// calder archive
// 2021-03-29
// <aq@okaq.com>
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	INDEX = "web.html"
	ARCHIVE = "https://calder.org/archive/all/works/"
	HTML = "html/ac.html"
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	fmt.Printf("%s\n", body)
	fmt.Println(res.StatusCode)
	w.Write(body)
}

func arc() {
	res, err := http.Get(ARCHIVE)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	fmt.Printf("%s\n", body)
	fmt.Println(res.StatusCode)
}

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// http.ServeFile(w,r,HTML)
	b0, err := ioutil.ReadFile(HTML)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-type", "text/plain")
	w.Write(b0)
}

func TextHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("ok"))
}

func main() {
	motd()
	// arc()
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/a", SaveHandler)
	http.HandleFunc("/b", ArchiveHandler)
	http.HandleFunc("/c", HtmlHandler)
	http.HandleFunc("/d", TextHandler)
	http.ListenAndServe(":8080", nil)
}


