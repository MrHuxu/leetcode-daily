package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/MrHuxu/leetcode-daily/cmd/utils"

	"github.com/gosuri/uiprogress"
)

func fetch() {
	var itemIDs []string
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

							itemIDs = append(itemIDs, utils.ParseItemID(bs))
						}
					}
				}
			}
		}
	}

	uiprogress.Start()
	bar := uiprogress.AddBar(len(itemIDs)).AppendCompleted().PrependElapsed()
	bar.Incr()

	data := make(map[string]utils.ItemInfo)
	var mu sync.Mutex

	ch := make(chan string, 5)

	go func() {
		for _, itemID := range itemIDs {
			ch <- itemID
		}
		close(ch)
	}()

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for itemID := range ch {
				item := utils.GetItem(itemID)

				mu.Lock()
				data[item.ID] = item
				mu.Unlock()

				bar.Incr()
			}
		}()
	}
	wg.Wait()

	bs, _ := json.Marshal(data)
	os.WriteFile("output/data.json", bs, 0666)
}
