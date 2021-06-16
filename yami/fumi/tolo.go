// bitmap sample & encode
// aq@okaq.com
// 2021-06-16
package main

import (
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

func main() {
	motd()
	load()
	http.HandleFunc("/", ToloHandler)
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


