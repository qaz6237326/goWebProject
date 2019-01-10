package cookie

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.HandleFunc("/delCookie", delCookie)
}

// 设置cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	cookie := http.Cookie{Name: "IamCookie", Value: "asd123", Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w,"HELLO, I AM Cookie")
}

// 获取cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	cookie, err := r.Cookie("IamCookie")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cookie)
	fmt.Fprintf(w,"HELLO, I AM 获取cookie")
}

// 删除cookie
func delCookie(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	cookie := http.Cookie{Name: "IamCookie", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w,"HELLO, I AM delCookie")
}