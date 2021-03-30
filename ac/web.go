// calder archive
// 2021-03-29
// <aq@okaq.com>
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	INDEX = "web.html"
	ARCHIVE = "https://calder.org/archive/all/works/"
	HTML = "html/ac.html"
)

var (
	Img []string
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

func ImgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	b0, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	f0, err := os.Create("html/img.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f0.Close()
	n0, err := f0.Write(b0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("image json %d bytes\n", n0)
	var j0 []string
	err = json.Unmarshal(b0, &j0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("image list count: %d\n", len(j0))
	Img = j0
	go pop()
	s0 := fmt.Sprintf("bytes decoded: %d", len(b0))
	b1 := []byte(s0)
	w.Header().Set("Content-type", "text/plain")
	w.Write(b1)
}

func pop() {
	// for i, img0 := range Img {
	for i := 1015; i < len(Img); i++ {
		img0 := Img[i]
		fmt.Println(i,img0)
		res, err := http.Get(img0)
		if err != nil {
			fmt.Println(err)
		}
		s0 := fmt.Sprintf("img/%05d.jpeg", i)
		fmt.Printf("creating file: %s\n", s0)
		f0, err := os.Create(s0)
		if err != nil {
			fmt.Println(err)
		}
		defer f0.Close()
		b0, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		n0, err := f0.Write(b0)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("wrote %d bytes\n", n0)
	}
	// too many files open, close on sync, not defer
	// 1300 jpegs, 386 MB calder archive on disk
}

func main() {
	motd()
	// arc()
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/a", SaveHandler)
	http.HandleFunc("/b", ArchiveHandler)
	http.HandleFunc("/c", HtmlHandler)
	http.HandleFunc("/d", TextHandler)
	http.HandleFunc("/e", ImgHandler)
	http.ListenAndServe(":8080", nil)
}


