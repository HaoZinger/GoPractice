package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

type user struct {
	Name string
	Age  int
}

func (u user) String() string {
	return fmt.Sprintf("user{name=%s,age=%d}", u.Name, u.Age)
}

func main() {
	formDemo()
}

func formDemo() {
	http.HandleFunc("/form", formHandler)
	if error := http.ListenAndServe(":80", nil); error != nil {
		panic(error)
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		writer.Write([]byte(form))
	case http.MethodPost:
		var user user
		bytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(bytes, &user)
		writer.Write([]byte(user.String()))
	}
}
