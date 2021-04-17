package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MrHuxu/leetcode-daily/cmd/utils"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

var data = make(map[string]utils.ItemInfo)

func main() {
	if bs, err := os.ReadFile("./output/data.json"); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(bs, &data); err != nil {
			log.Fatal(err)
		}
	}

	server := gin.Default()
	server.LoadHTMLGlob("./website/templates/*")

	server.GET("/", index)
	server.GET("/:year", year)
	server.GET("/:year/:month", month)
	server.GET("/:year/:month/:day", day)

	server.Run(":3542")
}

func renderError(ctx *gin.Context, err error) {
	ctx.HTML(http.StatusInternalServerError, "base.tmpl", gin.H{
		"page":  "error",
		"error": err.Error(),
		"title": "Oops!!!",
	})
}

func index(ctx *gin.Context) {
	var years []Year
	var err error
	defer func() {
		if err != nil {
			renderError(ctx, err)
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page":  "index",
				"years": years,
				"title": "LeetCode Daily - xhu",
			})
		}
	}()

	entries, err := os.ReadDir(fmt.Sprintf("./questions"))
	if err != nil {
		return
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

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
			renderError(ctx, err)
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page":  "year",
				"year":  year,
				"title": fmt.Sprintf("LeetCode Daily - %s - xhu", year.Title),
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
			renderError(ctx, err)
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page":  "month",
				"month": month,
				"title": fmt.Sprintf("LeetCode Daily - %s - xhu", month.Title),
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
			renderError(ctx, err)
		} else {
			ctx.HTML(http.StatusOK, "base.tmpl", gin.H{
				"page":  "day",
				"day":   day,
				"title": fmt.Sprintf("LeetCode Daily - %s: %s - xhu", day.Title, day.ItemInfo.Question.Title),
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
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

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
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

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
	ItemInfo    utils.ItemInfo
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
	itemID := utils.ParseItemID(bs)
	d.ItemInfo = data[itemID]

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
