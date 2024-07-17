package main

import (
	"fmt"
	"net/http"
)

func hello (w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world!\n")
}


func headers (w http.ResponseWriter, req * http.Request) {
	for name, header := range req.Header {
		for _, c := range header{
			fmt.Fprintf(w, "%v:%v\n", name, c)
			
		}
	}
}

func main1() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)




	http.ListenAndServe(":8000", nil)
}