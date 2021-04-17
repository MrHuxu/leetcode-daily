package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MrHuxu/leetcode-daily/cmd/utils"
)

func fetch() {
	data := make(map[string]Item)
	if years, err := os.ReadDir("questions"); err != nil {
		log.Fatal(err)
	} else {
		for _, year := range years {
			if months, err := os.ReadDir(fmt.Sprintf("questions/%s", year.Name())); err != nil {
				log.Fatal(err)
			} else {
				for _, month := range months {
					if days, err := os.ReadDir(fmt.Sprintf("questions/%s/%s", year.Name(), month.Name())); err != nil {
						log.Fatal(err)
					} else {
						for _, day := range days {
							if strings.HasPrefix(day.Name(), ".") {
								continue
							}

							bs, err := os.ReadFile(fmt.Sprintf("questions/%s/%s/%s/description.md", year.Name(), month.Name(), day.Name()))
							if err != nil {
								log.Fatal(err)
							}

							itemID := utils.ParseItemID(bs)
							item := utils.GetItem(itemID)

							data[itemID] = Item{
								Title: item.Title,
							}
						}
					}
				}
			}
		}
	}

	bs, _ := json.Marshal(data)
	fmt.Println(string(bs))
}

type Item struct {
	Title string `json:"title"`
}
