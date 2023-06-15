package main

import (
	"b47s1/connection"
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id          int
	Title       string
	Start       time.Time
	End         time.Time
	Duration    string
	Description string
	Reactjs     bool
	Nextjs      bool
	Nodejs      bool
	Typescript  bool
	StartFormat string
	EndFormat   string
	Image       string
	Authorid    int
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	e.Static("/public", "public")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// Get
	e.GET("/", home)
	e.GET("/form-project", formProject)
	e.GET("/detail-project/:id", detailProject)
	e.GET("/contact", contact)
	e.GET("/update-project/:id", formUpdateProject)
	e.GET("/testimonials", testimonials)

	// Register
	e.GET("/form-register", formRegister)
	e.POST("/register", register)

	// Login
	e.GET("/form-login", formLogin)
	e.POST("login", login)

	e.POST("/logout", logout)

	// Post
	e.POST("/add-project", addNewProject)
	e.POST("/update-project/:id", updateProject)
	e.POST("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:65534"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, start_date, end_date, description, reactjs, nextjs, nodejs, typescript, duration FROM tb_projects")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.Title, &each.Start, &each.End, &each.Description, &each.Reactjs, &each.Nextjs, &each.Nodejs, &each.Typescript, &each.Duration)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		each.StartFormat = each.Start.Format("2 January 2006")
		each.EndFormat = each.End.Format("2 January 2006")

		result = append(result, each)
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"Projects":     result,
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "status")
	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), projects)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	session := map[string]interface{}{
		"DataSession": userData,
	}

	return tmpl.Execute(c.Response(), session)
}

func formProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	session := map[string]interface{}{
		"DataSession": userData,
	}

	return tmpl.Execute(c.Response(), session)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	session := map[string]interface{}{
		"DataSession": userData,
	}

	return tmpl.Execute(c.Response(), session)
}

func detailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var projectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, description, reactjs, nextjs, nodejs, typescript, duration FROM tb_projects WHERE id=$1", id).Scan(
		&projectDetail.Id, &projectDetail.Title, &projectDetail.Start, &projectDetail.End, &projectDetail.Description, &projectDetail.Reactjs, &projectDetail.Nextjs, &projectDetail.Nodejs, &projectDetail.Typescript, &projectDetail.Duration)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"Projects":    projectDetail,
		"StartDate":   projectDetail.Start.Format("2 January 2006"),
		"EndDate":     projectDetail.End.Format("2 January 2006"),
		"DataSession": userData,
	}

	var tmpl, errTemp = template.ParseFiles("views/project-detail.html")

	if errTemp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemp.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addNewProject(c echo.Context) error {
	c.Request().ParseForm()

	title := c.FormValue("inputTitle")
	start := c.FormValue("inputStart")
	end := c.FormValue("inputEnd")
	description := c.FormValue("inputDescription")

	var reactjs bool
	if c.FormValue("reactjs") == "reactjs" {
		reactjs = true
	}

	var nextjs bool
	if c.FormValue("nextjs") == "nextjs" {
		nextjs = true
	}

	var nodejs bool
	if c.FormValue("nodejs") == "nodejs" {
		nodejs = true
	}

	var typescript bool
	if c.FormValue("typescript") == "typescript" {
		typescript = true
	}

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_projects (title, start_date, end_date, description, reactjs, nextjs, nodejs, typescript, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		title, start, end, description, reactjs, nextjs, nodejs, typescript, duration(start, end))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func formUpdateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var projectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, description, reactjs, nextjs, nodejs, typescript, duration FROM tb_projects WHERE id=$1", id).Scan(
		&projectDetail.Id, &projectDetail.Title, &projectDetail.Start, &projectDetail.End, &projectDetail.Description, &projectDetail.Reactjs, &projectDetail.Nextjs, &projectDetail.Nodejs, &projectDetail.Typescript, &projectDetail.Duration)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"Projects":    projectDetail,
		"StartDate":   projectDetail.Start.Format("2006-01-02"),
		"EndDate":     projectDetail.End.Format("2006-01-02"),
		"DataSession": userData,
	}

	var tmpl, errTemp = template.ParseFiles("views/update-project.html")

	if errTemp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemp.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("inputTitle")
	start := c.FormValue("inputStart")
	end := c.FormValue("inputEnd")
	description := c.FormValue("inputDescription")

	var reactjs bool
	if c.FormValue("reactjs") == "reactjs" {
		reactjs = true
	}

	var nextjs bool
	if c.FormValue("nextjs") == "nextjs" {
		nextjs = true
	}

	var nodejs bool
	if c.FormValue("nodejs") == "nodejs" {
		nodejs = true
	}

	var typescript bool
	if c.FormValue("typescript") == "typescript" {
		typescript = true
	}

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_projects SET title=$1, start_date=$2, end_date=$3, description=$4, reactjs=$5, nextjs=$6, nodejs=$7, typescript=$8, duration=$9 WHERE id=$10",
		title, start, end, description, reactjs, nextjs, nodejs, typescript, duration(start, end), id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func duration(startDate string, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	hours := int(endTime.Sub(startTime).Hours())
	days := hours / 24
	weeks := days / 7
	months := weeks / 4
	years := months / 12

	var duration string

	if years > 1 {
		duration = strconv.Itoa(years) + " Tahun"
	} else if years > 0 {
		duration = strconv.Itoa(years) + " Tahun"
	} else {
		if months > 1 {
			duration = strconv.Itoa(months) + " Bulan"
		} else if months > 0 {
			duration = strconv.Itoa(months) + " Bulan"
		} else {
			if weeks > 1 {
				duration = strconv.Itoa(weeks) + " Minggu"
			} else if weeks > 0 {
				duration = strconv.Itoa(weeks) + " Minggu"
			} else {
				if days > 1 {
					duration = strconv.Itoa(days) + " Hari"
				} else {
					duration = strconv.Itoa(days) + " Hari"
				}
			}
		}
	}

	return duration
}

func formRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	session := map[string]interface{}{
		"DataSession": userData,
	}

	return tmpl.Execute(c.Response(), session)
}

func register(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("input-name")
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_users(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success!", true, "/form-login")
}

func formLogin(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	alerts := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "status")
	delete(sess.Values, "message")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), alerts)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_users WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800
	sess.Values["message"] = "Login Succes!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}
