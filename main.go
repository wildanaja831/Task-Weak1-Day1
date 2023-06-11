package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id          int
	Title       string
	Start       string
	End         string
	Duration    string
	Description string
	Reactjs     bool
	Nextjs      bool
	Nodejs      bool
	Typescript  bool
}

var dataProject = []Project{
	{
		Title:       "Mamang Racing Anjay",
		Start:       "2023-01-15",
		End:         "2023-05-15",
		Duration:    duration("2023-01-15", "2023-05-15"),
		Description: "Siap Kerja 24 Jam Non Stop!",
		Reactjs:     true,
		Nextjs:      true,
		Nodejs:      true,
		Typescript:  true,
	},
	{
		Title:       "Kerja Bagus!",
		Start:       "2023-02-15",
		End:         "2024-05-15",
		Duration:    duration("2023-02-15", "2024-05-15"),
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
	e.POST("/delete-project/:id", deleteProject)

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
				Duration:    data.Duration,
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
		Duration:    duration(start, end),
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
				Id:          id,
				Title:       data.Title,
				Start:       data.Start,
				End:         data.End,
				Duration:    data.Duration,
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
		Title:       title,
		Start:       start,
		End:         end,
		Duration:    duration(start, end),
		Description: description,
		Reactjs:     (reactjs == "reactjs"),
		Nextjs:      (nextjs == "nextjs"),
		Nodejs:      (nodejs == "nodejs"),
		Typescript:  (typescript == "typescript"),
	}

	dataProject[id] = newProject

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

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
