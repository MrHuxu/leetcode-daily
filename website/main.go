package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("./website/templates/*")

	server.GET("/", index)
	server.GET("/:year", year)
	server.GET("/:year/:month", month)
	server.GET("/:year/:month/:day", day)

	server.Run(":3542")
}

func index(ctx *gin.Context) {
	var years []Year
	var err error
	defer func() {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page":  "index",
				"years": years,
			})
		}
	}()

	entries, err := os.ReadDir(fmt.Sprintf("./questions"))
	if err != nil {
		return
	}

	for _, entry := range entries {
		y, err := getYear(entry.Name())
		if err != nil {
			return
		}
		years = append(years, y)
	}
}

func year(ctx *gin.Context) {
	var year Year
	var err error
	defer func() {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page": "year",
				"year": year,
			})
		}
	}()

	year, err = getYear(ctx.Param("year"))
}

func month(ctx *gin.Context) {
	var month Month
	var err error
	defer func() {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page":  "month",
				"month": month,
			})
		}
	}()

	month, err = getMonth(ctx.Param("year"), ctx.Param("month"))
}

func day(ctx *gin.Context) {
	var day Day
	var err error
	defer func() {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page": "day",
				"day":  day,
			})
		}
	}()

	day, err = getDay(ctx.Param("year"), ctx.Param("month"), ctx.Param("day"))
}

type Breadcrumb struct {
	Path, Label string
}

type Year struct {
	Idx, Path, Title string
	Breadcrumbs      []Breadcrumb

	Months []Month
}

func getYear(year string) (y Year, err error) {
	y.Idx = year
	y.Path = fmt.Sprintf("/%s", year)
	y.Title = fmt.Sprintf("%s", year)
	y.Breadcrumbs = []Breadcrumb{{
		Path:  fmt.Sprintf("/%s", year),
		Label: year,
	}}

	entries, err := os.ReadDir(fmt.Sprintf("./questions/%s", year))
	for _, entry := range entries {
		m, innerErr := getMonth(year, entry.Name())
		if innerErr != nil {
			err = innerErr
			return
		}
		y.Months = append(y.Months, m)
	}

	return
}

type Month struct {
	Idx, Path, Title string
	Breadcrumbs      []Breadcrumb

	Days []Day
}

func getMonth(year, month string) (m Month, err error) {
	m.Idx = month
	m.Path = fmt.Sprintf("/%s/%s", year, month)
	m.Title = fmt.Sprintf("%s.%s", year, month)
	m.Breadcrumbs = []Breadcrumb{{
		Path:  fmt.Sprintf("/%s", year),
		Label: year,
	}, {
		Path:  fmt.Sprintf("/%s/%s", year, month),
		Label: month,
	}}

	entries, err := os.ReadDir(fmt.Sprintf("./questions/%s/%s", year, month))
	for _, entry := range entries {
		d, innerErr := getDay(year, month, entry.Name())
		if innerErr != nil {
			err = innerErr
			return
		}
		m.Days = append(m.Days, d)
	}

	return
}

type Day struct {
	Idx, Path, Title string
	Breadcrumbs      []Breadcrumb

	Description template.HTML
	Code        string
	Test        string
}

func getDay(year, month, day string) (d Day, err error) {
	d.Idx = day
	d.Path = fmt.Sprintf("/%s/%s/%s", year, month, day)
	d.Title = fmt.Sprintf("%s.%s.%s", year, month, day)
	d.Breadcrumbs = []Breadcrumb{{
		Path:  fmt.Sprintf("/%s", year),
		Label: year,
	}, {
		Path:  fmt.Sprintf("/%s/%s", year, month),
		Label: month,
	}, {
		Path:  fmt.Sprintf("/%s/%s/%s", year, month, day),
		Label: day,
	}}

	bs, err := os.ReadFile(fmt.Sprintf("./questions/%s/%s/%s/description.md", year, month, day))
	if err == nil {
		d.Description = template.HTML(blackfriday.Run(bs))
	}

	bs, err = os.ReadFile(fmt.Sprintf("./questions/%s/%s/%s/main.go", year, month, day))
	if err != nil {
		return
	}
	d.Code = string(bs)

	bs, err = os.ReadFile(fmt.Sprintf("./questions/%s/%s/%s/main_test.go", year, month, day))
	if err != nil {
		return
	}
	d.Test = string(bs)

	return
}
