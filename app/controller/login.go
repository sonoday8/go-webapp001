package controller

import (
	"github.com/labstack/echo/v4"
)

// Index return error
func Index(c echo.Context) error {
	c.Logger().Debug("AAAAAAAA")
	return c.File("public/login.html")
}

// Login return error
func Login(c echo.Context) error {
	c.Logger().Debug("BB")
	return c.File("public/login.html")
}

// type LoginForm struct {
// 	UserId       string
// 	Password     string
// 	ErrorMessage string
// }

// func Login(c echo.Context) error {
// 	c.Logger().Debug("AAAAAAAA")
// 	loginForm := LoginForm{
// 		UserId:       c.FormValue("userId"),
// 		Password:     c.FormValue("password"),
// 		ErrorMessage: "",
// 	}
// 	userId := html.EscapeString(loginForm.UserId)
// 	password := html.EscapeString(loginForm.Password)

// 	if userId != "userId" && password != "password" {
// 		loginForm.ErrorMessage = "ユーザーID または パスワードが間違っています"
// 		return c.Render(http.StatusOK, "login", loginForm)
// 	}

// 	//セッションにデータを保存する
// 	session := session.Default(c)
// 	session.Set("loginCompleted", "completed")
// 	session.Save()

// 	// completeJson := CompleteJson{
// 	// 	Success: true,
// 	// }
// 	return c.File("public/login.html")
// }
