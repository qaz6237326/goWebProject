package main

import (
	"net/http"
	"fmt"
	"log"
	_"web/project/api"
	_"web/project/cookie"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 避免执行/favicion.ico的情况
	if r.RequestURI == "/favicon.ico" {
		return
	}
	fmt.Fprintf(w,"HELLO, I AM HELLO MODULE")
}

func main() {

	http.HandleFunc("/", sayHello)


	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}






