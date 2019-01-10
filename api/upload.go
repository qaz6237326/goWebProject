package api

import (
	"net/http"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"html/template"
	"os"
)

func init() {
	http.HandleFunc("/upload", Upload)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles(ViewsFilePath + "upload.html")
		t.Execute(w,token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		// 保存在根目录的test文件夹
		f, err := os.OpenFile("./test/" + handler.Filename,os.O_WRONLY | os.O_CREATE,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		// 储存文件
		io.Copy(f, file)

	}
}
