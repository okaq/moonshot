// bitmap sample & encode
// aq@okaq.com
// 2021-06-16
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	INDEX = "tolo.html"
	IMG = "img/"
)

var (
	Images []string
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("okaq web localhost:8080")
}

func ToloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func load() {
	// populate img director file names
	f0, err := ioutil.ReadDir(IMG)
	if err != nil {
		fmt.Println(err)
	}
	Images = make([]string, len(f0))
	for i, f1 := range f0 {
		f2 := fmt.Sprintf("%s%s", IMG, f1.Name())
		Images[i] = f2
	}
	fmt.Println(Images)
}

func PngHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// w.Write([]byte("ok go"))
	b0, err := json.Marshal(Images)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("encoding %d bytes json png list\n", len(b0))
	w.Header().Set("Content-type", "application/json")
	w.Write(b0)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// read req body into bytes buffer
	b0 := new(bytes.Buffer)
	b0.ReadFrom(r.Body)
	b1, err := json.Marshal(b0.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	s0 := fmt.Sprintf("%s%s.json", JSON, time.Now().UnixNano())
	ioutil.WriteFile(s0,b1,0666)
	s1 := fmt.Sprintf("file %s written %d bytes", s0, len(b1))
	b2 := []byte(s1)
	w.Write(b2)

	// buffered chan reciever for bytes
	// file name from dir list
	// json response {name,size}
}

func main() {
	motd()
	load()
	http.HandleFunc("/", ToloHandler)
	http.HandleFunc("/a", PngHandler)
	http.HandleFunc("/b", SaveHandler)
	fs0 := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/",fs0))
	http.ListenAndServe(":8080", nil)
}

// auto gen from PNG files
// workflow: user generates images in draw tool
// format: 1024 x 1024 png
// files are initially saved to img/ dir
// on web host start, json png file list created from os
// at first client request, xhr retrieves file list
// one by one, each file is fetched and rendered
// sampled as 1024 monochrome bit array
// white = 1, black = 0
// bit array is post'd to web server handler
// saved to file in base64 json binary string encoding
// file name preserved -> img/goop.png json/goop.json


