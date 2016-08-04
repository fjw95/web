package main

import "fmt"

import "html/template"
import "net/http"

//import "path/filepath"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var data = map[string]string{
			"Name": "john wick",
			"Message": "have a	nice day",
		}

		var t, err = template.New("template.html").ParseFiles("/home/fiky/work/src/template.html")
		if err != nil {

			fmt.Println(err.Error())
		}
		t.Execute(w, data)
	})

	fmt.Println("starting web server at	http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
