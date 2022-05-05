package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

//login处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {

	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.Id > 0 {
		//登录成功
		//fmt.Println("登录成功：", user)
		//若已是登录状态，则不创建cookie生成新的session

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) //验证（对比密码）
		if err != nil {
			fmt.Println("登录失败，请检查输入的密码。")
			t := template.Must(template.ParseFiles("views/pages/user/logining.html"))
			t.Execute(w, "登录失败，请检查输入的密码。")
		} else {
			fmt.Println("登录成功")
			_, Islogin := IsLogin(r)
			if Islogin == false {
				//生成UUID
				uuid := utils.CreateUUID()
				//创建session
				sess := &model.Session{
					SessionId: uuid,
					UserName:  user.Username,
					UserId:    user.Id,
				}
				dao.AddSession(sess)
				//创建cookie与sesion绑定
				cookie := &http.Cookie{
					Name:     "user",
					Value:    uuid,
					HttpOnly: true,
				}
				http.SetCookie(w, cookie)
			}

			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		}

	} else {
		//登录失败
		fmt.Println("登录失败,请检查输入的用户名")
		t := template.Must(template.ParseFiles("views/pages/user/logining.html"))
		t.Execute(w, "登录失败，请检查输入的用户名。")
	}

}

//退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//删除session
		dao.DelSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将设置后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//回到首页
	QueryBooksByPrice(w, r)
}
func ReLogin(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//删除session
		dao.DelSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将设置后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	t := template.Must(template.ParseFiles("views/pages/user/login.html"))
	t.Execute(w, r)
}

//判断是否有用户登录
func IsLogin(r *http.Request) (*model.Session, bool) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//获取session
		session, _ := dao.GetSession(cookieValue)
		if session.UserId > 0 {
			return session, true //已登录
		} else {
			return nil, false //未登录
		}
	} else {
		return nil, false //未登录
	}

}

//regist处理用户注册的函数
func Regist(w http.ResponseWriter, r *http.Request) {

	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中验证用户名的方法
	user, _ := dao.CheckUserName(username)
	//fmt.Println("user.Id=", user.Id)
	if user.Id > 0 {
		//用户存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！请重新输入。")
	} else {
		//用户名可用
		// 给密码加密
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
		if err != nil {
			fmt.Println(err)
		}
		encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不一样，只需保存一份便可
		// fmt.Println(encodePWD)

		dao.SaveUser(username, encodePWD, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func FindByUserName(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	//调用userdao中验证用户名的方法
	user, _ := dao.CheckUserName(username)
	//fmt.Println("FindUserByName user.Id=", user.Id)
	if user.Id > 0 {
		//用户存在
		w.Write([]byte("用户已存在，请重新输入"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用</font>"))
	}
}
