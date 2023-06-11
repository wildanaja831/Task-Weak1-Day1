package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/form-project", formProject)
	e.GET("/detail-project/:id", detailProject)
	e.GET("/contact", contact)
	e.POST("/add-project", addNewProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func formProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func detailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"ID":          id,
		"Title":       "Kerja Bagus!",
		"Start":       "11 April 2004",
		"End":         "Tak Terhingga",
		"Description": "Kerja Bagus Dan Luar biasa!",
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addNewProject(c echo.Context) error {
	title := c.FormValue("inputTitle")
	start := c.FormValue("inputStart")
	end := c.FormValue("inputEnd")
	description := c.FormValue("inputDescription")

	println("Title :" + title)
	println("Start Date :" + start)
	println("End Date :" + end)
	println("Description :" + description)

	return c.Redirect(http.StatusMovedPermanently, "/add-project")
}
