package main

import "net/http"



func main3() {

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}

}