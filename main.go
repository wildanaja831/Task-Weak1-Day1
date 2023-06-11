package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id          int
	Title       string
	Start       string
	End         string
	Description string
	Reactjs     bool
	Nextjs      bool
	Nodejs      bool
	Typescript  bool
}

var dataProject = []Project{
	{
		Title:       "Mamang Racing Anjay",
		Start:       "11 April 2004",
		End:         "11 April 2104",
		Description: "Siap Kerja 24 Jam Non Stop!",
		Reactjs:     true,
		Nextjs:      true,
		Nodejs:      true,
		Typescript:  true,
	},
	{
		Title:       "Kerja Bagus!",
		Start:       "11 April 2004",
		End:         "11 April 2104",
		Description: "Permainan Yang Sangat Baik!",
		Reactjs:     true,
		Nextjs:      true,
		Nodejs:      true,
		Typescript:  true,
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")

	// Get
	e.GET("/", home)
	e.GET("/form-project", formProject)
	e.GET("/detail-project/:id", detailProject)
	e.GET("/contact", contact)
	e.GET("/update-project/:id", formUpdateProject)

	// Post
	e.POST("/add-project", addNewProject)
	e.POST("/update-project/:id", updateProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"Projects": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
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

	var projectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			projectDetail = Project{
				Title:       data.Title,
				Start:       data.Start,
				End:         data.End,
				Description: data.Description,
				Reactjs:     data.Reactjs,
				Nextjs:      data.Nextjs,
				Nodejs:      data.Nodejs,
				Typescript:  data.Typescript,
			}
		}
	}

	data := map[string]interface{}{
		"Projects": projectDetail,
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
	reactjs := c.FormValue("reactjs")
	nextjs := c.FormValue("nextjs")
	nodejs := c.FormValue("nodejs")
	typescript := c.FormValue("typescript")

	var newProject = Project{
		Title:       title,
		Start:       start,
		End:         end,
		Description: description,
		Reactjs:     (reactjs == "reactjs"),
		Nextjs:      (nextjs == "nextjs"),
		Nodejs:      (nodejs == "nodejs"),
		Typescript:  (typescript == "typescript"),
	}

	dataProject = append(dataProject, newProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func formUpdateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var projectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			projectDetail = Project{
				Title:       data.Title,
				Start:       data.Start,
				End:         data.End,
				Description: data.Description,
				Reactjs:     data.Reactjs,
				Nextjs:      data.Nextjs,
				Nodejs:      data.Nodejs,
				Typescript:  data.Typescript,
			}
		}
	}

	data := map[string]interface{}{
		"Projects": projectDetail,
	}

	var tmpl, err = template.ParseFiles("views/update-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("inputTitle")
	start := c.FormValue("inputStart")
	end := c.FormValue("inputEnd")
	description := c.FormValue("inputDescription")
	reactjs := c.FormValue("reactjs")
	nextjs := c.FormValue("nextjs")
	nodejs := c.FormValue("nodejs")
	typescript := c.FormValue("typescript")

	var newProject = Project{
		Id:          id,
		Title:       title,
		Start:       start,
		End:         end,
		Description: description,
		Reactjs:     (reactjs == "reactjs"),
		Nextjs:      (nextjs == "nextjs"),
		Nodejs:      (nodejs == "nodejs"),
		Typescript:  (typescript == "typescript"),
	}

	dataProject[id] = newProject

	return c.Redirect(http.StatusMovedPermanently, "/")
}
