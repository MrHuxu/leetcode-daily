package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MrHuxu/leetcode-daily/cmd/utils"
)

func gen(url string, date string) {

	var year, month, day string
	if date == "" {
		now := time.Now()
		year = strconv.Itoa(now.Year())

		month = strconv.Itoa(int(now.Month()))
		if now.Month() < 10 {
			month = "0" + month
		}

		day = strconv.Itoa(now.Day())
		if now.Day() < 10 {
			day = "0" + day
		}
	} else {
		year, month, day = date[:4], date[4:6], date[6:8]
	}

	dir := fmt.Sprintf("questions/%s/%s/%s", year, month, day)
	os.MkdirAll(dir, os.ModePerm)

	os.WriteFile(dir+"/description.md", []byte(fmt.Sprintf(`\>\> [题目链接](%s)

题意:

解答:
`, url)), 0666)

	itemID := utils.ParseItemIDFromURL(url)
	item := utils.GetItem(itemID)
	question := utils.GetQuestion(item.Question.TitleSlug)

	os.WriteFile(dir+"/main.go", []byte(fmt.Sprintf(`package main

%s
`, question.CodeDefinitions["golang"].DefaultCode)), 0666)
	os.WriteFile(dir+"/main_test.go", []byte(`package main`), 0666)
}
