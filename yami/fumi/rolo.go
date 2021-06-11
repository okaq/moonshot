// okaq web serve
// bitmap sample and encode base64 json
// AQ <aq@okaq.com>
// 2019-12-06
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	INDEX = "rolo.html"
	// png dir
	PNG = "omg/"
	// json dir
	JSON = "pow/"
)

var (
	// png file list
	P map[string]string
	// json encoded png file list
	J []byte
)

func motd() {
	fmt.Println("okaq web serve at localhost:8080")
	fmt.Println(time.Now().String())
}

func RoloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,INDEX)
}

func PngHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// static png server at PNG
	w.Header().Set("Content-Type","application/json")
	w.Write(J)
}

func pictures() {
	// load png files into map
	f0, err := ioutil.ReadDir(PNG)
	if err != nil {
		fmt.Println(err)
	}
	/* debug
	for i0, f1 := range f0 {
		fmt.Printf("file: %s, item: %d\n", f1.Name(), i0)
		s0 := fmt.Sprintf("%s%s", PNG, f1.Name())
		fmt.Printf("path: %s\n", s0)
		s1 := strconv.FormatInt(f1.Size(), 10)
		s2 := fmt.Sprintf("%s:%s", s1, f1.ModTime().String())
		fmt.Printf("size:time::%s\n", s2)
	}
	*/
	for _, f1 := range f0 {
		k0 := fmt.Sprintf("%s%s", PNG, f1.Name())
		s0 := strconv.FormatInt(f1.Size(), 10)
		v0 := fmt.Sprintf("%s|%s", s0, f1.ModTime().String())
		P[k0] = v0
	}
	// fmt.Println(P)
}

func compact() {
	// encode png file list map
	// json for http writer
	var err error
	J, err = json.Marshal(P)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(len(J))
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	// req body bytes buffer
	b0 := new(bytes.Buffer)
	b0.ReadFrom(r.Body)
	b1, err := json.Marshal(b0.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	s0 := fmt.Sprintf("%s%d.json", JSON, time.Now().UnixNano())
	ioutil.WriteFile(s0,b1,0666)
	w.Write([]byte("ok save"))
}

func main() {
	motd()
	// load png file list
	P = make(map[string]string)
	pictures()
	compact()
	http.HandleFunc("/", RoloHandler)
	http.HandleFunc("/a", PngHandler)
	http.Handle("/omg/", http.StripPrefix("/omg/", http.FileServer(http.Dir("omg/"))))
	http.HandleFunc("/b", SaveHandler)
	http.ListenAndServe(":8080", nil)
}



