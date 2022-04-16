package engine

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"spider/config"
	"spider/fetcher"
	"spider/questioner"
	"spider/task"
)

type SectionListResponse struct {
	Data struct {
		Booklet struct {
			BaseInfo struct {
				Title string `json:"title"`
			} `json:"base_info"`
		} `json:"booklet"`
		Sections []*Section `json:"sections"`
	} `json:"data"`
}

func Run() {
	if _, err := ioutil.ReadFile(config.CookieFile); err != nil {
		task.WeChatLogin()
	}

	answers, err := questioner.Ask()
	if err != nil {
		panic(err)
	}
	log.Printf("BookId: %s", answers.BookId)

	sectionListBytes, err := fetcher.FetchSectionList(answers.BookId)
	if err != nil {
		panic(err)
	}

	var sectionListResponse SectionListResponse
	if err := json.Unmarshal(sectionListBytes, &sectionListResponse); err != nil {
		panic(err)
	}

	bookTitle := sectionListResponse.Data.Booklet.BaseInfo.Title

	log.Printf("bookTitle: %s", bookTitle)
	for index, section := range sectionListResponse.Data.Sections {
		section.BookTitle = bookTitle
		section.Index = index + 1
		log.Printf("SectionIndex: %d", section.Index)
		if err := downloadSection(section); err != nil {
			log.Printf("downloadSection error: %s", err)
		}
	}
}
