package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	name string
	age  int
}

func (p *User) String() string {
	return fmt.Sprintf("[name=%s;age=%d]", p.name, p.age)
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("Inside HelloServer handler")
	name := req.URL.Path[len("/hello/"):]
	u := &User{name, 21}
	if req.Method == http.MethodGet {
		fmt.Fprintln(w, u)
	}
	switch req.Method {
	case http.MethodGet:
		fmt.Println(req.Method)
	case http.MethodPost:
		all, err:= ioutil.ReadAll(req.Body)
		checkErr(err)
		fmt.Printf("%s\n", string(all))
	default:
		fmt.Println("default")
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello/", handleHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
