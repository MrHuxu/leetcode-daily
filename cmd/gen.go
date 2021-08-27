package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MrHuxu/leetcode-daily/cmd/utils"

	"github.com/manifoldco/promptui"
)

func gen() {
	url, date := prompt()

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

	itemID := utils.ParseItemIDFromURL(url)
	item := utils.GetItem(itemID)
	question := utils.GetQuestion(item.Question.TitleSlug)

	ioutil.WriteFile(dir+"/description.md", []byte(fmt.Sprintf(`\>\> [题目链接](%s)

题意:

解答:

***original_content***

%s
`, url, question.Content)), 0666)

	ioutil.WriteFile(dir+"/main.go", []byte(fmt.Sprintf(`package main

%s
`, question.CodeDefinitions["golang"].DefaultCode)), 0666)

	time.Sleep(time.Millisecond * 5)
}

func prompt() (string, string) {
	urlPrompt := promptui.Prompt{
		Label:    "URL",
		Validate: urlValidator,
	}
	url, err := urlPrompt.Run()
	if err != nil {
		log.Fatalf("URL prompt failed %v\n", err)
	}

	datePrompt := promptui.Prompt{
		Label:    "Date",
		Validate: dateValidator,
	}
	date, err := datePrompt.Run()
	if err != nil {
		log.Fatalf("Date prompt failed %v\n", err)
	}

	return url, date
}

func urlValidator(input string) error {
	if !strings.HasPrefix(input, "http") {
		return errors.New("Invalid url")
	}

	if arr := strings.Split(input, "/"); len(arr) < 2 {
		return errors.New("Invalid URL")
	} else if _, err := strconv.Atoi(arr[len(arr)-2]); err != nil {
		return errors.New("Invalid url")
	}

	return nil
}

func dateValidator(input string) error {
	if len(input) != 0 && len(input) != 8 {
		return errors.New("Invalid Date")
	}
	return nil
}
