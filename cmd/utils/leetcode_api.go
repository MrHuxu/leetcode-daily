package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func GetItem(itemID string) ItemInfo {
	bs, _ := json.Marshal(map[string]interface{}{
		"operationName": "GetItem",
		"variables": map[string]interface{}{
			"itemId": itemID,
		},
		"query": `
			query GetItem($itemId: String!) {
			  item(id: $itemId) {
			    id
			    title
			    type
			    paidOnly
			    lang
			    isEligibleForCompletion
			    hasAppliedTimeTravelTicket
			    question {
			      questionId
			      title
			      titleSlug
			    }
		  }
		}
		`,
	})

	resp, err := http.Post("https://leetcode.com/graphql", "application/json", bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}

	var item Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		log.Fatal(err)
	}
	return item.Data.Item
}

type Item struct {
	Data struct {
		Item ItemInfo `json:"item"`
	} `json:"data"`
}

type ItemInfo struct {
	ID                         string      `json:"id"`
	Title                      string      `json:"title"`
	Type                       int         `json:"type"`
	PaidOnly                   bool        `json:"paidOnly"`
	Lang                       interface{} `json:"lang"`
	IsEligibleForCompletion    bool        `json:"isEligibleForCompletion"`
	HasAppliedTimeTravelTicket bool        `json:"hasAppliedTimeTravelTicket"`
	Question                   struct {
		QuestionID string `json:"questionId"`
		Title      string `json:"title"`
		TitleSlug  string `json:"titleSlug"`
	} `json:"question"`
}

func GetQuestion(slug string) QuestionInfo {
	bs, _ := json.Marshal(map[string]interface{}{
		"operationName": "GetQuestion",
		"variables": map[string]interface{}{
			"titleSlug": slug,
		},
		"query": `
			query GetQuestion($titleSlug: String!) {
			  question(titleSlug: $titleSlug) {
			    questionId
			    questionTitle
			    categoryTitle
			    codeDefinition
			    sampleTestCase
			    content
		  }
		}
		`,
	})

	resp, err := http.Post("https://leetcode.com/graphql", "application/json", bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}

	var question Question
	if err := json.NewDecoder(resp.Body).Decode(&question); err != nil {
		log.Fatal(err)
	}

	var tmp []CodeDefinition
	if err := json.Unmarshal([]byte(question.Data.Question.CodeDefinitionsRaw), &tmp); err != nil {
		log.Fatal(err)
	}
	question.Data.Question.CodeDefinitions = make(map[string]CodeDefinition)
	for _, definition := range tmp {
		question.Data.Question.CodeDefinitions[definition.Value] = definition
	}

	return question.Data.Question
}

type Question struct {
	Data struct {
		Question QuestionInfo `json:"question"`
	} `json:"data"`
}

type QuestionInfo struct {
	QuestionID         string `json:"questionId"`
	QuestionTitle      string `json:"questionTitle"`
	CategoryTitle      string `json:"categoryTitle"`
	CodeDefinitionsRaw string `json:"codeDefinition"`
	CodeDefinitions    map[string]CodeDefinition
	SampleTestCase     string `json:"sampleTestCase"`
	Content            string `json:"content"`
}

type CodeDefinition struct {
	Value       string `json:"value"`
	Text        string `json:"text"`
	DefaultCode string `json:"defaultCode"`
}
