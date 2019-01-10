package cookie

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.HandleFunc("/delCookie", delCookie)
	http.HandleFunc("/addCookie", addCookie)
}

// 设置cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	cookie := http.Cookie{
		Name:   "IamCookie",
		Value:  "asd123",
		Path:   "/",
		MaxAge: 86400,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w,"HELLO, I AM Cookie")
}

// 增加cookie，使用w.Header().Add("Set-Cookie", cookie对象)
func addCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "vanyar",
		HttpOnly: true,
	}
	if r.RequestURI == "/favicon.ico" {
		return
	}
	w.Header().Add("Set-Cookie",c1.String())
}

// 获取cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	// 根据key来获取某个cookie
	cookie, err := r.Cookie("IamCookie")
	if err != nil {
		fmt.Println(err)
		fmt.Println("IamCookie不存在")
	}
	fmt.Println(" r.Cookie: ", cookie)

	// 使用r.Header["Cookie"]获取所有的cookie，返回一个数组
	h := r.Header["Cookie"]
	fmt.Println("r.Header['Cookie']: ", h)

	// 使用 r.Cookies() 获取所有的cookie
	c2 := r.Cookies()
	fmt.Println("r.Cookies(): ", c2)

	fmt.Fprintf(w,"HELLO, I AM 获取cookie")

}

// 删除cookie
func delCookie(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}

	cookie := http.Cookie{
		Name:   "IamCookie",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w,"HELLO, I AM delCookie")
}