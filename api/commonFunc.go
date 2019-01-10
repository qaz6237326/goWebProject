package api

import (
	"net/http"
	"fmt"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

var checkboxs = map[string]string{}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles(ViewsFilePath + "login.html")
		t.Execute(w, token)
	} else {
		r.ParseForm()

		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token存在，开始验证token合法性，token: ", token)
		} else {
			fmt.Println("token不存在，出错了")
		}

		fmt.Println("username: ", r.Form["username"])
		fmt.Println("passowrd: ", r.Form["password"])
		// 这个是字段是一个checkbox，得到的是一个数组
		fmt.Println("interest: ", r.Form["interest"])


		/* 对用户输入的内容转义
		func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
		func HTMLEscapeString(s string) string //转义s之后返回结果字符串
		func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串
		*/
		template.HTMLEscape(w,[]byte(r.Form.Get("username")))


		// 为了检验用户提交给后端的是不是我们定义的数据，后端需要一个包含所有checkbox的map
		//checkboxs["football"] = "football"
		//checkboxs["basketball"] = "basketball"
		//checkboxs["tennis"] = "tennis"
		//fmt.Println(r.Form["interset"])

		// 如果 兴趣项 的数组长度不等于0，那就表示有选择兴趣 复选框
		//if len(r.Form["interest"]) != 0 {
		//	// 循环遍历 这个数组
		//	for _, v := range r.Form["interest"] {
		//		// 把循环出来的每一项跟 预设的值做比较，如果存在，则表示这个值是正确的。不存在，则前端发送了一个预设之外的值过来，表示异常。
		//		if val, ok := checkboxs[v]; ok {
		//			fmt.Println("复选空的值", val)
		//		} else {
		//			fmt.Println("不存在的值", v)
		//		}
		//	}
		//}

		// slice := []string{"football", "basketball", "tennis"}


		// 使用转化为int类型的方法来检查是否是正整数和空值
		/*getInt, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Fprintf(w, "年龄为空，或者格式错误")
			fmt.Println("年龄为空，或者格式错误")
			return
		} else {
			fmt.Println("age: ", getInt)
		}*/

		// 使用正则匹配 m 是 Boole类型
		//if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		//	fmt.Fprintf(w, "年龄为空，或者格式错误")
		//	fmt.Println("年龄为空，或者格式错误")
		//	return
		//} else {
		//	fmt.Println("age: ", m)
		//}

		// 用 Get()函数来获取值，返回值类型是string，如果没有，则返回一个空值 ""
		//if r.Form.Get("username") == "" {
		//	fmt.Fprintf(w, "用户名为空")
		//	return
		//}
		//if r.Form.Get("password") == "" {
		//	fmt.Fprintf(w, "密码为空")
		//	return
		//}

		// 返回一个字符串
		// fmt.Fprintf(w, "登录成功！Welcom")
	}
}