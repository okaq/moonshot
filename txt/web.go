// async font render
// AQ <aq@okaq.com>
// 2021-03-18
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	INDEX = "web.html"
	INDEX2 = "web2.html"
	DATA = "txt_0.json"
)

var (
	tick int
)

func motd() {
	fmt.Println("font load and display...")
	fmt.Println(time.Now().String())
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,DATA)
}

func Web2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX2)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// r.ParseMultipartForm(10 << 20)
	f0, h0, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
	}
	defer f0.Close()
	fmt.Println(f0)
	fmt.Printf("File: %+v\n", h0.Filename)
	fmt.Printf("Size: %+v\n", h0.Size)
	fmt.Printf("Mime: %+v\n", h0.Header)
	s0 := fmt.Sprintf("bytes recieved: %+v", h0.Size)
	b0 := []byte(s0)
	s1 := fmt.Sprintf("img/%04d.png", tick)
	f1, err := os.Create(s1)
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()
	f2, _ := h0.Open()
	b1, _ := ioutil.ReadAll(f2)
	fmt.Printf("bytes read: %+v\n", len(b1))
	n0, _ := f1.Write(b1)
	fmt.Printf("bytes written: %+v\n", n0)
	tick = tick + 1
	w.Write(b0)
}

func main() {
	motd()
	tick = 0
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/a", DataHandler)
	http.HandleFunc("/b", Web2Handler)
	http.HandleFunc("/c", SaveHandler)
	// http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("img/"))))
	http.ListenAndServe(":8080", nil)
}


